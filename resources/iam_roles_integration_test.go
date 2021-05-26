package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationIamRoles(t *testing.T) {
	awsTestIntegrationHelper(t, IamRoles(), providertest.ResourceIntegrationVerification{})
}
