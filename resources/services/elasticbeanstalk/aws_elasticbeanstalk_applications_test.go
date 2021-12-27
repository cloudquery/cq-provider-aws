// +build integration

package elasticbeanstalk

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElasticbeanstalkApplications(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkApplications(),
		"./snapshots")
}
