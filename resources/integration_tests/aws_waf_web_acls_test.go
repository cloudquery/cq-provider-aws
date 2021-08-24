package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWAFWebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafWebAcls(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_waf_web_acls",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                "waf_web_acl_1",
					"metric_name":         "wafwebacl1",
					"default_action_type": "ALLOW",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{{
				Name:           "aws_waf_web_acl_rules",
				ForeignKeyName: "web_acl_cq_id",
				ExpectedValues: []providertest.ExpectedValue{{
					Count: 1,
					Data: map[string]interface{}{
						"priority":    float64(1),
						"action_type": "BLOCK",
						"type":        "REGULAR",
					},
				}},
			}},
		}
	})
}
