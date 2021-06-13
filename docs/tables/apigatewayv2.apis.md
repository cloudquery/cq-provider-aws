
# Table: aws_apigatewayv2_apis

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|name|text||
|protocol_type|text||
|route_selection_expression|text||
|api_endpoint|text||
|api_gateway_managed|boolean||
|api_id|text||
|api_key_selection_expression|text||
|cors_configuration_allow_credentials|boolean||
|cors_configuration_allow_headers|text[]||
|cors_configuration_allow_methods|text[]||
|cors_configuration_allow_origins|text[]||
|cors_configuration_expose_headers|text[]||
|cors_configuration_max_age|integer||
|created_date|timestamp without time zone||
|description|text||
|disable_execute_api_endpoint|boolean||
|disable_schema_validation|boolean||
|import_info|text[]||
|tags|jsonb||
|version|text||
|warnings|text[]||
## Relations
## Table: aws_apigatewayv2_api_authorizers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|name|text||
|authorizer_credentials_arn|text||
|authorizer_id|text||
|authorizer_payload_format_version|text||
|authorizer_result_ttl_in_seconds|integer||
|authorizer_type|text||
|authorizer_uri|text||
|enable_simple_responses|boolean||
|identity_source|text[]||
|identity_validation_expression|text||
|jwt_configuration_audience|text[]||
|jwt_configuration_issuer|text||
## Table: aws_apigatewayv2_api_deployments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|auto_deployed|boolean||
|created_date|timestamp without time zone||
|deployment_id|text||
|deployment_status|text||
|deployment_status_message|text||
|description|text||
## Table: aws_apigatewayv2_api_integrations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|api_gateway_managed|boolean||
|connection_id|text||
|connection_type|text||
|content_handling_strategy|text||
|credentials_arn|text||
|description|text||
|integration_id|text||
|integration_method|text||
|integration_response_selection_expression|text||
|integration_subtype|text||
|integration_type|text||
|integration_uri|text||
|passthrough_behavior|text||
|payload_format_version|text||
|request_parameters|jsonb||
|request_templates|jsonb||
|response_parameters|jsonb||
|template_selection_expression|text||
|timeout_in_millis|integer||
|tls_config_server_name_to_verify|text||
## Relations
## Table: aws_apigatewayv2_api_integration_responses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_integration_id|uuid||
|integration_response_key|text||
|content_handling_strategy|text||
|integration_response_id|text||
|response_parameters|jsonb||
|response_templates|jsonb||
|template_selection_expression|text||
## Table: aws_apigatewayv2_api_models

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|model_template|text||
|name|text||
|content_type|text||
|description|text||
|model_id|text||
|schema|text||
## Table: aws_apigatewayv2_api_routes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|route_key|text||
|api_gateway_managed|boolean||
|api_key_required|boolean||
|authorization_scopes|text[]||
|authorization_type|text||
|authorizer_id|text||
|model_selection_expression|text||
|operation_name|text||
|request_models|jsonb||
|request_parameters|jsonb||
|route_id|text||
|route_response_selection_expression|text||
|target|text||
## Relations
## Table: aws_apigatewayv2_api_route_responses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_route_id|uuid||
|route_response_key|text||
|model_selection_expression|text||
|response_models|jsonb||
|response_parameters|jsonb||
|route_response_id|text||
## Table: aws_apigatewayv2_api_stages

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid||
|stage_name|text||
|access_log_settings_destination_arn|text||
|access_log_settings_format|text||
|api_gateway_managed|boolean||
|auto_deploy|boolean||
|client_certificate_id|text||
|created_date|timestamp without time zone||
|route_settings_data_trace_enabled|boolean||
|route_settings_detailed_metrics_enabled|boolean||
|route_settings_logging_level|text||
|route_settings_throttling_burst_limit|integer||
|route_settings_throttling_rate_limit|float||
|deployment_id|text||
|description|text||
|last_deployment_status_message|text||
|last_updated_date|timestamp without time zone||
|route_settings|jsonb||
|stage_variables|jsonb||
|tags|jsonb||
