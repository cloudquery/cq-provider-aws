// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsClusterParameterGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusterParameterGroups(),
		"./snapshots")
}
