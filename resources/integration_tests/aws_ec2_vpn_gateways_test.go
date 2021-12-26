package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2VpnGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpnGateways(),
		"./snapshots/ec2")
}
