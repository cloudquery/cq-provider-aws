// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudwatchAlarms(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudwatchAlarms(),
		"./snapshots/cloudwatch")
}
