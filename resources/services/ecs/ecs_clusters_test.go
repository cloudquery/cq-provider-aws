// +build integration

package ecs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsClusters(),
		"./snapshots")
}
