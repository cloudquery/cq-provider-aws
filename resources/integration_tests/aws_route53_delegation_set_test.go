package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRoute53DelegationSet(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Route53ReusableDelegationSets(),
		"./snapshots/route53")
}
