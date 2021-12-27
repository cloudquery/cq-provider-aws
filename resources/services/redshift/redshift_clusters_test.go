// +build integration

package redshift

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRedshiftClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RedshiftClusters(),
		"./snapshots")
}
