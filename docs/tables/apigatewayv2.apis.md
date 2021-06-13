
# Table: aws_apigatewayv2_apis
Represents an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|name|text|The name of the API.|
|protocol_type|text|The API protocol.|
|route_selection_expression|text|The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.|
|api_endpoint|text|The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The stage name is typically appended to this URI to form a complete path to a deployed API stage.|
|api_gateway_managed|boolean|Specifies whether an API is managed by API Gateway. You can't update or delete a managed API by using API Gateway. A managed API can be deleted only through the tooling or service that created it.|
|api_id|text|The API ID.|
|api_key_selection_expression|text|An API key selection expression. Supported only for WebSocket APIs. See API Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).|
|cors_configuration_allow_credentials|boolean|Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.|
|cors_configuration_allow_headers|text[]|Represents a collection of allowed headers. Supported only for HTTP APIs.|
|cors_configuration_allow_methods|text[]|Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.|
|cors_configuration_allow_origins|text[]|Represents a collection of allowed origins. Supported only for HTTP APIs.|
|cors_configuration_expose_headers|text[]|Represents a collection of exposed headers. Supported only for HTTP APIs.|
|cors_configuration_max_age|integer|The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.|
|created_date|timestamp without time zone|The timestamp when the API was created.|
|description|text|The description of the API.|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.|
|disable_schema_validation|boolean|Avoid validating models when creating a deployment. Supported only for WebSocket APIs.|
|import_info|text[]|The validation information during API import. This may include particular properties of your OpenAPI definition which are ignored during import. Supported only for HTTP APIs.|
|tags|jsonb|A collection of tags associated with the API.|
|version|text|A version identifier for the API.|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import.|
## Relations
## Table: aws_apigatewayv2_api_authorizers
Represents an authorizer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|name|text|The name of the authorizer.|
|authorizer_credentials_arn|text|Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, don't specify this parameter. Supported only for REQUEST authorizers.|
|authorizer_id|text|The authorizer identifier.|
|authorizer_payload_format_version|text|Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are 1.0 and 2.0. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).|
|authorizer_result_ttl_in_seconds|integer|The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.|
|authorizer_type|text|The authorizer type. Specify REQUEST for a Lambda function using incoming request parameters. Specify JWT to use JSON Web Tokens (supported only for HTTP APIs).|
|authorizer_uri|text|The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers, this must be a well-formed Lambda function URI, for example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations. In general, the URI has this form: arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial /. For Lambda functions, this is usually of the form /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST authorizers.|
|enable_simple_responses|boolean|Specifies whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html)|
|identity_source|text[]|The identity source for which authorization is requested. For a REQUEST authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with $, for example, $request.header.Auth, $request.querystring.Name. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html). For JWT, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example $request.header.Authorization.|
|identity_validation_expression|text|The validation expression does not apply to the REQUEST authorizer.|
|jwt_configuration_audience|text[]|A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list. See RFC 7519 (https://tools.ietf.org/html/rfc7519#section-4.1.3). Supported only for HTTP APIs.|
|jwt_configuration_issuer|text|The base domain of the identity provider that issues JSON Web Tokens. For example, an Amazon Cognito user pool has the following format: https://cognito-idp.{region}.amazonaws.com/{userPoolId} . Required for the JWT authorizer type. Supported only for HTTP APIs.|
## Table: aws_apigatewayv2_api_deployments
An immutable representation of an API that can be called by users.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|auto_deployed|boolean|Specifies whether a deployment was automatically released.|
|created_date|timestamp without time zone|The date and time when the Deployment resource was created.|
|deployment_id|text|The identifier for the deployment.|
|deployment_status|text|The status of the deployment: PENDING, FAILED, or SUCCEEDED.|
|deployment_status_message|text|May contain additional feedback on the status of an API deployment.|
|description|text|The description for the deployment.|
## Table: aws_apigatewayv2_api_integrations
Represents an integration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|api_gateway_managed|boolean|Specifies whether an integration is managed by API Gateway. If you created an API using using quick create, the resulting integration is managed by API Gateway. You can update a managed integration, but you can't delete it.|
|connection_id|text|The ID of the VPC link for a private integration. Supported only for HTTP APIs.|
|connection_type|text|The type of the network connection to the integration endpoint. Specify INTERNET for connections through the public routable internet or VPC_LINK for private connections between API Gateway and resources in a VPC. The default value is INTERNET.|
|content_handling_strategy|text|Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded string. If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.|
|credentials_arn|text|Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, specify null.|
|description|text|Represents the description of an integration.|
|integration_id|text|Represents the identifier of an integration.|
|integration_method|text|Specifies the integration's HTTP method type.|
|integration_response_selection_expression|text|The integration response selection expression for the integration. Supported only for WebSocket APIs. See Integration Response Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions).|
|integration_subtype|text|Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service action to invoke. To learn more, see Integration subtype reference (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html).|
|integration_type|text|The integration type of an integration. One of the following: AWS: for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs. AWS_PROXY: for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration. HTTP: for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs. HTTP_PROXY: for integrating the route or method request with an HTTP endpoint, with the client request passed through as-is. This is also referred to as HTTP proxy integration. MOCK: for integrating the route or method request with API Gateway as a "loopback" endpoint without invoking any backend. Supported only for WebSocket APIs.|
|integration_uri|text|For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If you specify the ARN of an AWS Cloud Map service, API Gateway uses DiscoverInstances to identify resources. You can use query parameters to target specific resources. To learn more, see DiscoverInstances (https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html). For private integrations, all resources must be owned by the same AWS account.|
|passthrough_behavior|text|Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource. There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported only for WebSocket APIs. WHEN_NO_MATCH passes the request body for unmapped content types through to the integration backend without transformation. NEVER rejects unmapped content types with an HTTP 415 Unsupported Media Type response. WHEN_NO_TEMPLATES allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same HTTP 415 Unsupported Media Type response.|
|payload_format_version|text|Specifies the format of the payload sent to an integration. Required for HTTP APIs.|
|request_parameters|jsonb|For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of method.request.{location}.{name} , where {location} is querystring, path, or header; and {name} must be a valid and unique method request parameter name. For HTTP API integrations with a specified integrationSubtype, request parameters are a key-value map specifying parameters that are passed to AWS_PROXY integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Working with AWS service integrations for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services.html). For HTTP API itegrations, without a specified integrationSubtype request parameters are a key-value map specifying how to transform HTTP requests before sending them to backend integrations. The key should follow the pattern <action>:<header|querystring|path>.<location>. The action can be append, overwrite or remove. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).|
|request_templates|jsonb|Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.|
|response_parameters|jsonb|Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend integration before returning the response to clients. Specify a key-value map from a selection key to response parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key must match pattern <action>:<header>.<location> or overwrite.statuscode. The action can be append, overwrite or remove. The value can be a static value, or map to response data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).|
|template_selection_expression|text|The template selection expression for the integration. Supported only for WebSocket APIs.|
|timeout_in_millis|integer|Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.|
|tls_config_server_name_to_verify|text|If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.|
## Relations
## Table: aws_apigatewayv2_api_integration_responses
Represents an integration response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_integration_id|uuid|Unique ID of aws_apigatewayv2_api_integrations table (FK)|
|integration_response_key|text|The integration response key.|
|content_handling_strategy|text|Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded string. If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.|
|integration_response_id|text|The integration response ID.|
|response_parameters|jsonb|A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of method.response.header.{name}, where name is a valid and unique header name. The mapped non-static value must match the pattern of integration.response.header.{name} or integration.response.body.{JSON-expression}, where name is a valid and unique response header name and JSON-expression is a valid JSON expression without the $ prefix.|
|response_templates|jsonb|The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.|
|template_selection_expression|text|The template selection expressions for the integration response.|
## Table: aws_apigatewayv2_api_models
Represents a data model for an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|model_template|text||
|name|text|The name of the model. Must be alphanumeric.|
|content_type|text|The content-type for the model, for example, "application/json".|
|description|text|The description of the model.|
|model_id|text|The model identifier.|
|schema|text|The schema for the model. For application/json models, this should be JSON schema draft 4 model.|
## Table: aws_apigatewayv2_api_routes
Represents a route.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|route_key|text|The route key for the route.|
|api_gateway_managed|boolean|Specifies whether a route is managed by API Gateway. If you created an API using quick create, the $default route is managed by API Gateway. You can't modify the $default route key.|
|api_key_required|boolean|Specifies whether an API key is required for this route. Supported only for WebSocket APIs.|
|authorization_scopes|text[]|A list of authorization scopes configured on a route. The scopes are used with a JWT authorizer to authorize the method invocation. The authorization works by matching the route scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any route scope matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the route scope is configured, the client must provide an access token instead of an identity token for authorization purposes.|
|authorization_type|text|The authorization type for the route. For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer For HTTP APIs, valid values are NONE for open access, JWT for using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.|
|authorizer_id|text|The identifier of the Authorizer resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.|
|model_selection_expression|text|The model selection expression for the route. Supported only for WebSocket APIs.|
|operation_name|text|The operation name for the route.|
|request_models|jsonb|The request models for the route. Supported only for WebSocket APIs.|
|request_parameters|jsonb|The request parameters for the route. Supported only for WebSocket APIs.|
|route_id|text|The route ID.|
|route_response_selection_expression|text|The route response selection expression for the route. Supported only for WebSocket APIs.|
|target|text|The target for the route.|
## Relations
## Table: aws_apigatewayv2_api_route_responses
Represents a route response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_route_id|uuid|Unique ID of aws_apigatewayv2_api_routes table (FK)|
|route_response_key|text|Represents the route response key of a route response.|
|model_selection_expression|text|Represents the model selection expression of a route response. Supported only for WebSocket APIs.|
|response_models|jsonb|Represents the response models of a route response.|
|response_parameters|jsonb|Represents the response parameters of a route response.|
|route_response_id|text|Represents the identifier of a route response.|
## Table: aws_apigatewayv2_api_stages
Represents an API stage.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|stage_name|text|The name of the stage.|
|access_log_settings_destination_arn|text|The ARN of the CloudWatch Logs log group to receive access logs.|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId.|
|api_gateway_managed|boolean|Specifies whether a stage is managed by API Gateway. If you created an API using quick create, the $default stage is managed by API Gateway. You can't modify the $default stage.|
|auto_deploy|boolean|Specifies whether updates to an API automatically trigger a new deployment. The default value is false.|
|client_certificate_id|text|The identifier of a client certificate for a Stage. Supported only for WebSocket APIs.|
|created_date|timestamp without time zone|The timestamp when the stage was created.|
|route_settings_data_trace_enabled|boolean|Specifies whether (true) or not (false) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.|
|route_settings_detailed_metrics_enabled|boolean|Specifies whether detailed metrics are enabled.|
|route_settings_logging_level|text|Specifies the logging level for this route: INFO, ERROR, or OFF. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.|
|route_settings_throttling_burst_limit|integer|Specifies the throttling burst limit.|
|route_settings_throttling_rate_limit|float|Specifies the throttling rate limit.|
|deployment_id|text|The identifier of the Deployment that the Stage is associated with. Can't be updated if autoDeploy is enabled.|
|description|text|The description of the stage.|
|last_deployment_status_message|text|Describes the status of the last deployment of a stage. Supported only for stages with autoDeploy enabled.|
|last_updated_date|timestamp without time zone|The timestamp when the stage was last updated.|
|route_settings|jsonb|Route settings for the stage, by routeKey.|
|stage_variables|jsonb|A map that defines the stage variables for a stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
