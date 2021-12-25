// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv2LoadBalancers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2LoadBalancers())
}
