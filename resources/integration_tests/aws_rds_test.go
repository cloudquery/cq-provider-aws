package integration_tests

// TODO - RDS certificate cannot be created

//import (
//	"fmt"
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
//)

//func TestIntegrationRdsCertificates(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.RdsCertificates(), []string{"aws_rds_instances.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_rds_certificates",
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"certificate_identifier": fmt.Sprintf("rds-ca-%s%s", res.Prefix, res.Suffix),
//					},
//				},
//			},
//		}
//	})
//}
