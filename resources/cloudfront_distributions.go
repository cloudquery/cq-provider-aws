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
		Description:  "A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.",
		Resolver:     fetchCloudfrontDistributions,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				//+
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudfrontDistributionTags,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) for the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "caller_reference",
				Description: "A unique value (for example, a date-time stamp) that ensures that the request can't be replayed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.CallerReference"),
			},
			{
				Name:        "comment",
				Description: "Any comments you want to include about the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Comment"),
			},
			{
				Name:        "cache_behavior_target_origin_id",
				Description: "The value of ID for the origin that you want CloudFront to route requests to when they use the default cache behavior.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TargetOriginId"),
			},
			{
				Name:        "cache_behavior_viewer_protocol_policy",
				Description: "The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ViewerProtocolPolicy"),
			},
			{
				Name:        "cache_behavior_allowed_methods_items",
				Description: "A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.  This member is required.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.Items"),
			},
			{
				Name:        "cache_behavior_allowed_methods_quantity",
				Description: "The number of HTTP methods that you want CloudFront to forward to your origin. Valid values are 2 (for GET and HEAD requests), 3 (for GET, HEAD, and OPTIONS requests) and 7 (for GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests).  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.Quantity"),
			},
			{
				Name:        "cache_behavior_allowed_methods_cached_methods_items",
				Description: "A complex type that contains the HTTP methods that you want CloudFront to cache responses to.  This member is required.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.CachedMethods.Items"),
			},
			{
				Name:        "cache_behavior_allowed_methods_cached_methods_quantity",
				Description: "The number of HTTP methods for which you want CloudFront to cache responses. Valid values are 2 (for caching responses to GET and HEAD requests) and 3 (for caching responses to GET, HEAD, and OPTIONS requests).  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.CachedMethods.Quantity"),
			},
			{
				Name:        "behavior_target_origin_id",
				Description: "The unique identifier of the cache policy that is attached to the default cache behavior",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.CachePolicyId"),
			},
			{
				Name:        "cache_behavior_compress",
				Description: "Whether you want CloudFront to automatically compress certain files for this cache behavior",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.Compress"),
			},
			{
				Name:        "cache_behavior_default_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.DefaultTTL"),
			},
			{
				Name:        "cache_behavior_field_level_encryption_id",
				Description: "The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:        "cache_behavior_forwarded_values_cookies_forward",
				Description: "This field is deprecated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.Forward"),
			},
			{
				Name:        "cache_behavior_forwarded_values_cookies_whitelisted_names_quantity",
				Description: "The number of cookie names in the Items list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames.Quantity"),
			},
			{
				Name:        "cache_behavior_forwarded_values_cookies_whitelisted_names_items",
				Description: "A list of cookie names.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames.Items"),
			},
			{
				Name:        "cache_behavior_forwarded_values_query_string",
				Description: "This field is deprecated",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryString"),
			},
			{
				Name:        "cache_behavior_forwarded_values_headers_quantity",
				Description: "The number of header names in the Items list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Headers.Quantity"),
			},
			{
				Name:        "cache_behavior_forwarded_values_headers_items",
				Description: "A list of HTTP header names.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Headers.Items"),
			},
			{
				Name:        "cache_behavior_forwarded_values_query_string_cache_keys_quantity",
				Description: "The number of whitelisted query string parameters for a cache behavior.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys.Quantity"),
			},
			{
				Name:        "cache_behavior_forwarded_values_query_string_cache_keys_items",
				Description: "A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys.Items"),
			},
			{
				Name:        "cache_behavior_lambda_function_associations_quantity",
				Description: "The number of Lambda function associations for this cache behavior.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.LambdaFunctionAssociations.Quantity"),
			},
			{
				Name:        "cache_behavior_max_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.MaxTTL"),
			},
			{
				Name:        "cache_behavior_min_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.MinTTL"),
			},
			{
				Name:        "cache_behavior_origin_request_policy_id",
				Description: "The unique identifier of the origin request policy that is attached to the default cache behavior",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.OriginRequestPolicyId"),
			},
			{
				Name:        "cache_behavior_realtime_log_config_arn",
				Description: "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.RealtimeLogConfigArn"),
			},
			{
				Name:        "cache_behavior_smooth_streaming",
				Description: "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.SmoothStreaming"),
			},
			{
				Name:        "cache_behavior_trusted_key_groups_enabled",
				Description: "This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedKeyGroups.Enabled"),
			},
			{
				Name:        "cache_behavior_trusted_key_groups_quantity",
				Description: "The number of key groups in the list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedKeyGroups.Quantity"),
			},
			{
				Name:        "cache_behavior_trusted_key_groups_items",
				Description: "A list of key groups identifiers.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedKeyGroups.Items"),
			},
			{
				Name:        "cache_behavior_trusted_signers_enabled",
				Description: "This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedSigners.Enabled"),
			},
			{
				Name:        "cache_behavior_trusted_signers_quantity",
				Description: "The number of AWS accounts in the list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedSigners.Quantity"),
			},
			{
				Name:        "cache_behavior_trusted_signers_items",
				Description: "A list of AWS account identifiers.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedSigners.Items"),
			},
			{
				Name:        "enabled",
				Description: "From this field, you can enable or disable the selected distribution.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Enabled"),
			},
			{
				Name:        "origins_quantity",
				Description: "The number of origins for this distribution.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.Origins.Quantity"),
			},
			{
				Name:        "aliases_quantity",
				Description: "The number of CNAME aliases, if any, that you want to associate with this distribution.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.Aliases.Quantity"),
			},
			{
				Name:        "aliases_items",
				Description: "A complex type that contains the CNAME aliases, if any, that you want to associate with this distribution.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.Aliases.Items"),
			},
			{
				Name:        "cache_behaviors_quantity",
				Description: "The number of cache behaviors for this distribution.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.CacheBehaviors.Quantity"),
			},
			{
				Name:        "custom_error_responses_quantity",
				Description: "The number of HTTP status codes for which you want to specify a custom error page and/or a caching duration",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.CustomErrorResponses.Quantity"),
			},
			{
				Name:        "default_root_object",
				Description: "The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution (http://www.example.com) instead of an object in your distribution (http://www.example.com/product-description.html)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultRootObject"),
			},
			{
				Name:        "http_version",
				Description: "(Optional) Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.HttpVersion"),
			},
			{
				Name:        "is_ipv6_enabled",
				Description: "If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.IsIPV6Enabled"),
			},
			{
				Name:        "logging_bucket",
				Description: "The Amazon S3 bucket to store the access logs in, for example, myawslogbucket.s3.amazonaws.com.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Bucket"),
			},
			{
				Name:        "logging_enabled",
				Description: "Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Enabled"),
			},
			{
				Name:        "logging_include_cookies",
				Description: "Specifies whether you want CloudFront to include cookies in access logs, specify true for IncludeCookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.IncludeCookies"),
			},
			{
				Name:        "logging_prefix",
				Description: "An optional string that you want CloudFront to prefix to the access log filenames for this distribution, for example, myprefix/",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Prefix"),
			},
			{
				Name:        "origin_groups_quantity",
				Description: "The number of origin groups.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.OriginGroups.Quantity"),
			},
			{
				Name:        "price_class",
				Description: "The price class that corresponds with the maximum price that you want to pay for CloudFront service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.PriceClass"),
			},
			{
				Name:        "restrictions_geo_restriction_quantity",
				Description: "When geo restriction is enabled, this is the number of countries in your whitelist or blacklist",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DistributionConfig.Restrictions.GeoRestriction.Quantity"),
			},
			{
				Name:        "restrictions_geo_restriction_restriction_type",
				Description: "The method that you want to use to restrict distribution of your content by country:  * none: No geo restriction is enabled, meaning access to content is not restricted by client geo location.  * blacklist: The Location elements specify the countries in which you don't want CloudFront to distribute your content.  * whitelist: The Location elements specify the countries in which you want CloudFront to distribute your content.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Restrictions.GeoRestriction.RestrictionType"),
			},
			{
				Name:        "restrictions_geo_restriction_items",
				Description: "A complex type that contains a Location element for each country in which you want CloudFront either to distribute your content (whitelist) or not distribute your content (blacklist)",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.Restrictions.GeoRestriction.Items"),
			},
			{
				Name:        "viewer_certificate_a_c_m_certificate_arn",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html), provide the Amazon Resource Name (ARN) of the ACM certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.ACMCertificateArn"),
			},
			{
				Name:        "viewer_certificate",
				Description: "This field is deprecated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.Certificate"),
			},
			{
				Name:        "viewer_certificate_certificate_source",
				Description: "This field is deprecated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.CertificateSource"),
			},
			{
				Name:        "viewer_certificate_cloud_front_default_certificate",
				Description: "If the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net, set this field to true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.CloudFrontDefaultCertificate"),
			},
			{
				Name:        "viewer_certificate_iam_certificate_id",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Identity and Access Management (AWS IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html), provide the ID of the IAM certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.IAMCertificateId"),
			},
			{
				Name:        "viewer_certificate_minimum_protocol_version",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.MinimumProtocolVersion"),
			},
			{
				Name:        "viewer_certificate_ssl_support_method",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.  * sni-only – The distribution accepts HTTPS connections from only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.SSLSupportMethod"),
			},
			{
				Name:        "web_acl_id",
				Description: "A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.WebACLId"),
			},
			{
				Name:        "domain_name",
				Description: "The domain name corresponding to the distribution, for example, d111111abcdef8.cloudfront.net.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier for the distribution",
				Type:        schema.TypeString,
			},
			{
				Name:        "in_progress_invalidation_batches",
				Description: "The number of invalidation batches currently in progress.  This member is required.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "last_modified_time",
				Description: "The date and time the distribution was last modified.  This member is required.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status",
				Description: "This response element indicates the current status of the distribution",
				Type:        schema.TypeString,
			},
			{
				Name:        "active_trusted_key_groups_enabled",
				Description: "This field is true if any of the key groups have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ActiveTrustedKeyGroups.Enabled"),
			},
			{
				Name:        "active_trusted_key_groups_quantity",
				Description: "The number of key groups in the list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ActiveTrustedKeyGroups.Quantity"),
			},
			{
				Name:        "active_trusted_signers_enabled",
				Description: "This field is true if any of the AWS accounts in the list have active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ActiveTrustedSigners.Enabled"),
			},
			{
				Name:        "active_trusted_signers_quantity",
				Description: "The number of AWS accounts in the list.  This member is required.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ActiveTrustedSigners.Quantity"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_cloudfront_distribution_default_cache_behavior_lambda_function_associations_items",
				Description: "A complex type that contains a Lambda function association.",
				Resolver:    fetchCloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationsItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "event_type",
						Description: "Specifies the event type that triggers a Lambda function invocation",
						Type:        schema.TypeString,
					},
					{
						Name:        "lambda_function_arn",
						Description: "The ARN of the Lambda function",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LambdaFunctionARN"),
					},
					{
						Name:        "include_body",
						Description: "A flag that allows a Lambda function to have read access to the body content. For more information, see Accessing the Request Body by Choosing the Include Body Option (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_origins_items",
				Description: "An origin",
				Resolver:    fetchCloudfrontDistributionOriginsItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "domain_name",
						Description: "The domain name for the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "A unique identifier for the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_attempts",
						Description: "The number of times that CloudFront attempts to connect to the origin",
						Type:        schema.TypeInt,
					},
					{
						Name:        "connection_timeout",
						Description: "The number of seconds that CloudFront waits when trying to establish a connection to the origin",
						Type:        schema.TypeInt,
					},
					{
						Name:        "custom_headers_quantity",
						Description: "The number of custom headers, if any, for this distribution.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomHeaders.Quantity"),
					},
					{
						Name:        "custom_origin_config_http_port",
						Description: "The HTTP port that CloudFront uses to connect to the origin",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomOriginConfig.HTTPPort"),
					},
					{
						Name:        "custom_origin_config_https_port",
						Description: "The HTTPS port that CloudFront uses to connect to the origin",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomOriginConfig.HTTPSPort"),
					},
					{
						Name:        "custom_origin_config_origin_protocol_policy",
						Description: "Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginProtocolPolicy"),
					},
					{
						Name:        "custom_origin_config_origin_keepalive_timeout",
						Description: "Specifies how long, in seconds, CloudFront persists its connection to the origin",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginKeepaliveTimeout"),
					},
					{
						Name:        "custom_origin_config_origin_read_timeout",
						Description: "Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the origin response timeout",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginReadTimeout"),
					},
					{
						Name:        "custom_origin_config_origin_ssl_protocols_items",
						Description: "A list that contains allowed SSL/TLS protocols for this distribution.  This member is required.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginSslProtocols.Items"),
					},
					{
						Name:        "custom_origin_config_origin_ssl_protocols_quantity",
						Description: "The number of SSL/TLS protocols that you want to allow CloudFront to use when establishing an HTTPS connection with this origin.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginSslProtocols.Quantity"),
					},
					{
						Name:        "origin_path",
						Description: "An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "origin_shield_enabled",
						Description: "A flag that specifies whether Origin Shield is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("OriginShield.Enabled"),
					},
					{
						Name:        "origin_shield_region",
						Description: "The AWS Region for Origin Shield",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OriginShield.OriginShieldRegion"),
					},
					{
						Name:        "s3_origin_config_origin_access_identity",
						Description: "The CloudFront origin access identity to associate with the origin",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3OriginConfig.OriginAccessIdentity"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_cloudfront_distribution_origins_item_custom_headers_items",
						Description: "A complex type that contains HeaderName and HeaderValue elements, if any, for this distribution.",
						Resolver:    fetchCloudfrontDistributionOriginsItemCustomHeadersItems,
						Columns: []schema.Column{
							{
								Name:        "distribution_origins_item_cq_id",
								Description: "Unique CloudQuery ID of aws_cloudfront_distribution_origins_items table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "header_name",
								Description: "The name of a header that you want CloudFront to send to your origin",
								Type:        schema.TypeString,
							},
							{
								Name:        "header_value",
								Description: "The value for the header that you specified in the HeaderName field.  This member is required.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_cache_behaviors_items",
				Description: "A complex type that describes how CloudFront processes requests",
				Resolver:    fetchCloudfrontDistributionCacheBehaviorsItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "path_pattern",
						Description: "The pattern (for example, images/*.jpg) that specifies which requests to apply the behavior to",
						Type:        schema.TypeString,
					},
					{
						Name:        "target_origin_id",
						Description: "The value of ID for the origin that you want CloudFront to route requests to when they match this cache behavior.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "viewer_protocol_policy",
						Description: "The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_methods_items",
						Description: "A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.  This member is required.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("AllowedMethods.Items"),
					},
					{
						Name:        "allowed_methods_quantity",
						Description: "The number of HTTP methods that you want CloudFront to forward to your origin. Valid values are 2 (for GET and HEAD requests), 3 (for GET, HEAD, and OPTIONS requests) and 7 (for GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests).  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("AllowedMethods.Quantity"),
					},
					{
						Name:        "allowed_methods_cached_methods_items",
						Description: "A complex type that contains the HTTP methods that you want CloudFront to cache responses to.  This member is required.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("AllowedMethods.CachedMethods.Items"),
					},
					{
						Name:        "allowed_methods_cached_methods_quantity",
						Description: "The number of HTTP methods for which you want CloudFront to cache responses. Valid values are 2 (for caching responses to GET and HEAD requests) and 3 (for caching responses to GET, HEAD, and OPTIONS requests).  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("AllowedMethods.CachedMethods.Quantity"),
					},
					{
						Name:        "cache_policy_id",
						Description: "The unique identifier of the cache policy that is attached to this cache behavior",
						Type:        schema.TypeString,
					},
					{
						Name:        "compress",
						Description: "Whether you want CloudFront to automatically compress certain files for this cache behavior",
						Type:        schema.TypeBool,
					},
					{
						Name:        "default_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DefaultTTL"),
					},
					{
						Name:        "field_level_encryption_id",
						Description: "The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for this cache behavior.",
						Type:        schema.TypeString,
					},
					{
						Name:        "forwarded_values_cookies_forward",
						Description: "This field is deprecated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ForwardedValues.Cookies.Forward"),
					},
					{
						Name:        "forwarded_values_cookies_whitelisted_names_quantity",
						Description: "The number of cookie names in the Items list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ForwardedValues.Cookies.WhitelistedNames.Quantity"),
					},
					{
						Name:        "forwarded_values_cookies_whitelisted_names_items",
						Description: "A list of cookie names.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ForwardedValues.Cookies.WhitelistedNames.Items"),
					},
					{
						Name:        "forwarded_values_query_string",
						Description: "This field is deprecated",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ForwardedValues.QueryString"),
					},
					{
						Name:        "forwarded_values_headers_quantity",
						Description: "The number of header names in the Items list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ForwardedValues.Headers.Quantity"),
					},
					{
						Name:        "forwarded_values_headers_items",
						Description: "A list of HTTP header names.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ForwardedValues.Headers.Items"),
					},
					{
						Name:        "forwarded_values_query_string_cache_keys_quantity",
						Description: "The number of whitelisted query string parameters for a cache behavior.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ForwardedValues.QueryStringCacheKeys.Quantity"),
					},
					{
						Name:        "forwarded_values_query_string_cache_keys_items",
						Description: "A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ForwardedValues.QueryStringCacheKeys.Items"),
					},
					{
						Name:        "lambda_function_associations_quantity",
						Description: "The number of Lambda function associations for this cache behavior.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LambdaFunctionAssociations.Quantity"),
					},
					{
						Name:        "max_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("MaxTTL"),
					},
					{
						Name:        "min_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("MinTTL"),
					},
					{
						Name:        "origin_request_policy_id",
						Description: "The unique identifier of the origin request policy that is attached to this cache behavior",
						Type:        schema.TypeString,
					},
					{
						Name:        "realtime_log_config_arn",
						Description: "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior",
						Type:        schema.TypeString,
					},
					{
						Name:        "smooth_streaming",
						Description: "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false",
						Type:        schema.TypeBool,
					},
					{
						Name:        "trusted_key_groups_enabled",
						Description: "This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.  This member is required.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TrustedKeyGroups.Enabled"),
					},
					{
						Name:        "trusted_key_groups_quantity",
						Description: "The number of key groups in the list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("TrustedKeyGroups.Quantity"),
					},
					{
						Name:        "trusted_key_groups_items",
						Description: "A list of key groups identifiers.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("TrustedKeyGroups.Items"),
					},
					{
						Name:        "trusted_signers_enabled",
						Description: "This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TrustedSigners.Enabled"),
					},
					{
						Name:        "trusted_signers_quantity",
						Description: "The number of AWS accounts in the list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("TrustedSigners.Quantity"),
					},
					{
						Name:        "trusted_signers_items",
						Description: "A list of AWS account identifiers.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("TrustedSigners.Items"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_cloudfront_distribution_cache_behaviors_item_lambda_function_associations_items",
						Description: "A complex type that contains a Lambda function association.",
						Resolver:    fetchCloudfrontDistributionCacheBehaviorsItemLambdaFunctionAssociationsItems,
						Columns: []schema.Column{
							{
								Name:        "distribution_cache_behaviors_item_cq_id",
								Description: "Unique CloudQuery ID of aws_cloudfront_distribution_cache_behaviors_items table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "event_type",
								Description: "Specifies the event type that triggers a Lambda function invocation",
								Type:        schema.TypeString,
							},
							{
								Name:        "lambda_function_arn",
								Description: "The ARN of the Lambda function",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("LambdaFunctionARN"),
							},
							{
								Name:        "include_body",
								Description: "A flag that allows a Lambda function to have read access to the body content. For more information, see Accessing the Request Body by Choosing the Include Body Option (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
								Type:        schema.TypeBool,
							},
						},
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_custom_error_responses_items",
				Description: "A complex type that controls:  * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.  * How long CloudFront caches HTTP status codes in the 4xx and 5xx range.  For more information about custom error pages, see Customizing Error Responses (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the Amazon CloudFront Developer Guide.",
				Resolver:    fetchCloudfrontDistributionCustomErrorResponsesItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "error_code",
						Description: "The HTTP status code for which you want to specify a custom error page and/or a caching duration.  This member is required.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "error_caching_min_ttl",
						Description: "The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ErrorCachingMinTTL"),
					},
					{
						Name:        "response_code",
						Description: "The HTTP status code that you want CloudFront to return to the viewer along with the custom error page",
						Type:        schema.TypeString,
					},
					{
						Name:        "response_page_path",
						Description: "The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by ErrorCode, for example, /4xx-errors/403-forbidden.html",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_origin_groups_items",
				Description: "An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify",
				Resolver:    fetchCloudfrontDistributionOriginGroupsItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "failover_criteria_status_codes_items",
						Description: "The items (status codes) for an origin group.  This member is required.",
						Type:        schema.TypeIntArray,
						Resolver:    schema.PathResolver("FailoverCriteria.StatusCodes.Items"),
					},
					{
						Name:        "failover_criteria_status_codes_quantity",
						Description: "The number of status codes.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FailoverCriteria.StatusCodes.Quantity"),
					},
					{
						Name:        "id",
						Description: "The origin group's ID.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "members_quantity",
						Description: "The number of origins in an origin group.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Members.Quantity"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_cloudfront_distribution_origin_groups_item_members_items",
						Description: "An origin in an origin group.",
						Resolver:    fetchCloudfrontDistributionOriginGroupsItemMembersItems,
						Columns: []schema.Column{
							{
								Name:        "distribution_origin_groups_item_cq_id",
								Description: "Unique CloudQuery ID of aws_cloudfront_distribution_origin_groups_items table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "origin_id",
								Description: "The ID for an origin in an origin group.  This member is required.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_active_trusted_key_groups_items",
				Description: "A list of identifiers for the public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies.",
				Resolver:    fetchCloudfrontDistributionActiveTrustedKeyGroupsItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key_group_id",
						Description: "The identifier of the key group that contains the public keys.",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_pair_ids_quantity",
						Description: "The number of key pair identifiers in the list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("KeyPairIds.Quantity"),
					},
					{
						Name:        "key_pair_ids_items",
						Description: "A list of CloudFront key pair identifiers.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("KeyPairIds.Items"),
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_active_trusted_signers_items",
				Description: "A list of AWS accounts and the active CloudFront key pairs in each account that CloudFront can use to verify the signatures of signed URLs and signed cookies.",
				Resolver:    fetchCloudfrontDistributionActiveTrustedSignersItems,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "aws_account_number",
						Description: "An AWS account number that contains active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_pair_ids_quantity",
						Description: "The number of key pair identifiers in the list.  This member is required.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("KeyPairIds.Quantity"),
					},
					{
						Name:        "key_pair_ids_items",
						Description: "A list of CloudFront key pair identifiers.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("KeyPairIds.Items"),
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_alias_i_c_p_recordals",
				Description: "AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they've added to CloudFront",
				Resolver:    fetchCloudfrontDistributionAliasICPRecordals,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "c_n_a_m_e",
						Description: "A domain name associated with a distribution.",
						Type:        schema.TypeString,
					},
					{
						Name:        "i_c_p_recordal_status",
						Description: "The Internet Content Provider (ICP) recordal status for a CNAME",
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
func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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

		if response.DistributionList != nil {
			res <- response.DistributionList.Items
		}

		if aws.ToString(response.DistributionList.Marker) == "" {
			break
		}
		config.Marker = response.DistributionList.Marker
	}
	return nil
}
func resolveCloudfrontDistributionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionOriginsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionOriginsItemCustomHeadersItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionCacheBehaviorsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionCacheBehaviorsItemLambdaFunctionAssociationsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionCustomErrorResponsesItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionOriginGroupsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionOriginGroupsItemMembersItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionActiveTrustedKeyGroupsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionActiveTrustedSignersItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionAliasICPRecordals(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- distribution.AliasICPRecordals
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchCloudfrontDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigOriginsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigOriginsItemCustomHeadersItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigCacheBehaviorsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigCacheBehaviorsItemLambdaFunctionAssociationsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigCustomErrorResponsesItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigOriginGroupsItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionConfigOriginGroupsItemMembersItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchCloudfrontDistributionCacheBehaviours(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.CacheBehaviors != nil {
		res <- distribution.CacheBehaviors.Items
	}
	return nil
}
func fetchCloudfrontDistributionCustomErrorResponses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.CustomErrorResponses != nil {
		res <- distribution.CustomErrorResponses.Items
	}
	return nil
}
func resolveCloudfrontDistributionOriginCustomHeaders(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Origin)
	if r.CustomHeaders == nil {
		return nil
	}
	tags := map[string]string{}
	for _, t := range r.CustomHeaders.Items {
		tags[*t.HeaderName] = *t.HeaderValue
	}
	return resource.Set("custom_headers", tags)
}
func fetchCloudfrontDistributionOrigins(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.Origins != nil {
		res <- distribution.Origins.Items
	}
	return nil
}
func fetchCloudfrontDistributionOriginGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.OriginGroups != nil {
		res <- distribution.OriginGroups.Items
	}
	return nil
}
func resolveCloudfrontDistributionOriginGroupMembers(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	if r.Members == nil {
		return nil
	}
	members := make([]string, 0, *r.Members.Quantity)
	for _, t := range r.Members.Items {
		members = append(members, *t.OriginId)
	}
	return resource.Set("members_origin_ids", members)
}
func resolveFailoverCriteriaStatusCodeItems(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	if r.FailoverCriteria == nil || r.FailoverCriteria.StatusCodes == nil {
		return nil
	}
	members := make([]int, 0, *r.Members.Quantity)
	for _, item := range r.FailoverCriteria.StatusCodes.Items {
		members = append(members, int(item))
	}
	return resource.Set(c.Name, members)
}
func fetchCloudfrontDistributionDefaultCacheBehaviourLambdaFunctionAssociations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.DefaultCacheBehavior != nil && distribution.DefaultCacheBehavior.LambdaFunctionAssociations != nil {
		res <- distribution.DefaultCacheBehavior.LambdaFunctionAssociations.Items
	}
	return nil
}
