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
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "authentication_strategy",
				Type: schema.TypeString,
			},
			{
				Name: "auto_minor_version_upgrade",
				Type: schema.TypeBool,
			},
			{
				Name: "broker_arn",
				Type: schema.TypeString,
			},
			{
				Name: "broker_id",
				Type: schema.TypeString,
			},
			{
				Name: "broker_name",
				Type: schema.TypeString,
			},
			{
				Name: "broker_state",
				Type: schema.TypeString,
			},
			{
				Name: "created",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "deployment_mode",
				Type: schema.TypeString,
			},
			{
				Name:     "encryption_options_use_aws_owned_key",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EncryptionOptions.UseAwsOwnedKey"),
			},
			{
				Name:     "encryption_options_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionOptions.KmsKeyId"),
			},
			{
				Name: "engine_type",
				Type: schema.TypeString,
			},
			{
				Name: "engine_version",
				Type: schema.TypeString,
			},
			{
				Name: "host_instance_type",
				Type: schema.TypeString,
			},
			{
				Name: "publicly_accessible",
				Type: schema.TypeBool,
			},
			{
				Name: "security_groups",
				Type: schema.TypeStringArray,
			},
			{
				Name: "storage_type",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "aws_mq_broker_users",
				Resolver:     fetchMqBrokerUsers,
				Multiplex:    client.AccountRegionMultiplex,
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:     "broker_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "account_id",
						Type:     schema.TypeString,
						Resolver: client.ResolveAWSAccount,
					},
					{
						Name:     "region",
						Type:     schema.TypeString,
						Resolver: client.ResolveAWSRegion,
					},
					{
						Name: "console_access",
						Type: schema.TypeBool,
					},
					{
						Name: "groups",
						Type: schema.TypeStringArray,
					},
					{
						Name: "username",
						Type: schema.TypeString,
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
