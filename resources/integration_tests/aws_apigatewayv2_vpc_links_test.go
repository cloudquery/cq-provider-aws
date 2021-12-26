package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayv2VpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Apigatewayv2VpcLinks(),
		"./snapshots/apigatewayv2")
}
