// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsDBSnapshots(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsDbSnapshots(),
		"./snapshots")
}
