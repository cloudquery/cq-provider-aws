// +build integration

package sns

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSnsSubscriptions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsSubscriptions(),
		"./snapshots")
}

func TestIntegrationSnsTopics(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SnsTopics(),
		"./snapshots")
}
