package resources

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticbeanstalkConfigurationOptions() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_configuration_options",
		Description:  "Describes the possible values for a configuration option.",
		Resolver:     fetchElasticbeanstalkConfigurationOptions,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"application_arn", "date_created", "name", "namespace"}},
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
				Name:        "application_arn",
				Description: "The arn of the associated application.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationArn"),
			}, {
				Name:        "date_created",
				Description: "The date when the application was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name: "change_severity",
				Description: `An indication of which action is required if the value for this configuration option changes:  
				* NoInterruption : There is no interruption to the environment or application availability.  
				* RestartEnvironment : The environment is entirely restarted, all AWS resources are deleted and recreated, and the environment is unavailable during the process.  
				* RestartApplicationServer : The environment is available the entire time`,
				Type: schema.TypeString,
			},
			{
				Name:        "default_value",
				Description: "The default value for this configuration option.",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_length",
				Description: "If specified, the configuration option must be a string value no longer than this value.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "max_value",
				Description: "If specified, the configuration option must be a numeric value less than this value.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "min_value",
				Description: "If specified, the configuration option must be a numeric value greater than this value.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The name of the configuration option.",
				Type:        schema.TypeString,
			},
			{
				Name:        "namespace",
				Description: "A unique namespace identifying the option's associated AWS resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "regex_label",
				Description: "A unique name representing this regular expression.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Regex.Label"),
			},
			{
				Name:        "regex_pattern",
				Description: "The regular expression pattern that a string configuration option value with this restriction must match.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Regex.Pattern"),
			},
			{
				Name:        "user_defined",
				Description: "An indication of whether the user defined this configuration option:  * true : This configuration option was defined by the user",
				Type:        schema.TypeBool,
			},
			{
				Name:        "value_options",
				Description: "If specified, values for the configuration option are selected from this list.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "value_type",
				Description: "An indication of which type of values this option has and whether it is allowable to select one or more than one of the possible values:  * Scalar : Values for this option are a single selection from the possible values, or an unformatted string, or numeric value governed by the MIN/MAX/Regex constraints.  * List : Values for this option are multiple selections from the possible values.  * Boolean : Values for this option are either true or false .  * Json : Values for this option are a JSON representation of a ConfigDocument.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticbeanstalkConfigurationOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {

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
		configOptionsIn := elasticbeanstalk.DescribeConfigurationOptionsInput{
			ApplicationName: environment.ApplicationName,
			EnvironmentName: environment.EnvironmentName,
		}
		output, err := svc.DescribeConfigurationOptions(ctx, &configOptionsIn, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		for _, option := range output.Options {
			res <- ConfigOptions{
				option, client.GenerateResourceARN("elasticbeanstalk", "application", *environment.ApplicationName, c.Region, c.AccountID), *environment.DateCreated,
			}
		}
	}
	return nil
}

type ConfigOptions struct {
	types.ConfigurationOptionDescription
	ApplicationArn string
	DateCreated    time.Time
}
