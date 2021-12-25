// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEcsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsClusters())
}
