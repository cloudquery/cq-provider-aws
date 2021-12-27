// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamUsers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamUsers(),
		"./snapshots")
}
