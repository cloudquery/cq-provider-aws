package client

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
)

func newRetryer(maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return retry.NewStandard(func(o *retry.StandardOptions) {
			o.MaxAttempts = maxRetries
			o.MaxBackoff = time.Second * time.Duration(maxBackoff)
		})
	}
}
