
# Table: aws_apigatewayv2_domain_names
Represents a domain name.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|domain_name|text|The name of the DomainName resource.|
|api_mapping_selection_expression|text|The API mapping selection expression.|
|mutual_tls_authentication_truststore_uri|text|An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.|
|mutual_tls_authentication_truststore_version|text|The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.|
|mutual_tls_authentication_truststore_warnings|text[]|A list of warnings that API Gateway returns while processing your truststore. Invalid certificates produce warnings. Mutual TLS is still enabled, but some clients might not be able to access your API. To resolve warnings, upload a new truststore to S3, and then update you domain name to use the new version.|
|tags|jsonb|The collection of tags associated with a domain name.|
## Relations
## Table: aws_apigatewayv2_domain_name_configurations
The domain name configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid|Unique ID of aws_apigatewayv2_domain_names table (FK)|
|api_gateway_domain_name|text|A domain name for the API.|
|certificate_arn|text|An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.|
|certificate_name|text|The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.|
|certificate_upload_date|timestamp without time zone|The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.|
|domain_name_status|text|The status of the domain name migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.|
|domain_name_status_message|text|An optional text message containing detailed information about status of the domain name migration.|
|endpoint_type|text|The endpoint type.|
|hosted_zone_id|text|The Amazon Route 53 Hosted Zone ID of the endpoint.|
|security_policy|text|The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are TLS_1_0 and TLS_1_2.|
## Table: aws_apigatewayv2_domain_name_rest_api_mappings
Represents an API mapping.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid|Unique ID of aws_apigatewayv2_domain_names table (FK)|
|api_id|text|The API identifier.|
|stage|text|The API stage.|
|api_mapping_id|text|The API mapping identifier.|
|api_mapping_key|text|The API mapping key.|
