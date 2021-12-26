package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2EbsVolumes(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2EbsVolumes(),
		"./snapshots/ec2")
}
