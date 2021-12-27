// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsClusterSnapshots(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusterSnapshots(), "./snapshots")
}
