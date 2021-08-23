package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/Masterminds/squirrel"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCognitoUserPools(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CognitoUserPools(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_cognito_user_pools",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("cognito_user_pool%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("cognito_user_pool%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
		}
	})
}
