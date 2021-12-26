package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationFsxBackups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.FsxBackups(),
		"./snapshots/fsx")
}
