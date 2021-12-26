// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamRoles(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamRoles(),
		"./snapshots/iam")
}
