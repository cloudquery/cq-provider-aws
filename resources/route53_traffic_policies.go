package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53TrafficPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_traffic_policies",
		Resolver:     fetchRoute53TrafficPolicies,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "document",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "version",
				Type: schema.TypeInt,
			},
			{
				Name: "comment",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53TrafficPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config route53.ListTrafficPoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListTrafficPolicies(ctx, &config, func(o *route53.Options) {})
		if err != nil {
			return err
		}

		for _, tps := range response.TrafficPolicySummaries {
			tpResponse, err := svc.GetTrafficPolicy(ctx, &route53.GetTrafficPolicyInput{Id: tps.Id}, func(o *route53.Options) {})
			if err != nil {
				return err
			}
			res <- tpResponse.TrafficPolicy
		}

		if aws.ToString(response.TrafficPolicyIdMarker) == "" {
			break
		}
		config.TrafficPolicyIdMarker = response.TrafficPolicyIdMarker
	}
	return nil
}
