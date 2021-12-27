// +build integration

package fsx

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationFsxBackups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.FsxBackups(),
		"./snapshots")
}
