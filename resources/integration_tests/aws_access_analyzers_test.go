// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationAccessAnalyzers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.AccessAnalyzerAnalyzer(), "./snapshots/access_analyzers")
}
