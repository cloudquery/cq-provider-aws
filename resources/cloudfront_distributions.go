package resources

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
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

			// DefaultCacheBehavior start
			{
				Name:     "default_cache_behaviour_target_origin_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TargetOriginId"),
			},
			{
				Name:     "default_cache_behaviour_viewer_protocol_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ViewerProtocolPolicy"),
			},
			{
				Name:     "default_cache_behaviour_allowed_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.AllowedMethods.Items"),
			},
			{
				Name:     "default_cache_behaviour_allowed_methods_cached_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.AllowedMethods.CachedMethods.Items"),
			},
			{
				Name:     "default_cache_behaviour_cache_policy_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.CachePolicyId"),
			},
			{
				Name:     "default_cache_behaviour_compress",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.Compress"),
			},
			{
				Name:     "default_cache_behaviour_default_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.DefaultTTL"),
			},
			{
				Name:     "default_cache_behaviour_field_level_encryption_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:     "default_cache_behaviour_field_level_encryption_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:     "default_cache_behaviour_field_level_encryption_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_cookies_forward",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Cookies.Forward"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_cookies_forward",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Cookies.Forward"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_cookies_white_listed_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames.Items"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_query_string",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.QueryString"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Headers.Items"),
			},
			{
				Name:     "default_cache_behaviour_forwarded_values_query_string_cache_keys",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys.Items"),
			},
			{
				Name:     "default_cache_behaviour_max_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.MaxTTL"),
			},
			{
				Name:     "default_cache_behaviour_min_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.MinTTL"),
			},
			{
				Name:     "default_cache_behaviour_origin_request_policy_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.OriginRequestPolicyId"),
			},
			{
				Name:     "default_cache_behaviour_realtime_log_config_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.RealtimeLogConfigArn"),
			},
			{
				Name:     "default_cache_behaviour_smooth_streaming",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.SmoothStreaming"),
			},
			{
				Name:     "default_cache_behaviour_trusted_key_groups_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedKeyGroups.Enabled"),
			},
			{
				Name:     "default_cache_behaviour_trusted_key_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedKeyGroups.Items"),
			},
			{
				Name:     "default_cache_behaviour_trusted_signers_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedSigners.Enabled"),
			},
			{
				Name:     "default_cache_behaviour_trusted_signers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedSigners.Items"),
			},
			// DefaultCacheBehavior end

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
				Type: schema.TypeString,
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
				Type: schema.TypeString,
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
			//todo learn how to add DefaultCacheBehavior.LambdaFunctionsAssociations
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
	var config cloudfront.ListDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	for {
		response, err := svc.ListDistributions(ctx, nil, func(options *cloudfront.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.DistributionList.Items
		if aws.ToString(response.DistributionList.Marker) == "" {
			break
		}
		config.Marker = response.DistributionList.Marker
	}
	return nil
}
