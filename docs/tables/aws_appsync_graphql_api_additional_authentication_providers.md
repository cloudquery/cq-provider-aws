
# Table: aws_appsync_graphql_api_additional_authentication_providers
Describes an authentication provider.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|graphql_api_cq_id|uuid|Unique CloudQuery ID of aws_appsync_graphql_apis table (FK)|
|authentication_type|text|The authentication type|
|lambda_authorizer_authorize_uri|text|The Amazon Resource Name (ARN) of the Lambda function to be called for authorization|
|lambda_authorizer_result_ttl_in_seconds|integer|The number of seconds a response should be cached for|
|lambda_authorizer_identity_validation_expression|text|A regular expression for validation of tokens before the Lambda function is called|
|openid_connect_config_issuer|text|The issuer for the OIDC configuration|
|openid_connect_config_auth_ttl|integer|The number of milliseconds that a token is valid after being authenticated|
|openid_connect_config_client_id|text|The client identifier of the relying party at the OpenID identity provider|
|openid_connect_config_iat_ttl|integer|The number of milliseconds that a token is valid after it's issued to a user|
|user_pool_config_aws_region|text|The AWS Region in which the user pool was created|
|user_pool_config_user_pool_id|text|The user pool ID|
|user_pool_config_app_id_client_regex|text|A regular expression for validating the incoming Amazon Cognito user pool app client ID|
