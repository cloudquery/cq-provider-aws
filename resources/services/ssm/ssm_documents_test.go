// +build integration

package ssm

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSSMDocuments(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SsmDocuments(),
		"./snapshots")
}
