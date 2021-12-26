package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Eips(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Eips(),
		"./snapshots/ec2")
}
