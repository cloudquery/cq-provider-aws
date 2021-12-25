// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamOpenidConnectIdentityProviders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamOpenidConnectIdentityProviders())
}
