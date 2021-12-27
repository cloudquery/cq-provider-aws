// +build integration

package guardduty

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationGuarddutyDetectors(t *testing.T) {
	awsTestIntegrationHelper(t, resources.GuarddutyDetectors(),
		"./snapshots")
}
