package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSecretsmanagerSecrets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SecretsmanagerSecrets(),
		"./snapshots/sagemaker")
}
