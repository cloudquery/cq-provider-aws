package resources

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudfrontDistributions() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudfront_distributions",
		Resolver:     fetchCloudfrontDistributions,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			//todo check columns
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
				Name:     "aliases_items",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Aliases.Items"),
			},
			{
				Name: "comment",
				Type: schema.TypeString,
			},
			//{
			//	Name:     "default_cache_behaviour_target_origin_id",
			//	Type:     schema.TypeString,
			//	Resolver: schema.PathResolver("DefaultCacheBehavior.TargetOriginId"), //todo learn how to add complex object
			//},
			{
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "http_version",
				Type: schema.TypeString, //todo use type conversion for this value
			},
			{
				Name: "id",
				Type: schema.TypeString,
			},
			{
				Name: "ip_v6_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "last_modified_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "price_class",
				Type: schema.TypeString, //todo use type conversion for this value
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "web_acl_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			//todo add CacheBehaviors.Items
			//todo add CustomErrorResponses.Items
			//todo add DefaultCacheBehavior or add it to root object
			//todo add Origins
			//todo add Restrictions
			//todo add ViewerCertificate
			//todo add AliasICPRecordals
			//todo add OriginGroups
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	response, err := svc.ListDistributions(ctx, nil, func(options *cloudfront.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.DistributionList
	return nil
}
