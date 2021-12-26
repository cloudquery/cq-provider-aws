// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationGuarddutyDetectors(t *testing.T) {
	awsTestIntegrationHelper(t, resources.GuarddutyDetectors(),
		"./snapshots/guardduty")
}
