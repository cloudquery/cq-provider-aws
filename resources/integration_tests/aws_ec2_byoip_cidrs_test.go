package integration_tests

//
//func TestIntegrationEc2ByoipCidrs(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.Ec2ByoipCidrs(), []string{"aws_ec2_vpc.tf", "aws_vpc_tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: resources.Ec2ByoipCidrs().Name,
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"volume_type":          "gp2",
//						"multi_attach_enabled": false,
//						"encrypted":            false,
//						"fast_restored":        false,
//						"size":                 float64(5),
//						"tags": map[string]interface{}{
//							"Type":   "integration_test",
//							"TestId": res.Suffix,
//							"Name":   fmt.Sprintf("ebs-%s%s", res.Prefix, res.Suffix),
//						},
//					},
//				},
//			},
//		}
//	})
//}
