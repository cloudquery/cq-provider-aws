package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayUsagePlans(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayUsagePlans(), []string{"aws_apigateway_rest_apis.tf", "aws_apigateway_usage_plans.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_usage_plans",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                 "my-usage-plan",
					"description":          "my description",
					"product_code":         "MYCODE",
					"quota_limit":          float64(20),
					"quota_offset":         float64(2),
					"quota_period":         "WEEK",
					"throttle_burst_limit": float64(5),
					"throttle_rate_limit":  float64(10),
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_apigateway_usage_plan_api_stages",
					ForeignKeyName: "usage_plan_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"stage": "cqtestingstage1integration",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"stage": "cqtestingstage2integration",
							},
						},
					},
				},
				{
					Name:           "aws_apigateway_usage_plan_keys",
					ForeignKeyName: "usage_plan_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name": "my_key",
							"type": "API_KEY",
						},
					}},
				},
			},
		}
	})
}
