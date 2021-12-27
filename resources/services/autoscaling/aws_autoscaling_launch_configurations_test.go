// +build integration

package autoscaling

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationAutoscalingLaunchConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t,
		resources.AutoscalingLaunchConfigurations(),
		"./snapshots")
}
