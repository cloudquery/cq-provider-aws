// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv2TargetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2TargetGroups())
}
