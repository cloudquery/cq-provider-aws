package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRedshiftSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RedshiftSubnetGroups(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_redshift_subnet_groups",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"cluster_subnet_group_name": "redshift-subnet-group-1",
					"description":               "my test description",
					"tags": map[string]interface{}{
						"TestId":      res.Suffix,
						"Type":        "integration_test",
						"environment": "Production",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_redshift_subnet_group_subnets",
					ForeignKeyName: "subnet_group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 2,
					}},
				},
			},
		}
	})
}
