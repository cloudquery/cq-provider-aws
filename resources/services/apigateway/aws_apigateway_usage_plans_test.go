// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayUsagePlans(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayUsagePlans(),
		"./snapshots")
}
