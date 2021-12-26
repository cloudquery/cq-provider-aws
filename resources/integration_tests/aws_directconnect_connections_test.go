// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDirectconnectConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectConnections(),
		"./snapshots/directconnect")
}
