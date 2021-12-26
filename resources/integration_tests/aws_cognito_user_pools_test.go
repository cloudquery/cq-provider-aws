package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCognitoUserPools(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CognitoUserPools(),
		"./snapshots/coginto")
}
