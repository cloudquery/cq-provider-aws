// +build integration

package codebuild

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CodebuildProjects(),
		"./snapshots/codebuild")
}
