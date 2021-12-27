// +build integration

package emr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEmrClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EmrClusters(),
		"./snapshots")
}
