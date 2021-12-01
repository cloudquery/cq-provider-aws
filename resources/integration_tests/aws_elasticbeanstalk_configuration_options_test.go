package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElasticbeanstalkConfigurationOptions(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkConfigurationOptions(), []string{"aws_elasticbeanstalk_environments.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elasticbeanstalk_configuration_options",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{

					"application_arn": fmt.Sprintf("arn:aws:elasticbeanstalk:%%:%%:application/beanstalk-ea-%s", res.Suffix),
					"name":            "Protocol",
					"namespace":       "aws:elbv2:listener",
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"value_options": []interface{}{"HTTP", "HTTPS"},
					},
				},
			},
		}
	})
}
