package integration_tests

import (
	"github.com/cloudquery/cq-provider-aws/resources"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMQBrokers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.MqBrokers(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_mq_brokers",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
				},
			},
		}
	})
}
