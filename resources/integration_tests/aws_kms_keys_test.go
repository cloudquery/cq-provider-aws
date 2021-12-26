package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationKmsKeys(t *testing.T) {
	awsTestIntegrationHelper(t, resources.KmsKeys(),
		"./snapshots/kms")
}
