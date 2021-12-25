// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcsTaskDefinitions(t *testing.T) {
	resource := resources.EcsTaskDefinitions()
	awsTestIntegrationHelper(t, resource)
}
