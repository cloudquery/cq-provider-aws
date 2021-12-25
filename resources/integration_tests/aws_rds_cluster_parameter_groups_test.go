// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsClusterParameterGroups(t *testing.T) {
	table := resources.RdsClusterParameterGroups()
	awsTestIntegrationHelper(t, table)
}
