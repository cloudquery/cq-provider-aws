// +build integration

package cognito

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCognitoIdentityPools(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CognitoIdentityPools(),
		"./snapshots/coginto")
}
