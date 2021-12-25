// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsEventSubscriptions(t *testing.T) {
	table := resources.RdsEventSubscriptions()
	awsTestIntegrationHelper(t, table)
}
