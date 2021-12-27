// +build integration

package ecr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcrRepositories(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcrRepositories(),
		"./snapshots")
}
