// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSnsSubscriptions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsSubscriptions(),
		"./snapshots/sns")
}

func TestIntegrationSnsTopics(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsTopics(),
		"./snapshots/sns")
}
