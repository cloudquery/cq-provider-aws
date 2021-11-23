package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	resource := resources.CodebuildProjects()
	awsTestIntegrationHelper(t, resource, []string{"aws_codebuild_projects.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resource.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("project-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"project_visibility":        "PRIVATE",
						"queued_timeout_in_minutes": float64(480),
						"source_type":               "GITHUB",
					},
				},
			},
		}
	})
}
