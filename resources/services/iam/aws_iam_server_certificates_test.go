// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamServerCertificates(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamServerCertificates(),
		"./snapshots")
}
