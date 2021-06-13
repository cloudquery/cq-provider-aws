
# Table: aws_apigateway_rest_apis
Represents a REST API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|api_key_source|text|The source of the API key for metering requests according to a usage plan. Valid values are:|
|binary_media_types|text[]|The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.|
|created_date|timestamp without time zone|The timestamp when the API was created.|
|description|text|The API's description.|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.|
|endpoint_configuration_types|text[]|A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is "EDGE". For a regional API and its custom domain name, the endpoint type is REGIONAL. For a private API, the endpoint type is PRIVATE.|
|endpoint_configuration_vpc_endpoint_ids|text[]|A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for PRIVATE endpoint type.|
|resource_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|minimum_compression_size|integer|A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.|
|name|text|The API's name.|
|policy|text|A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|version|text|A version identifier for the API.|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import.|
## Relations
## Table: aws_apigateway_rest_api_authorizers
Represents an authorization layer for methods.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|auth_type|text|Optional customer-defined field, used in OpenAPI imports and exports without functional impact.|
|authorizer_credentials|text|Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.|
|authorizer_result_ttl_in_seconds|integer|The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.|
|authorizer_uri|text|Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or REQUEST authorizers, this must be a well-formed Lambda function URI, for example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations. In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api}, where {region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial /. For Lambda functions, this is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.|
|resource_id|text|The identifier for the authorizer resource.|
|identity_source|text|The identity source for which authorization is requested.|
|identity_validation_expression|text|A validation expression for the incoming identity token. For TOKEN authorizers, this value is a regular expression. For COGNITO_USER_POOLS authorizers, API Gateway will match the aud field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the REQUEST authorizer.|
|name|text|[Required] The name of the authorizer.|
|provider_arns|text[]|A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer. Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}. For a TOKEN or REQUEST authorizer, this is not defined.|
|type|text|The authorizer type. Valid values are TOKEN for a Lambda function using a single authorization token submitted in a custom header, REQUEST for a Lambda function using incoming request parameters, and COGNITO_USER_POOLS for using an Amazon Cognito user pool.|
## Table: aws_apigateway_rest_api_deployments
An immutable representation of a RestApi resource that can be called by users using Stages.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|api_summary|jsonb|A summary of the RestApi at the date and time that the deployment resource was created.|
|created_date|timestamp without time zone|The date and time that the deployment resource was created.|
|description|text|The description for the deployment resource.|
|resource_id|text|The identifier for the deployment resource.|
## Table: aws_apigateway_rest_api_documentation_parts
A documentation part for a targeted API entity.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|documentation_part_id|text|The DocumentationPart identifier, generated by API Gateway when the DocumentationPart is created.|
|location_type|text|[Required] The type of API entity to which the documentation content applies. Valid values are API, AUTHORIZER, MODEL, RESOURCE, METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. Content inheritance does not apply to any entity of the API, AUTHORIZER, METHOD, MODEL, REQUEST_BODY, or RESOURCE type.|
|location_method|text|The HTTP verb of a method. It is a valid field for the API entity types of METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is * for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other location attributes, the child entity's method attribute must match that of the parent entity exactly.|
|location_name|text|The name of the targeted API entity. It is a valid and required field for the API entity types of AUTHORIZER, MODEL, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY and RESPONSE_HEADER. It is an invalid field for any other entity type.|
|location_path|text|The URL path of the target. It is a valid field for the API entity types of RESOURCE, METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is / for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other location attributes, the child entity's path attribute must match that of the parent entity as a prefix.|
|location_status_code|text|The HTTP status code of a response. It is a valid field for the API entity types of RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is * for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other location attributes, the child entity's statusCode attribute must match that of the parent entity exactly.|
|properties|text|A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only OpenAPI-compliant documentation-related fields from the properties map are exported and, hence, published as part of the API entity definitions, while the original documentation parts are exported in a OpenAPI extension of x-amazon-apigateway-documentation.|
## Table: aws_apigateway_rest_api_documentation_versions
A snapshot of the documentation of an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|created_date|timestamp without time zone|The date when the API documentation snapshot is created.|
|description|text|The description of the API documentation snapshot.|
|version|text|The version identifier of the API documentation snapshot.|
## Table: aws_apigateway_rest_api_gateway_responses
A gateway response of a given response type and status code, with optional response parameters and mapping templates.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|default_response|boolean|A Boolean flag to indicate whether this GatewayResponse is the default gateway response (true) or not (false). A default gateway response is one generated by API Gateway without any customization by an API developer.|
|response_parameters|jsonb|Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs.|
|response_templates|jsonb|Response templates of the GatewayResponse as a string-to-string map of key-value pairs.|
|response_type|text|The response type of the associated GatewayResponse. Valid values are|
|status_code|text|The HTTP status code for this GatewayResponse.|
## Table: aws_apigateway_rest_api_models
Represents the data structure of a method's request or response payload.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|model_template|text||
|content_type|text|The content-type for the model.|
|description|text|The description of the model.|
|resource_id|text|The identifier for the model resource.|
|name|text|The name of the model. Must be an alphanumeric string.|
|schema|text|The schema for the model. For application/json models, this should be JSON schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.|
## Table: aws_apigateway_rest_api_request_validators
A set of validation rules for incoming Method requests.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|resource_id|text|The identifier of this RequestValidator.|
|name|text|The name of this RequestValidator|
|validate_request_body|boolean|A Boolean flag to indicate whether to validate a request body according to the configured Model schema.|
|validate_request_parameters|boolean|A Boolean flag to indicate whether to validate request parameters (true) or not (false).|
## Table: aws_apigateway_rest_api_resources
Represents an API resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|resource_id|text|The resource's identifier.|
|parent_id|text|The parent resource's identifier.|
|path|text|The full path for this resource.|
|path_part|text|The last path segment for this resource.|
|resource_methods|jsonb|Gets an API resource's method of a given HTTP verb. The resource methods are a map of methods indexed by methods' HTTP verbs enabled on the resource. This method map is included in the 200 OK response of the GET /restapis/{restapi_id}/resources/{resource_id} or GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods request. Example: Get the GET method of an API resource|
## Table: aws_apigateway_rest_api_stages
Represents a unique identifier for a version of a deployed RestApi that is callable by users.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|access_log_settings_destination_arn|text|The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-.|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least $context.requestId.|
|cache_cluster_enabled|boolean|Specifies whether a cache cluster is enabled for the stage.|
|cache_cluster_size|text|The size of the cache cluster for the stage, if enabled.|
|cache_cluster_status|text|The status of the cache cluster for the stage, if enabled.|
|canary_settings_deployment_id|text|The ID of the canary deployment.|
|canary_settings_percent_traffic|float|The percent (0-100) of traffic diverted to a canary deployment.|
|canary_settings_stage_variable_overrides|jsonb|Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.|
|canary_settings_use_stage_cache|boolean|A Boolean flag to indicate whether the canary deployment uses the stage cache or not.|
|client_certificate_id|text|The identifier of a client certificate for an API stage.|
|created_date|timestamp without time zone|The timestamp when the stage was created.|
|deployment_id|text|The identifier of the Deployment that the stage points to.|
|description|text|The stage's description.|
|documentation_version|text|The version of the associated API documentation.|
|last_updated_date|timestamp without time zone|The timestamp when the stage last updated.|
|method_settings|jsonb|A map that defines the method settings for a Stage resource. Keys (designated as /{method_setting_key below) are method paths defined as {resource_path}/{http_method} for an individual method override, or /\*/\* for overriding all methods in the stage.|
|stage_name|text|The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|tracing_enabled|boolean|Specifies whether active tracing with X-ray is enabled for the Stage.|
|variables|jsonb|A map that defines the stage variables for a Stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.|
|web_acl_arn|text|The ARN of the WebAcl associated with the Stage.|
