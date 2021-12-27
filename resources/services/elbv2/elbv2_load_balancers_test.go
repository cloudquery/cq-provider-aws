// +build integration

package elbv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv2LoadBalancers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2LoadBalancers(),
		"./snapshots")
}
