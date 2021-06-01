//+build integration_skip all

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

// todo use some domain for tests of route53 related resources
func TestIntegrationApigatewayDomainNames(t *testing.T) {
	awsTestIntegrationHelper(t, ApigatewayDomainNames(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_domain_names",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"endpoint_configuration_types": []interface{}{"REGIONAL"},
					"api_key_source":               "HEADER",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
		}
	})
}
