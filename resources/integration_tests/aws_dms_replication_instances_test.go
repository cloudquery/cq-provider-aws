package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDmsReplicationInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DmsReplicationInstances(),
		"./snapshots/dms")
}
