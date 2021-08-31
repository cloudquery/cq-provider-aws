package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFv2WebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2WebAcls(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_wafv2_web_acls",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":        "wafv2-web-acl-" + res.Prefix + res.Suffix,
					"description": "Example of a managed rule.",
					"default_action": map[string]interface{}{
						"Allow": map[string]interface{}{
							"CustomRequestHandling": nil,
						},
						"Block": nil,
					},
					"visibility_config_cloud_watch_metrics_enabled": false,
					"visibility_config_metric_name":                 "friendly-metric-name",
					"visibility_config_sampled_requests_enabled":    false,
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
						"Tag1":   "Value1",
						"Tag2":   "Value2",
					},
				},
			}},
		}
	})
}
