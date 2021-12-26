package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsEventSubscriptions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsEventSubscriptions(),
		"./snapshots/rds")
}
