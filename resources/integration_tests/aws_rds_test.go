package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsInstances(),
		"./snapshots/rds")
}

func TestIntegrationRdsSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsSubnetGroups(),
		"./snapshots/rds")
}

func TestIntegrationRdsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusters(),
		"./snapshots/rds")
}
