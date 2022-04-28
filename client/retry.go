package client

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/smithy-go"
	"github.com/hashicorp/go-hclog"
)

type retryer struct {
	aws.Retryer
	logger hclog.Logger
}

func newRetryer(logger hclog.Logger, maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return &retryer{
			//Retryer: retry.NewStandard(func(o *retry.StandardOptions) {
			//	o.MaxAttempts = maxRetries
			//	o.MaxBackoff = time.Second * time.Duration(maxBackoff)
			//}),
			Retryer: retry.NewAdaptiveMode(func(o *retry.AdaptiveModeOptions) {
				o.StandardOptions = []func(o *retry.StandardOptions){
					func(o *retry.StandardOptions) {
						o.MaxAttempts = maxRetries
						o.MaxBackoff = time.Second * time.Duration(maxBackoff)
					},
				}
			}),
			logger: logger,
		}
	}
}

func (r *retryer) RetryDelay(attempt int, err error) (time.Duration, error) {
	dur, retErr := r.Retryer.RetryDelay(attempt, err)

	logParams := []interface{}{
		"duration", dur.String(),
		"attempt", attempt,
		"retrier_err", retErr,
		"err", err,
	}
	var oe *smithy.OperationError
	if errors.As(err, &oe) {
		logParams = append(logParams, []interface{}{
			"op", oe.Operation(),
			"service_id", oe.Service(),
		})
	}
	if retErr != nil {
		r.logger.Debug("retryDelay returned err", logParams...)
	} else {
		r.logger.Debug("retryDelay will wait", logParams...)
	}
	return dur, retErr
}
