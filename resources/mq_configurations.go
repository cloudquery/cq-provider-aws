package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MqConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_mq_configurations",
		Resolver:     fetchMqConfigurations,
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
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "authentication_strategy",
				Type: schema.TypeString,
			},
			{
				Name: "created",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
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
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "latest_revision_created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LatestRevision.Created"),
			},
			{
				Name:     "latest_revision_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestRevision.Description"),
			},
			{
				Name:     "latest_revision",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LatestRevision.Revision"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMqConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config mq.ListConfigurationsInput
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for {
		response, err := svc.ListConfigurations(ctx, &config, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Configurations
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
