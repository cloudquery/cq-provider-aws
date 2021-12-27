// +build integration

package secretsmanager

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSecretsmanagerSecrets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SecretsmanagerSecrets(),
		"./snapshots")
}
