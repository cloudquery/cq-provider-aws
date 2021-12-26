// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcrRepositories(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcrRepositories(),
		"./snapshots/ecr")
}
