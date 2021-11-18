package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEcsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsClusters(), []string{"aws_ecs_clusters.tf", "aws_vpc.tf", "aws_elbv2_load_balancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ecs_clusters",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"name":   fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix),
							"stage":  "test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
