package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcEndpoints(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpc_endpoints",
			//Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
			//	return sq.Where(squirrel.And{
			//		squirrel.Eq{"tags->>'TestId'": res.Suffix},
			//		squirrel.NotEq{"status_code": "deleted"},
			//	})
			//},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"vpc_endpoint_type":   "Gateway",
						"requester_managed":   false,
						"private_dns_enabled": false,
						"tags": map[string]interface{}{
							"Type":        "integration_test",
							"Environment": "test",
							"TestId":      res.Suffix,
						},
					},
				},
			},
		}
	})
}
