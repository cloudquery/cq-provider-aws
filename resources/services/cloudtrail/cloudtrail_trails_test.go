// +build integration

package cloudtrail

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCloudtrailTrails(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudtrailTrails(),
		"./snapshots")
}
