package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationIamOpenidConnectIdentityProviders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamOpenidConnectIdentityProviders(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_openid_connect_identity_providers",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"url": fmt.Sprintf("openidprovider%s.com", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"url": fmt.Sprintf("openidprovider%s.com", res.Suffix),
					},
				},
			},
		}
	})
}
