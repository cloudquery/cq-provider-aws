
# Table: aws_appsync_graphql_apis
Describes a GraphQL API
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|id|text|The API's unique identifier|
|arn|text|The Amazon Resource Name (ARN) that identifies the resource|
|name|text|The name of the API|
|uris|jsonb|The URIs associated with the GraphQL API|
|waf_web_acl_arn|text|The ARN of the WAF access control list (ACL) associated with this GraphqlApi, if one exists|
|xray_enabled|boolean|A flag indicating whether to use X-Ray tracing for this GraphqlApi|
|authentication_type|text|The authentication type|
|lambda_authorizer_authorize_uri|text|The Amazon Resource Name (ARN) of the Lambda function to be called for authorization|
|lambda_authorizer_result_ttl_in_seconds|integer|The number of seconds a response should be cached for|
|lambda_authorizer_identity_validation_expression|text|A regular expression for validation of tokens before the Lambda function is called|
|log_config_cloudwatch_logs_role_arn|text|The service role that AppSync assumes to publish to CloudWatch logs|
|log_config_field_log_level|text|The field logging level|
|log_config_exclude_verbose_content|boolean|Whether to exclude sections that contain information such as headers, context, and evaluated mapping templates|
|openid_connect_config_issuer|text|The issuer for the OIDC configuration|
|openid_connect_config_auth_ttl|integer|The number of milliseconds that a token is valid after being authenticated|
|openid_connect_config_client_id|text|The client identifier of the relying party at the OpenID identity provider|
|openid_connect_config_iat_ttl|integer|The number of milliseconds that a token is valid after it's issued to a user|
|user_pool_config_aws_region|text|The AWS Region in which the user pool was created|
|user_pool_config_default_action|text|The action to take when a request doesn't match the Amazon Cognito user pool configuration|
|user_pool_config_user_pool_id|text|The user pool ID|
|user_pool_config_app_id_client_regex|text|A regular expression for validating the incoming Amazon Cognito user pool app client ID|
