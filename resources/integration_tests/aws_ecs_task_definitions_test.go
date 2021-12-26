package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcsTaskDefinitions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsTaskDefinitions(),
		"./snapshots/ecs")
}
