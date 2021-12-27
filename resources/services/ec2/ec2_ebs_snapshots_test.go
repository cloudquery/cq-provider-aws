// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2EbsSnapshots(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2EbsSnapshots(),
		"./snapshots")
}
