package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRoute53HostedZones(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Route53HostedZones(),
		"./snapshots/route53")
}
