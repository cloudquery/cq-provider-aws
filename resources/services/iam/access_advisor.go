package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const MAX_GOROUTINES = 10

func fetchDetailEntities(ctx context.Context, res chan<- interface{}, svc client.IamClient, sla types.ServiceLastAccessed, jobId string) error {
	config := iam.GetServiceLastAccessedDetailsWithEntitiesInput{
		JobId:            &jobId,
		ServiceNamespace: sla.ServiceNamespace,
	}
	details := AccessedDetails{
		ServiceLastAccessed: sla,
	}
	for {
		output, err := svc.GetServiceLastAccessedDetailsWithEntities(ctx, &config)
		if err != nil {
			return err
		}
		details.Entities = append(details.Entities, output.EntityDetailsList...)
		if output.Marker == nil {
			break
		}
		if output.Marker != nil {
			config.Marker = output.Marker
		}
	}
	res <- details
	return nil
}

func fetchIamAccessDetails(ctx context.Context, res chan<- interface{}, svc client.IamClient, arn string) error {
	config := iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         &arn,
		Granularity: types.AccessAdvisorUsageGranularityTypeActionLevel,
	}
	output, err := svc.GenerateServiceLastAccessedDetails(ctx, &config)
	if err != nil {
		return diag.WrapError(err)
	}

	getDetails := iam.GetServiceLastAccessedDetailsInput{
		JobId: output.JobId,
	}
	for {
		details, err := svc.GetServiceLastAccessedDetails(ctx, &getDetails)
		if err != nil {
			return diag.WrapError(err)
		}

		requestTime := 0
		switch details.JobStatus {
		case types.JobStatusTypeInProgress:
			time.Sleep(time.Millisecond * 200)
			requestTime += 200
			continue
		case types.JobStatusTypeFailed:
			return diag.WrapError(fmt.Errorf("failed to get last acessed details with error: %s - %s", *details.Error.Code, *details.Error.Message))
		case types.JobStatusTypeCompleted:
			var sem = semaphore.NewWeighted(int64(MAX_GOROUTINES))
			errs, ctx := errgroup.WithContext(ctx)
			for _, s := range details.ServicesLastAccessed {
				if *s.TotalAuthenticatedEntities > 0 {
					if err := sem.Acquire(ctx, 1); err != nil {
						return diag.WrapError(err)
					}
					func(sla types.ServiceLastAccessed, jobId string) {
						errs.Go(func() error {
							defer sem.Release(1)
							return fetchDetailEntities(ctx, res, svc, sla, jobId)
						})
					}(s, *output.JobId)

				}
			}
			err = errs.Wait()
			if err != nil {
				return diag.WrapError(err)
			}
			if details.Marker == nil {
				return nil
			}
			if details.Marker != nil {
				getDetails.Marker = details.Marker
			}
		}
	}
}
