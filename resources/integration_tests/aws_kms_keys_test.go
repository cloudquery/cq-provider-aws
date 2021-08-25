package integration_tests

// todo fix "created keys is not fetched"
//func TestIntegrationKmsKeys(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.KmsKeys(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_kms_keys",
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Eq{"description": fmt.Sprintf("kms-key-%s%s", res.Prefix, res.Suffix)})
//			},
//			ExpectedValues: []providertest.ExpectedValue{{
//				Count: 1,
//				Data: map[string]interface{}{
//					"description": fmt.Sprintf("kms-key-%s%s", res.Prefix, res.Suffix),
//				},
//			}},
//		}
//	})
//}
