// +build integration

package integration_tests

import (
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

// IntegrationTestsEnabledVar is the name of the environment variable that enables integration tests from this package.
// Set it to one of "1", "y", "yes", "true" to enable the tests.
const IntegrationTestsEnabledVar = "INTEGRATION_TESTS"

func firstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func awsTestIntegrationHelper(t *testing.T, table *schema.Table, resourceFiles []string) {

	cfg := `
	regions = ["us-east-1"]
	aws_debug = false
	`

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:        table,
		Config:       cfg,
		SnapshotsDir: "./snapshots",
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
