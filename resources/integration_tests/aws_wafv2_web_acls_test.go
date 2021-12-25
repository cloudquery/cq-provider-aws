// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationWAFv2WebACLs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Wafv2WebAcls())
}
