// +build integration

package lambda

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationLambdaLayers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.LambdaLayers(),
		"./snapshots")
}
