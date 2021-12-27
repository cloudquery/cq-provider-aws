// +build integration

package cloudwatch

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudwatchAlarms(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudwatchAlarms(),
		"./snapshots")
}
