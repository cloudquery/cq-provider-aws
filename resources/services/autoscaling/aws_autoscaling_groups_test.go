// +build integration

package autoscaling

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationAutoscalingGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.AutoscalingGroups(),
		"./snapshots/autoscaling")
}
