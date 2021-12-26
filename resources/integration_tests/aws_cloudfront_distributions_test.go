package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudfrontDistributions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudfrontDistributions(),
		"./snapshots/cloudfront")
}
