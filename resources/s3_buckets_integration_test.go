//+build integration

package resources

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	//"github.com/hashicorp/terraform-exec/tfinstall"
)

func TestIntegrationS3Buckets(t *testing.T) {
	awsTestIntegrationHelper(t, S3Buckets(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_s3_buckets",
			Values: []providertest.VerificationRow{{
				"name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
			}},
		}
	})
}
