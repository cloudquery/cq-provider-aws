// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsDbParameterGroups(t *testing.T) {
	table := resources.RdsDbParameterGroups()
	awsTestIntegrationHelper(t, table)
}
