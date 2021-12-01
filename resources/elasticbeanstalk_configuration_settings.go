package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticbeanstalkConfigurationSettings() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_configuration_settings",
		Description:  "Describes the settings for a configuration set.",
		Resolver:     fetchElasticbeanstalkConfigurationSettings,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"application_name", "date_created"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "application_name",
				Description: "The name of the application associated with this configuration set.",
				Type:        schema.TypeString,
			},
			{
				Name:        "date_created",
				Description: "The date (in UTC time) when this configuration set was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "date_updated",
				Description: "The date (in UTC time) when this configuration set was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deployment_status",
				Description: "If this configuration set is associated with an environment, the DeploymentStatus parameter indicates the deployment status of this configuration set:  * null: This configuration is not associated with a running environment.  * pending: This is a draft configuration that is not deployed to the associated environment but is in the process of deploying.  * deployed: This is the configuration that is currently deployed to the associated running environment.  * failed: This is a draft configuration that failed to successfully deploy.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "Describes this configuration set.",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment_name",
				Description: "If not null, the name of the environment for this configuration set.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_arn",
				Description: "The ARN of the platform version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "solution_stack_name",
				Description: "The name of the solution stack this configuration set uses.",
				Type:        schema.TypeString,
			},
			{
				Name:        "template_name",
				Description: "If not null, the name of the configuration template for this configuration set.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticbeanstalk_configuration_setting_option_settings",
				Description: "A specification identifying an individual configuration option along with its current value",
				Resolver:    fetchElasticbeanstalkConfigurationSettingOptionSettings,
				Columns: []schema.Column{
					{
						Name:        "configuration_setting_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticbeanstalk_configuration_settings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "namespace",
						Description: "A unique namespace that identifies the option's associated AWS resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "option_name",
						Description: "The name of the configuration option.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_name",
						Description: "A unique resource name for the option setting",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The current value for the configuration option.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticbeanstalkConfigurationSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	var config elasticbeanstalk.DescribeEnvironmentsInput
	allEnvs := make([]types.EnvironmentDescription, 0)
	for {
		response, err := svc.DescribeEnvironments(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		allEnvs = append(allEnvs, response.Environments...)
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	for _, environment := range allEnvs {
		configOptionsIn := elasticbeanstalk.DescribeConfigurationSettingsInput{
			ApplicationName: environment.ApplicationName,
			EnvironmentName: environment.EnvironmentName,
		}
		output, err := svc.DescribeConfigurationSettings(ctx, &configOptionsIn, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		for _, option := range output.ConfigurationSettings {
			res <- option
		}
	}
	return nil
}

func fetchElasticbeanstalkConfigurationSettingOptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	option, ok := parent.Item.(types.ConfigurationSettingsDescription)
	if !ok {
		meta.Logger().Error("parent.Item", "Item", parent.Item)
		return fmt.Errorf("not types.ConfigurationOptionSetting")
	}
	for _, t := range option.OptionSettings {
		res <- t
	}

	return nil
}
