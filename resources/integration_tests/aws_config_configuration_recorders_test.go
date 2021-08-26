package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationConfigConfigurationRecorders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConfigurationRecorders(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_config_configuration_recorders",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                          "config_recorder",
					"recording_group_all_supported": true,
					"recording_group_include_global_resource_types": false,
					"recording_group_resource_types":                []interface{}{},
				},
			}},
		}
	})
}
