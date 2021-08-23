package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElbv2TargetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2TargetGroups(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elbv2_target_groups",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":     fmt.Sprintf("%s%s", res.Prefix, res.Suffix)[0:31],
						"protocol": "HTTP",
						"port":     float64(80),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix)[0:31]})
			},
		}
	})
}
