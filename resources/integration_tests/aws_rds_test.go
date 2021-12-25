// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationRdsInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsInstances())
}

func TestIntegrationRdsSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsSubnetGroups())
}

func TestIntegrationRdsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusters())
}
