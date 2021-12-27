// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsInstances(),
		"./snapshots")
}

func TestIntegrationRdsSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsSubnetGroups(),
		"./snapshots")
}

func TestIntegrationRdsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusters(),
		"./snapshots")
}
