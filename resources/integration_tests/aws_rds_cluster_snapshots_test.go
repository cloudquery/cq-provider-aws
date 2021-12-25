// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsClusterSnapshots(t *testing.T) {
	table := resources.RdsClusterSnapshots()
	awsTestIntegrationHelper(t, table)
}
