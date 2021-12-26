// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEfsFilesystems(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EfsFilesystems(),
		"./snapshots/efs")
}
