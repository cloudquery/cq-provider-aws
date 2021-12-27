// +build integration

package waf

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationWAFRules(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafRules(),
		"./snapshots")
}
