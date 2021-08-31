package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationAccessAnalyzers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.AccessAnalyzerAnalyzer(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_access_analyzer_analyzers",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": "analyzer-name",
					"type": "ACCOUNT",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
		}
	})
}
