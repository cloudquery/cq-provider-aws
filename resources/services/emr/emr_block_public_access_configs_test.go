// +build integration

package emr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEmrBlockPublicAccessConfigs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EmrBlockPublicAccessConfigs(),
		"./snapshots")
}
