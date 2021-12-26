// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElasticbeanstalkApplications(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkApplications(),
		"./snapshots/elasticbeanstalk")
}
