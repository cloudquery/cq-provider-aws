//go:build integration

package provider

import (
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {

	cfg := `
	accounts "cq-provider-aws" {
		role_arn = "arn:aws:iam::707066037989:role/CqProviderAWSAssumeRoleTest"
	}
	accounts "cq-dev" {
		role_arn = "arn:aws:iam::704956590351:role/CqProviderAWSAssumeRoleTest"
	}
  	max_retries = 20
	max_backoff = 60
	`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider:              Provider(),
		Config:                cfg,
		NotParallel:           true,
		ParallelFetchingLimit: 10000,
	})
}
