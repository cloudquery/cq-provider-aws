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

func ElasticbeanstalkApplications() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_applications",
		Description:  "Describes the properties of an application.",
		Resolver:     fetchElasticbeanstalkApplications,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationArn"),
			},
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationName"),
			},
			{
				Name:        "configuration_templates",
				Description: "The names of the configuration templates associated with this application.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "date_created",
				Description: "The date when the application was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "date_updated",
				Description: "The date when the application was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "User-defined description of the application.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_lifecycle_config_service_role",
				Description: "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a VersionLifecycleConfig for the application in one of the supporting calls (CreateApplication or UpdateApplicationResourceLifecycle)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.ServiceRole"),
			},
			{
				Name:        "lifecycle_config_max_age_rule_enabled",
				Description: "Specify true to apply the rule, or false to disable it.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.Enabled"),
			},
			{
				Name:        "lifecycle_config_max_age_rule_delete_source_from_s3",
				Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.DeleteSourceFromS3"),
			},
			{
				Name:        "lifecycle_config_max_age_rule_max_age_in_days",
				Description: "Specify the number of days to retain an application versions.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.MaxAgeInDays"),
			},
			{
				Name:        "lifecycle_config_max_count_rule_enabled",
				Description: "Specify true to apply the rule, or false to disable it.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.Enabled"),
			},
			{
				Name:        "lifecycle_config_max_count_rule_delete_source_from_s3",
				Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.DeleteSourceFromS3"),
			},
			{
				Name:        "lifecycle_config_max_count_rule_max_count",
				Description: "Specify the maximum number of application versions to retain.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.MaxCount"),
			},
			{
				Name:        "versions",
				Description: "The names of the versions for this application.",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticbeanstalk_environments",
				Description: "Describes the properties of an environment.",
				Resolver:    fetchElasticbeanstalkEnvironments,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},

				Relations: []*schema.Table{
					{
						Name:        "aws_elasticbeanstalk_environment_links",
						Description: "A link to another environment, defined in the environment's manifest",
						Resolver:    fetchElasticbeanstalkEnvironmentLinks,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"environment_cq_id", "link_name"}},
						Columns: []schema.Column{
							{
								Name:        "environment_cq_id",
								Description: "Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "environment_name",
								Description: "The name of the linked environment (the dependency).",
								Type:        schema.TypeString,
							},
							{
								Name:        "link_name",
								Description: "The name of the link.",
								Type:        schema.TypeString,
							},
						},
					},
				},
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
						Name:        "tags",
						Type:        schema.TypeJSON,
						Description: "Any tags assigned to the resource",
						Resolver:    resolveElasticbeanstalkEnvironmentTags,
					},
					{
						Name:        "abortable_operation_in_progress",
						Description: "Indicates if there is an in-progress environment configuration update or application version deployment that you can cancel",
						Type:        schema.TypeBool,
					},
					{
						Name:        "application_name",
						Description: "The name of the application associated with this environment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cname",
						Description: "The URL to the CNAME for this environment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CNAME"),
					},
					{
						Name:        "date_created",
						Description: "The creation date for this environment.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "date_updated",
						Description: "The last modified date for this environment.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "Describes this environment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_url",
						Description: "For load-balanced, autoscaling environments, the URL to the LoadBalancer",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointURL"),
					},
					{
						Name:        "arn",
						Description: "The environment's Amazon Resource Name (ARN), which can be used in other API requests that require an ARN.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EnvironmentArn"),
					},
					{
						Name:        "id",
						Description: "The ID of this environment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EnvironmentId"),
					},
					{
						Name:        "name",
						Description: "The name of this environment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EnvironmentName"),
					},
					{
						Name:        "health",
						Description: "Describes the health status of the environment",
						Type:        schema.TypeString,
					},
					{
						Name:        "health_status",
						Description: "Returns the health status of the application running in your environment",
						Type:        schema.TypeString,
					},
					{
						Name:        "operations_role",
						Description: "The Amazon Resource Name (ARN) of the environment's operations role",
						Type:        schema.TypeString,
					},
					{
						Name:        "platform_arn",
						Description: "The ARN of the platform version.",
						Type:        schema.TypeString,
					},
					{
						Name:        "load_balancer_domain",
						Description: "The domain name of the LoadBalancer.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Resources.LoadBalancer.Domain"),
					},
					{
						Name:        "listeners",
						Description: "A list of Listeners used by the LoadBalancer.",
						Type:        schema.TypeJSON,
						Resolver:    resolveElasticbeanstalkEnvironmentListeners,
					},
					{
						Name:        "load_balancer_name",
						Description: "The name of the LoadBalancer.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Resources.LoadBalancer.LoadBalancerName"),
					},
					{
						Name:        "solution_stack_name",
						Description: "The name of the SolutionStack deployed with this environment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The current operational status of the environment:  * Launching: Environment is in the process of initial deployment.  * Updating: Environment is in the process of updating its configuration settings or application version.  * Ready: Environment is available to have an action performed on it, such as update or terminate.  * Terminating: Environment is in the shut-down process.  * Terminated: Environment is not running.",
						Type:        schema.TypeString,
					},
					{
						Name:        "template_name",
						Description: "The name of the configuration template used to originally launch this environment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tier_name",
						Description: "The name of this environment tier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Tier.Name"),
					},
					{
						Name:        "tier_type",
						Description: "The type of this environment tier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Tier.Type"),
					},
					{
						Name:        "tier_version",
						Description: "The version of this environment tier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Tier.Version"),
					},
					{
						Name:        "version_label",
						Description: "The application version deployed in this environment.",
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
func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.ApplicationDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", parent.Item)
	}
	var config elasticbeanstalk.DescribeEnvironmentsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	config.ApplicationName = aws.String(*p.ApplicationName)
	for {
		response, err := svc.DescribeEnvironments(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Environments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveElasticbeanstalkEnvironmentTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", resource.Item)
	}
	if p.Resources == nil || p.Resources.LoadBalancer == nil {
		return nil
	}
	listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
	for _, l := range p.Resources.LoadBalancer.Listeners {
		listeners[l.Port] = l.Protocol
	}
	return resource.Set(c.Name, listeners)
}
func resolveElasticbeanstalkEnvironmentListeners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", resource.Item)
	}
	svc := meta.(*client.Client).Services().ElasticBeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.EnvironmentArn,
	}, func(o *elasticbeanstalk.Options) {})
	if err != nil {
		return err
	}
	if len(tagsOutput.ResourceTags) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.ResourceTags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
func fetchElasticbeanstalkEnvironmentLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", parent.Item)
	}
	res <- p.EnvironmentLinks
	return nil
}

func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	output, err := svc.DescribeApplications(ctx, &config, func(options *elasticbeanstalk.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	fmt.Printf("application data %+v", output.Applications)
	res <- output.Applications
	return nil
}
