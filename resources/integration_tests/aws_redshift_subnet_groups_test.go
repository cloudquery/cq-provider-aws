// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRedshiftSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RedshiftSubnetGroups(),
		"./snapshots/redshift")
}
