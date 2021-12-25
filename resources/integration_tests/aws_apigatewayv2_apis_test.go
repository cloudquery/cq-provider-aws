// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayv2ApisTest(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Apigatewayv2Apis())
}
