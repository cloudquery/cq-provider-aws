//+build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	awsTestIntegrationHelper(t, ApigatewayRestApis(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_rest_apis",
			Values: []map[string]interface{}{{
				"endpoint_configuration_types": []interface{}{"REGIONAL"},
				"api_key_source":               "HEADER",
			}},
			Children: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_apigateway_rest_api_deployments",
					ForeignKeyName: "rest_api_id",
					Values: []map[string]interface{}{{
						"description": "test description",
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_authorizers",
					ForeignKeyName: "rest_api_id",
					Values: []map[string]interface{}{{
						"auth_type":                        "custom",
						"authorizer_result_ttl_in_seconds": float64(500),
						"type":                             "TOKEN",
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_stages",
					ForeignKeyName: "rest_api_id",
					Values: []map[string]interface{}{{
						"tracing_enabled": false,
						"tags": map[string]interface{}{
							"hello": "world",
						},
					}},
				},
			},
		}
	})
}
