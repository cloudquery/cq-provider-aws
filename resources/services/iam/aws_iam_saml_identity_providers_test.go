// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamSAMLIdentityProviders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamSamlIdentityProviders(),
		"./snapshots")
}
