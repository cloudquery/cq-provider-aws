package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Instances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Instances(),
		"./snapshots/ec2")
}
