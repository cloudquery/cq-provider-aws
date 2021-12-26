// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayVpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayVpcLinks(),
		"./snapshots/apigateway")
}
