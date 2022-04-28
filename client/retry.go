package client

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/hashicorp/go-hclog"
)

type retryer struct {
	*retry.Standard
	backoff *retry.ExponentialJitterBackoff
	logger  hclog.Logger
}

func newRetryer(logger hclog.Logger, maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		maxbackoff := time.Second * time.Duration(maxBackoff)
		backoff := retry.NewExponentialJitterBackoff(maxbackoff)

		return &retryer{
			Standard: retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = maxRetries
				o.MaxBackoff = maxbackoff
				o.Backoff = backoff
			}),
			backoff: backoff,
			logger:  logger,
		}
	}
}

func (r *retryer) RetryDelay(attempt int, err error) (time.Duration, error) {
	dur, err := r.backoff.BackoffDelay(attempt, err)
	if err != nil {
		r.logger.Debug("retryDelay returned err", "err", err, "duration", dur.String())
	} else {
		r.logger.Debug("retryDelay will wait", "duration", dur.String())
	}
	return dur, err
}
