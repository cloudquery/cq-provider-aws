// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationAutoscalingLaunchConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t,
		resources.AutoscalingLaunchConfigurations(),
		"./snapshots/autoscaling")
}
