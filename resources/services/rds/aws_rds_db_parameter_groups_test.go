// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsDbParameterGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsDbParameterGroups(), "./snapshots")
}
