package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	//"github.com/hashicorp/terraform-exec/tfinstall"
)

func TestIntegrationS3Buckets(t *testing.T) {
	awsTestIntegrationHelper(t, S3Buckets(), providertest.ResourceIntegrationVerification{})
}
