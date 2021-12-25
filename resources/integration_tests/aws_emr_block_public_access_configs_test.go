// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEmrBlockPublicAccessConfigs(t *testing.T) {
	table := resources.EmrBlockPublicAccessConfigs()
	awsTestIntegrationHelper(t, table)
}
