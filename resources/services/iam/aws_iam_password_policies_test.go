// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIamPasswordPolicies(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamPasswordPolicies(),
		"./snapshots")
}
