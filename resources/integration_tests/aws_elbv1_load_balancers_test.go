// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv1LoadBalancers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv1LoadBalancers(),
		"./snapshots/elbv1")
}
