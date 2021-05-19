package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayVpcLinks() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_vpc_links",
		Resolver:     fetchApigatewayVpcLinks,
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
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "security_group_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "subnet_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "vpc_link_id",
				Type: schema.TypeString,
			},
			{
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "vpc_link_status",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_link_status_message",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_link_version",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayVpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
