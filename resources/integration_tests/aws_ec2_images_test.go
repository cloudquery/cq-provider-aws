// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Images(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Images(),
		"./snapshots/ec2")
}
