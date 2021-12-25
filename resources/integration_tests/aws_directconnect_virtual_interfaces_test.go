// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDirectconnectVirtualInterfaces(t *testing.T) {
	t.Skipf("resource missing")
	awsTestIntegrationHelper(t, resources.DirectconnectVirtualInterfaces())
}
