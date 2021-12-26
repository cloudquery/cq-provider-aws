// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2FlowLogs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2FlowLogs(),
		"./snapshots/ec2")
}
