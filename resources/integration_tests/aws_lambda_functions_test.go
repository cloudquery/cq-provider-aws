package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationLambdaFunctions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.LambdaFunctions(),
		"./snapshots/lambda")
}
