// +build integration

package config

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationConfigConfigurationRecorders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConfigurationRecorders(),
		"./snapshots")
}

func TestIntegrationConfigConformancePack(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConformancePack(),
		"./snapshots")
}
