// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDirectconnectVirtualInterfaces(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectVirtualInterfaces(),
		"./snapshots/directconnect")
}
