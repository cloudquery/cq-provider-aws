
# Table: aws_cloudfront_distributions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|arn|text||
|aliases_items|text[]||
|comment|text||
|cache_behaviour_target_origin_id|text||
|cache_behaviour_viewer_protocol_policy|text||
|cache_behaviour_allowed_methods|text[]||
|cache_behaviour_allowed_methods_cached_methods|text[]||
|cache_behaviour_cache_policy_id|text||
|cache_behaviour_compress|boolean||
|cache_behaviour_default_ttl|bigint||
|cache_behaviour_field_level_encryption_id|text||
|cache_behaviour_forwarded_values_cookies_forward|text||
|cache_behaviour_forwarded_values_cookies_white_listed_names|text[]||
|cache_behaviour_forwarded_values_query_string|boolean||
|cache_behaviour_forwarded_values_headers|text[]||
|cache_behaviour_forwarded_values_query_string_cache_keys|text[]||
|cache_behaviour_max_ttl|bigint||
|cache_behaviour_min_ttl|bigint||
|cache_behaviour_origin_request_policy_id|text||
|cache_behaviour_realtime_log_config_arn|text||
|cache_behaviour_smooth_streaming|boolean||
|cache_behaviour_trusted_key_groups_enabled|boolean||
|cache_behaviour_trusted_key_groups|text[]||
|cache_behaviour_trusted_signers_enabled|boolean||
|cache_behaviour_trusted_signers|text[]||
|domain_name|text||
|enabled|boolean||
|http_version|text||
|resource_id|text||
|ip_v6_enabled|boolean||
|last_modified_time|timestamp without time zone||
|price_class|text||
|restrictions_geo_restriction_restriction_type|text||
|restrictions_geo_restriction_restriction_items|text[]||
|status|text||
|viewer_certificate_acm_certificate_arn|text||
|viewer_certificate|text||
|viewer_certificate_source|text||
|viewer_certificate_cloudfront_default_certificate|boolean||
|viewer_certificate_iam_certificate_id|text||
|viewer_certificate_minimum_protocol_version|text||
|viewer_certificate_ssl_support_method|text||
|web_acl_id|text||
## Relations
## Table: aws_cloudfront_distribution_cache_behaviours

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|path_pattern|text||
|target_origin_id|text||
|viewer_protocol_policy|text||
|allowed_methods|text[]||
|cached_methods|text[]||
## Table: cache_behaviour_lambda_function_associations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|event_type|text||
|lambda_function_arn|text||
|include_body|boolean||
## Table: aws_cloudfront_distribution_custom_error_responses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|error_code|integer||
|error_caching_min_ttl|bigint||
|response_code|text||
|response_page_path|text||
## Table: aws_cloudfront_distribution_origins

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|domain_name|text||
|origin_id|text||
|connection_attempts|integer||
|connection_timeout|integer||
|custom_headers|jsonb||
|custom_origin_config_http_port|integer||
|custom_origin_config_https_port|integer||
|custom_origin_config_protocol_policy|text||
|custom_origin_config_keepalive_timeout|integer||
|custom_origin_config_read_timeout|integer||
|custom_origin_config_ssl_protocols|text[]||
|origin_path|text||
|origin_shield_enabled|boolean||
|origin_shield_region|text||
|s3_origin_config_origin_access_identity|text||
## Table: aws_cloudfront_distribution_alias_icp_recordals

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|cname|text||
|icp_recordal_status|text||
## Table: aws_cloudfront_distribution_origin_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|failover_criteria_status_codes_items|integer[]||
|origin_group_id|text||
|members_origin_ids|text[]||
