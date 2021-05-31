//+build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationLambdaFunctions(t *testing.T) {
	awsTestIntegrationHelper(t, LambdaFunctions(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_lambda_functions",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"tracing_config_mode": "PassThrough",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_lambda_function_aliases",
					ForeignKeyName: "function_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"description": "a sample description",
						},
					}},
				},
			},
		}
	})
}
