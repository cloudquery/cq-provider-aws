package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamSAMLIdentityProviders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamSamlIdentityProviders(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_saml_identity_providers",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("arn LIKE ?", fmt.Sprintf("%%/saml%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
				},
			},
		}
	})
}
