package integration_tests

import (
	"github.com/cloudquery/cq-provider-aws/resources"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationFsxBackups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.FsxBackups(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_fsx_backups",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
				},
			},
		}
	})
}
