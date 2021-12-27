// +build integration

package elasticbeanstalk

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationElasticbeanstalkEnvironments(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkEnvironments(),
		"./snapshots/elasticbeanstalk")
}
