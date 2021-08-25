package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDirectAutoscalingLaunchConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t, resources.AutoscalingLaunchConfigurations(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_autoscaling_launch_configurations",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("launch_configuration_name = ?", fmt.Sprintf("%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"instance_type": "t2.nano",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_autoscaling_launch_configuration_block_device_mappings",
					ForeignKeyName: "launch_configuration_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"device_name":     "test_name",
							"ebs_volume_size": float64(5),
							"ebs_volume_type": "standard",
						},
					}},
				},
			},
		}
	})
}
