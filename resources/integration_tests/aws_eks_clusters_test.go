package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEksClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EksClusters(),
		"./snapshots/eks")
}
