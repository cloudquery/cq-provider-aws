// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDirectconnectLags(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectLags(),
		"./snapshots/directconnect")
}
