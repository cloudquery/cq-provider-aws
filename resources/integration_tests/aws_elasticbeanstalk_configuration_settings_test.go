package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElasticbeanstalkConfigurationSettings(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkConfigurationSettings(), []string{"aws_elasticbeanstalk_environments.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elasticbeanstalk_configuration_settings",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"application_name": firstN(fmt.Sprintf("beanstalk-ea-%s", res.Suffix), 40)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"application_name":    fmt.Sprintf("beanstalk-ea-%s", res.Suffix),
						"deployment_status":   "deployed",
						"environment_name":    fmt.Sprintf("beanstalk-ee-%s", res.Suffix),
						"platform_arn":        "arn:aws:elasticbeanstalk:us-east-1::platform/Go 1 running on 64bit Amazon Linux 2/3.3.4",
						"solution_stack_name": "64bit Amazon Linux 2 v3.3.4 running Go 1",
						"template_name":       nil,
					},
				},
			},
		}
	})
}
