// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudwatchlogsFilters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudwatchlogsFilters(),
		"./snapshots/cloudwatchlogs")
}
