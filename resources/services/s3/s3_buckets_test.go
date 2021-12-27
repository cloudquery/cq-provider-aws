// +build integration

package s3

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationS3Buckets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.S3Buckets(),
		"./snapshots")
}
