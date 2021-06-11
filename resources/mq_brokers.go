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
				Name: "engine_type",
				Type: schema.TypeString,
			},
			{
				Name: "host_instance_type",
				Type: schema.TypeString,
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
		res <- response.BrokerSummaries
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
