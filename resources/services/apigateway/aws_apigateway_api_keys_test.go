// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationApigatewayAPIKeys(t *testing.T) {
	client.AwsTestIntegrationHelper(t, ApigatewayAPIKeys(), "./snapshots")
}
