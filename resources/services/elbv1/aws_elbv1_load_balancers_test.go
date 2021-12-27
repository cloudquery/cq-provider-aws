// +build integration

package elbv1

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv1LoadBalancers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv1LoadBalancers(),
		"./snapshots")
}
