// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSSMDocuments(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SsmDocuments(),
		"./snapshots/ssm")
}
