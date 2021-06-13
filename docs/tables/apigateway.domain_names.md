
# Table: aws_apigateway_domain_names

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|certificate_arn|text||
|certificate_name|text||
|certificate_upload_date|timestamp without time zone||
|distribution_domain_name|text||
|distribution_hosted_zone_id|text||
|domain_name|text||
|domain_name_status|text||
|domain_name_status_message|text||
|endpoint_configuration_types|text[]||
|endpoint_configuration_vpc_endpoint_ids|text[]||
|mutual_tls_authentication_truststore_uri|text||
|mutual_tls_authentication_truststore_version|text||
|mutual_tls_authentication_truststore_warnings|text[]||
|regional_certificate_arn|text||
|regional_certificate_name|text||
|regional_domain_name|text||
|regional_hosted_zone_id|text||
|security_policy|text||
|tags|jsonb||
## Relations
## Table: aws_apigateway_domain_name_base_path_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid||
|base_path|text||
|rest_api_id|text||
|stage|text||
