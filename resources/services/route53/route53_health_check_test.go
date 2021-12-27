// +build integration

package route53

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRoute53HealthChecks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Route53HealthChecks(),
		"./snapshots")
}
