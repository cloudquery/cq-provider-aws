package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationS3Account(t *testing.T) {
	awsTestIntegrationHelper(t, resources.S3Accounts(),
		"./snapshots/s3")
}
