// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsDBSnapshots(t *testing.T) {
	table := resources.RdsDbSnapshots()
	awsTestIntegrationHelper(t, table)
}
