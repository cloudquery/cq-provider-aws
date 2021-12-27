// +build integration

package elasticsearch

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElasticsearchDomains(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticsearchDomains(),
		"./snapshots")
}
