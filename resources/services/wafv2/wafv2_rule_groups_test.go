// +build integration

package wafv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationWAFv2RuleGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2RuleGroups(),
		"./snapshots")
}
