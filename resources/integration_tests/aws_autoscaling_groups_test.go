// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationAutoscalingGroups(t *testing.T) {
	resource := resources.AutoscalingGroups()
	awsTestIntegrationHelper(t, resource)
}
