// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamGroups(),
		"./snapshots")
}
