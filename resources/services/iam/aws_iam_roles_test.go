// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamRoles(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamRoles(),
		"./snapshots")
}
