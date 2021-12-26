// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayAPIKeys(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayAPIKeys(), "./snapshots/apigateway")
}
