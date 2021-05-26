package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	verification := providertest.ResourceIntegrationVerification{
		Name: "aws_apigateway_rest_apis",
		Values: map[string]interface{}{
			"endpoint_configuration_types": []interface{}{"REGIONAL"},
		},
		Children: []*providertest.ResourceIntegrationVerification{
			{
				Name:           "aws_apigateway_rest_api_deployments",
				ForeginKeyName: "rest_api_id",
				Values: map[string]interface{}{
					"description": "test description",
				},
			},
		},
	}
	awsTestIntegrationHelper(t, ApigatewayRestApis(), verification)
}
