package resources

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
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
			// Restrictions start
			{
				Name:     "restrictions_geo_restriction_restriction_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Restrictions.GeoRestriction.RestrictionType"),
			},
			{
				Name:     "restrictions_geo_restriction_restriction_items",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Restrictions.GeoRestriction.Items"),
			},
			// Restrictions End
			{
				Name: "status",
				Type: schema.TypeString,
			},
			//ViewerCertificate start
			{
				Name:     "viewer_certificate_acm_certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.ACMCertificateArn"),
			},
			{
				Name:     "viewer_certificate_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.Certificate"),
			},
			{
				Name:     "viewer_certificate_certificate_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.CertificateSource"),
			},
			{
				Name:     "viewer_certificate_cloudfront_default_certificate",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ViewerCertificate.CloudFrontDefaultCertificate"),
			},
			{
				Name:     "viewer_certificate_iam_certificate_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.IAMCertificateId"),
			},
			{
				Name:     "viewer_certificate_minimum_protocol_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.MinimumProtocolVersion"),
			},
			{
				Name:     "viewer_certificate_ssl_support_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.SSLSupportMethod"),
			},
			//ViewerCertificate end
			{
				Name: "web_acl_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			//todo learn how to add DefaultCacheBehavior.LambdaFunctionsAssociations
			{
				Name:     "aws_cloudfront_distribution_cache_behaviour",
				Resolver: fetchCloudfrontCacheBehaviours,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "path_pattern",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PathPattern"),
					},
					{
						Name:     "target_origin_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TargetOriginId"),
					},
					{
						Name:     "viewer_protocol_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ViewerProtocolPolicy"),
					},
					{
						Name:     "allowed_methods",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("AllowedMethods.Items"),
					},
					{
						Name:     "cached_methods",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("AllowedMethods.CachedMethods.Items"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_custom_error_response",
				Resolver: fetchCloudfrontCustomErrorResponses,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "error_code",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ErrorCode"),
					},
					{
						Name:     "error_caching_min_ttl",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ErrorCachingMinTTL"),
					},
					{
						Name:     "response_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResponseCode"),
					},
					{
						Name:     "response_page_path",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResponsePagePath"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_origin",
				Resolver: fetchCloudfrontOrigins,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "domain_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DomainName"),
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name:     "connection_attempts",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ConnectionAttempts"),
					},
					{
						Name:     "connection_timeout",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ConnectionTimeout"),
					},
					{
						Name:     "custom_headers",
						Type:     schema.TypeJSON,
						Resolver: resolveCloudfrontOriginCustomHeaders,
					},
					{
						Name:     "custom_origin_config_http_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPPort"),
					},
					{
						Name:     "custom_origin_config_https_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPSPort"),
					},
					{
						Name:     "custom_origin_config_origin_protocol_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginProtocolPolicy"),
					},
					{
						Name:     "custom_origin_config_origin_path",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginPath"),
					},
					{
						Name:     "custom_origin_config_origin_shield_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginShield.Enabled"),
					},
					{
						Name:     "custom_origin_config_origin_shield_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginShield.OriginShieldRegion"),
					},
					{
						Name:     "custom_origin_config_s3_origin_config_origin_access_identity",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.S3OriginConfig.OriginAccessIdentity"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_viewer_certificate",
				Resolver: fetchCloudfrontViewerCertificates,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "domain_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DomainName"),
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name:     "connection_attempts",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ConnectionAttempts"),
					},
					{
						Name:     "connection_timeout",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("ConnectionTimeout"),
					},
					{
						Name:     "custom_headers",
						Type:     schema.TypeJSON,
						Resolver: resolveCloudfrontOriginCustomHeaders,
					},
					{
						Name:     "custom_origin_config_http_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPPort"),
					},
					{
						Name:     "custom_origin_config_https_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPSPort"),
					},
					{
						Name:     "custom_origin_config_origin_protocol_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginProtocolPolicy"),
					},
					{
						Name:     "custom_origin_config_origin_path",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginPath"),
					},
					{
						Name:     "custom_origin_config_origin_shield_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginShield.Enabled"),
					},
					{
						Name:     "custom_origin_config_origin_shield_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginShield.OriginShieldRegion"),
					},
					{
						Name:     "custom_origin_config_s3_origin_config_origin_access_identity",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.S3OriginConfig.OriginAccessIdentity"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_alias_icp_recordal",
				Resolver: fetchCloudfrontAliasICPRecordals,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "cname",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CNAME"),
					},
					{
						Name:     "icp_recordal_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ICPRecordalStatus"), //todo remove redundant path resolvers
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_origin_group",
				Resolver: fetchCloudfrontOriginGroups,
				Columns: []schema.Column{
					{
						Name:     "distribution_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "failover_criteria_statuscodes_items",
						Type:     schema.TypeIntArray,
						Resolver: schema.PathResolver("FailoverCriteria.StatusCodes.Items"),
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"), //todo remove redundant path resolvers
					},
					{
						Name:     "members_origin_ids",
						Type:     schema.TypeStringArray,
						Resolver: resolveCloudfrontOriginGroupMembers,
					},
				},
			},
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

func fetchCloudfrontCacheBehaviours(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.CacheBehaviors.Items
	return nil
}

func fetchCloudfrontCustomErrorResponses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.CustomErrorResponses.Items
	return nil
}

func resolveCloudfrontOriginCustomHeaders(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Origin)
	tags := map[string]*string{}
	for _, t := range r.CustomHeaders.Items {
		tags[*t.HeaderName] = t.HeaderValue
	}
	resource.Set("custom_headers", tags)
	return nil
}

func fetchCloudfrontOrigins(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.Origins.Items
	return nil
}

func fetchCloudfrontViewerCertificates(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.ViewerCertificate
	return nil
}

func fetchCloudfrontAliasICPRecordals(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.AliasICPRecordals
	return nil
}

func fetchCloudfrontOriginGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- instance.OriginGroups.Items
	return nil
}

func resolveCloudfrontOriginGroupMembers(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	members := make([]string, 0, *r.Members.Quantity)
	for _, t := range r.Members.Items {
		members = append(members, *t.OriginId)
	}
	resource.Set("members_origin_ids", members)
	return nil
}
