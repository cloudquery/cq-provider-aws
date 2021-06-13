
# Table: aws_apigateway_rest_apis

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|api_key_source|text||
|binary_media_types|text[]||
|created_date|timestamp without time zone||
|description|text||
|disable_execute_api_endpoint|boolean||
|endpoint_configuration_types|text[]||
|endpoint_configuration_vpc_endpoint_ids|text[]||
|resource_id|text||
|minimum_compression_size|integer||
|name|text||
|policy|text||
|tags|jsonb||
|version|text||
|warnings|text[]||
## Relations
## Table: aws_apigateway_rest_api_authorizers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|auth_type|text||
|authorizer_credentials|text||
|authorizer_result_ttl_in_seconds|integer||
|authorizer_uri|text||
|resource_id|text||
|identity_source|text||
|identity_validation_expression|text||
|name|text||
|provider_arns|text[]||
|type|text||
## Table: aws_apigateway_rest_api_deployments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|api_summary|jsonb||
|created_date|timestamp without time zone||
|description|text||
|resource_id|text||
## Table: aws_apigateway_rest_api_documentation_parts

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|documentation_part_id|text||
|location_type|text||
|location_method|text||
|location_name|text||
|location_path|text||
|location_status_code|text||
|properties|text||
## Table: aws_apigateway_rest_api_documentation_versions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|created_date|timestamp without time zone||
|description|text||
|version|text||
## Table: aws_apigateway_rest_api_gateway_responses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|default_response|boolean||
|response_parameters|jsonb||
|response_templates|jsonb||
|response_type|text||
|status_code|text||
## Table: aws_apigateway_rest_api_models

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|model_template|text||
|content_type|text||
|description|text||
|resource_id|text||
|name|text||
|schema|text||
## Table: aws_apigateway_rest_api_request_validators

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|resource_id|text||
|name|text||
|validate_request_body|boolean||
|validate_request_parameters|boolean||
## Table: aws_apigateway_rest_api_resources

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|resource_id|text||
|parent_id|text||
|path|text||
|path_part|text||
|resource_methods|jsonb||
## Table: aws_apigateway_rest_api_stages

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid||
|access_log_settings_destination_arn|text||
|access_log_settings_format|text||
|cache_cluster_enabled|boolean||
|cache_cluster_size|text||
|cache_cluster_status|text||
|canary_settings_deployment_id|text||
|canary_settings_percent_traffic|float||
|canary_settings_stage_variable_overrides|jsonb||
|canary_settings_use_stage_cache|boolean||
|client_certificate_id|text||
|created_date|timestamp without time zone||
|deployment_id|text||
|description|text||
|documentation_version|text||
|last_updated_date|timestamp without time zone||
|method_settings|jsonb||
|stage_name|text||
|tags|jsonb||
|tracing_enabled|boolean||
|variables|jsonb||
|web_acl_arn|text||
