// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsClusterSnapshots(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusterSnapshots(), "./snapshots/rds")
}
