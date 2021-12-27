// +build integration

package cloudfront

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudfrontCachePolicies(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudfrontCachePolicies(),
		"./snapshots")
}
