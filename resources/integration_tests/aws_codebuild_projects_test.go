// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	resource := resources.CodebuildProjects()
	awsTestIntegrationHelper(t, resource)
}
