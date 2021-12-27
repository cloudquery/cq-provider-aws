// +build integration

package waf

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationWAFWebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.WafWebAcls(),
		"./snapshots")
}
