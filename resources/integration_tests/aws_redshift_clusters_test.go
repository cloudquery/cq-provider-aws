// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRedshiftClusters(t *testing.T) {
	const clusterFKName = "cluster_cq_id"
	awsTestIntegrationHelper(t, resources.RedshiftClusters())
}
