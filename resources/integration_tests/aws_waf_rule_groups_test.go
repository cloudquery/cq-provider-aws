package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationWAFRuleGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafRuleGroups(),
		"./snapshots/waf")
}
