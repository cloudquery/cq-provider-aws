// +build integration

package lambda

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationLambdaFunctions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.LambdaFunctions(),
		"./snapshots")
}
