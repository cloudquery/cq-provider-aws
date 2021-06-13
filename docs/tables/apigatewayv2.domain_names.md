
# Table: aws_apigatewayv2_domain_names

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|domain_name|text||
|api_mapping_selection_expression|text||
|mutual_tls_authentication_truststore_uri|text||
|mutual_tls_authentication_truststore_version|text||
|mutual_tls_authentication_truststore_warnings|text[]||
|tags|jsonb||
## Relations
## Table: aws_apigatewayv2_domain_name_configurations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid||
|api_gateway_domain_name|text||
|certificate_arn|text||
|certificate_name|text||
|certificate_upload_date|timestamp without time zone||
|domain_name_status|text||
|domain_name_status_message|text||
|endpoint_type|text||
|hosted_zone_id|text||
|security_policy|text||
## Table: aws_apigatewayv2_domain_name_api_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid||
|api_id|text||
|stage|text||
|api_mapping_id|text||
|api_mapping_key|text||
