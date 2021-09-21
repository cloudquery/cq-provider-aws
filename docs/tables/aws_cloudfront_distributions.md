
# Table: aws_cloudfront_distributions
A summary of the information about a CloudFront distribution.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|tags|jsonb||
|arn|text|The ARN (Amazon Resource Name) for the distribution|
|aliases|text[]|A complex type that contains the CNAME aliases, if any, that you want to associate with this distribution.|
|comment|text|The comment originally specified when this distribution was created.  This member is required.|
|cache_behaviour_target_origin_id|text|The value of ID for the origin that you want CloudFront to route requests to when they use the default cache behavior.  This member is required.|
|cache_behaviour_viewer_protocol_policy|text|The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern|
|cache_behaviour_allowed_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.  This member is required.|
|cache_behaviour_allowed_methods_cached_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to cache responses to.  This member is required.|
|cache_behaviour_cache_policy_id|text|The unique identifier of the cache policy that is attached to the default cache behavior|
|cache_behaviour_compress|boolean|Whether you want CloudFront to automatically compress certain files for this cache behavior|
|cache_behaviour_default_ttl|bigint|This field is deprecated|
|cache_behaviour_field_level_encryption_id|text|The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.|
|cache_behaviour_forwarded_values_cookies_forward|text|This field is deprecated|
|cache_behaviour_forwarded_values_cookies_whitelisted_names|text[]|A list of cookie names.|
|cache_behaviour_forwarded_values_query_string|boolean|This field is deprecated|
|cache_behaviour_forwarded_values_headers|text[]|A list of HTTP header names.|
|cache_behaviour_forwarded_values_query_string_cache_keys|text[]|A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior|
|cache_behaviour_max_ttl|bigint|This field is deprecated|
|cache_behaviour_min_ttl|bigint|This field is deprecated|
|cache_behaviour_origin_request_policy_id|text|The unique identifier of the origin request policy that is attached to the default cache behavior|
|cache_behaviour_realtime_log_config_arn|text|The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior|
|cache_behaviour_smooth_streaming|boolean|Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false|
|cache_behaviour_trusted_key_groups_enabled|boolean|This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.  This member is required.|
|cache_behaviour_trusted_key_groups|text[]|A list of key groups identifiers.|
|cache_behaviour_trusted_signers_enabled|boolean|This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies|
|cache_behaviour_trusted_signers|text[]|A list of AWS account identifiers.|
|domain_name|text|The domain name that corresponds to the distribution, for example, d111111abcdef8.cloudfront.net.  This member is required.|
|enabled|boolean|Whether the distribution is enabled to accept user requests for content.  This member is required.|
|http_version|text|Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront|
|id|text|The identifier for the distribution|
|ip_v6_enabled|boolean|Whether CloudFront responds to IPv6 DNS requests with an IPv6 address for your distribution.  This member is required.|
|last_modified_time|timestamp without time zone|The date and time the distribution was last modified.  This member is required.|
|price_class|text|A complex type that contains information about price class for this streaming distribution.  This member is required.|
|restrictions_geo_restriction_type|text|The method that you want to use to restrict distribution of your content by country:  * none: No geo restriction is enabled, meaning access to content is not restricted by client geo location.  * blacklist: The Location elements specify the countries in which you don't want CloudFront to distribute your content.  * whitelist: The Location elements specify the countries in which you want CloudFront to distribute your content.  This member is required.|
|restrictions_geo_restrictions|text[]|A complex type that contains a Location element for each country in which you want CloudFront either to distribute your content (whitelist) or not distribute your content (blacklist)|
|status|text|The current status of the distribution|
|viewer_certificate_acm_certificate_arn|text|If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html), provide the Amazon Resource Name (ARN) of the ACM certificate|
|viewer_certificate|text|This field is deprecated|
|viewer_certificate_source|text|This field is deprecated|
|viewer_certificate_cloudfront_default_certificate|boolean|If the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net, set this field to true|
|viewer_certificate_iam_certificate_id|text|If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Identity and Access Management (AWS IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html), provide the ID of the IAM certificate|
|viewer_certificate_minimum_protocol_version|text|If the distribution uses Aliases (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers|
|viewer_certificate_ssl_support_method|text|If the distribution uses Aliases (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.  * sni-only – The distribution accepts HTTPS connections from only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended|
|web_acl_id|text|The Web ACL Id (if any) associated with the distribution.  This member is required.|
