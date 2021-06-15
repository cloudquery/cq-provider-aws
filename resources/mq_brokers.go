package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MqBrokers() *schema.Table {
	return &schema.Table{
		Name:         "aws_mq_brokers",
		Resolver:     fetchMqBrokers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "authentication_strategy",
				Description: "The authentication strategy used to secure the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "broker_arn",
				Description: "The Amazon Resource Name (ARN) of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_id",
				Description: "The unique ID that Amazon MQ generates for the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_name",
				Description: "The name of the broker",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_state",
				Description: "The status of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "The time when the broker was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deployment_mode",
				Description: "The deployment mode of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_options_use_aws_owned_key",
				Description: "Enables the use of an AWS owned CMK using AWS Key Management Service (KMS).",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EncryptionOptions.UseAwsOwnedKey"),
			},
			{
				Name:        "encryption_options_kms_key_id",
				Description: "The symmetric customer master key (CMK) to use for the AWS Key Management Service (KMS).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionOptions.KmsKeyId"),
			},
			{
				Name:        "engine_type",
				Description: "The type of broker engine.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The version of the broker engine",
				Type:        schema.TypeString,
			},
			{
				Name:        "host_instance_type",
				Description: "The broker's instance type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "publicly_accessible",
				Description: "Enables connections from applications outside of the VPC that hosts the broker's subnets.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "security_groups",
				Description: "The list of security groups (1 minimum, 5 maximum) that authorizes connections to brokers.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "storage_type",
				Description: "The broker's storage type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The list of all tags associated with this broker.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "aws_mq_broker_configurations",
				Resolver:     fetchMqBrokerConfigurations,
				Multiplex:    client.AccountRegionMultiplex,
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:        "broker_id",
						Description: "Unique ID of aws_mq_brokers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
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
						Description: "The ARN of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "authentication_strategy",
						Description: "The authentication strategy associated with the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created",
						Description: "The date and time of the configuration revision.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "engine_type",
						Description: "The type of broker engine.",
						Type:        schema.TypeString,
					},
					{
						Name:        "engine_version",
						Description: "The version of the broker engine.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_id",
						Description: "The unique ID that Amazon MQ generates for the configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "latest_revision_created",
						Description: "The date and time of the configuration revision.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("LatestRevision.Created"),
					},
					{
						Name:        "latest_revision_description",
						Description: "The description of the configuration revision.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestRevision.Description"),
					},
					{
						Name:        "latest_revision",
						Description: "The revision number of the configuration.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LatestRevision.Revision"),
					},
					{
						Name:        "name",
						Description: "The name of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The list of all tags associated with this configuration.",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:         "aws_mq_broker_users",
				Resolver:     fetchMqBrokerUsers,
				Multiplex:    client.AccountRegionMultiplex,
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:        "broker_id",
						Description: "Unique ID of aws_mq_brokers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
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
						Name:        "console_access",
						Description: "Enables access to the the ActiveMQ Web Console for the ActiveMQ user.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "groups",
						Description: "The list of groups (20 maximum) to which the ActiveMQ user belongs",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "username",
						Description: "The username of the ActiveMQ user.",
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
func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config mq.ListBrokersInput
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for {
		response, err := svc.ListBrokers(ctx, &config, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, bs := range response.BrokerSummaries {
			output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, func(options *mq.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- output
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchMqBrokerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for _, cfg := range broker.Configurations.History {
		input := mq.DescribeConfigurationInput{ConfigurationId: cfg.Id}
		output, err := svc.DescribeConfiguration(ctx, &input, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for _, us := range broker.Users {
		input := mq.DescribeUserInput{
			BrokerId: broker.BrokerId,
			Username: us.Username,
		}
		output, err := svc.DescribeUser(ctx, &input, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}
