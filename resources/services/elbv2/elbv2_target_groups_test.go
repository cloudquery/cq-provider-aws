// +build integration

package elbv2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElbv2TargetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Elbv2TargetGroups(),
		"./snapshots")
}
