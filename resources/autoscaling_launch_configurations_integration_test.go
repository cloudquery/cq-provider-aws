//+build integration

package resources

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationDirectAutoscalingLaunchConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t, AutoscalingLaunchConfigurations(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("launch_configuration_name = ?", fmt.Sprintf("%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"instance_type": "t2.micro",
				},
			}},
			//todo add block device to terraform
			//Relations: []*providertest.ResourceIntegrationVerification{
			//	{
			//		Name:           "aws_autoscaling_launch_configuration_block_device_mappings",
			//		ForeignKeyName: "select * from aws_autoscaling_launch_c",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"metric_id": "m1",
			//			},
			//		}},
			//	},
			//},
		}
	})
}
