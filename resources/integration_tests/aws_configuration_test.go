// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationConfigConfigurationRecorders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConfigurationRecorders(),
		"./snapshots/configuration")
}

func TestIntegrationConfigConformancePack(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConformancePack(),
		"./snapshots/config")
}
