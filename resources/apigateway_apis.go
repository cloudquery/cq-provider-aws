package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayApis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_apis",
		Resolver:     fetchApigatewayApis,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "protocol_type",
				Type: schema.TypeString,
			},
			{
				Name: "route_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name: "api_endpoint",
				Type: schema.TypeString,
			},
			{
				Name: "api_gateway_managed",
				Type: schema.TypeBool,
			},
			{
				Name: "api_id",
				Type: schema.TypeString,
			},
			{
				Name: "api_key_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name:     "cors_configuration_allow_credentials",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CorsConfiguration.AllowCredentials"),
			},
			{
				Name:     "cors_configuration_allow_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowHeaders"),
			},
			{
				Name:     "cors_configuration_allow_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowMethods"),
			},
			{
				Name:     "cors_configuration_allow_origins",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowOrigins"),
			},
			{
				Name:     "cors_configuration_expose_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.ExposeHeaders"),
			},
			{
				Name:     "cors_configuration_max_age",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CorsConfiguration.MaxAge"),
			},
			{
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disable_execute_api_endpoint",
				Type: schema.TypeBool,
			},
			{
				Name: "disable_schema_validation",
				Type: schema.TypeBool,
			},
			{
				Name: "import_info",
				Type: schema.TypeStringArray,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "version",
				Type: schema.TypeString,
			},
			{
				Name: "warnings",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigateway_api_mappings",
				Resolver: fetchApigatewayApiMappings,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_id",
						Type: schema.TypeString,
					},
					{
						Name: "stage",
						Type: schema.TypeString,
					},
					{
						Name: "api_mapping_id",
						Type: schema.TypeString,
					},
					{
						Name: "api_mapping_key",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_api_authorizers",
				Resolver: fetchApigatewayApiAuthorizers,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_credentials_arn",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_id",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_payload_format_version",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_result_ttl_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "authorizer_type",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_uri",
						Type: schema.TypeString,
					},
					{
						Name: "enable_simple_responses",
						Type: schema.TypeBool,
					},
					{
						Name: "identity_source",
						Type: schema.TypeStringArray,
					},
					{
						Name: "identity_validation_expression",
						Type: schema.TypeString,
					},
					{
						Name:     "jwt_configuration_audience",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("JwtConfiguration.Audience"),
					},
					{
						Name:     "jwt_configuration_issuer",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("JwtConfiguration.Issuer"),
					},
				},
			},
			{
				Name:     "aws_apigateway_api_deployments",
				Resolver: fetchApigatewayApiDeployments,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "auto_deployed",
						Type: schema.TypeBool,
					},
					{
						Name: "created_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "deployment_id",
						Type: schema.TypeString,
					},
					{
						Name: "deployment_status",
						Type: schema.TypeString,
					},
					{
						Name: "deployment_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_api_integrations",
				Resolver: fetchApigatewayApiIntegrations,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "connection_id",
						Type: schema.TypeString,
					},
					{
						Name: "connection_type",
						Type: schema.TypeString,
					},
					{
						Name: "content_handling_strategy",
						Type: schema.TypeString,
					},
					{
						Name: "credentials_arn",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "integration_id",
						Type: schema.TypeString,
					},
					{
						Name: "integration_method",
						Type: schema.TypeString,
					},
					{
						Name: "integration_response_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "integration_subtype",
						Type: schema.TypeString,
					},
					{
						Name: "integration_type",
						Type: schema.TypeString,
					},
					{
						Name: "integration_uri",
						Type: schema.TypeString,
					},
					{
						Name: "passthrough_behavior",
						Type: schema.TypeString,
					},
					{
						Name: "payload_format_version",
						Type: schema.TypeString,
					},
					{
						Name: "request_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "request_templates",
						Type: schema.TypeJSON,
					},
					{
						Name: "response_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "template_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "timeout_in_millis",
						Type: schema.TypeInt,
					},
					{
						Name:     "tls_config_server_name_to_verify",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TlsConfig.ServerNameToVerify"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_apigateway_api_integration_responses",
						Resolver: fetchApigatewayApiIntegrationResponses,
						Columns: []schema.Column{
							{
								Name:     "api_integration_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "integration_response_key",
								Type: schema.TypeString,
							},
							{
								Name: "content_handling_strategy",
								Type: schema.TypeString,
							},
							{
								Name: "integration_response_id",
								Type: schema.TypeString,
							},
							{
								Name: "response_parameters",
								Type: schema.TypeJSON,
							},
							{
								Name: "response_templates",
								Type: schema.TypeJSON,
							},
							{
								Name: "template_selection_expression",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_apigateway_api_models",
				Resolver: fetchApigatewayApiModels,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayAPIModelModelTemplate,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "content_type",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "model_id",
						Type: schema.TypeString,
					},
					{
						Name: "schema",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_api_routes",
				Resolver: fetchApigatewayApiRoutes,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "route_key",
						Type: schema.TypeString,
					},
					{
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "api_key_required",
						Type: schema.TypeBool,
					},
					{
						Name: "authorization_scopes",
						Type: schema.TypeStringArray,
					},
					{
						Name: "authorization_type",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_id",
						Type: schema.TypeString,
					},
					{
						Name: "model_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "operation_name",
						Type: schema.TypeString,
					},
					{
						Name: "request_models",
						Type: schema.TypeJSON,
					},
					{
						Name: "request_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "route_id",
						Type: schema.TypeString,
					},
					{
						Name: "route_response_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "target",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_apigateway_api_route_responses",
						Resolver: fetchApigatewayApiRouteResponses,
						Columns: []schema.Column{
							{
								Name:     "api_route_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "route_response_key",
								Type: schema.TypeString,
							},
							{
								Name: "model_selection_expression",
								Type: schema.TypeString,
							},
							{
								Name: "response_models",
								Type: schema.TypeJSON,
							},
							{
								Name: "response_parameters",
								Type: schema.TypeJSON,
							},
							{
								Name: "route_response_id",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_apigateway_api_stages",
				Resolver: fetchApigatewayApiStages,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "stage_name",
						Type: schema.TypeString,
					},
					{
						Name:     "access_log_settings_destination_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccessLogSettings.DestinationArn"),
					},
					{
						Name:     "access_log_settings_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccessLogSettings.Format"),
					},
					{
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "auto_deploy",
						Type: schema.TypeBool,
					},
					{
						Name: "client_certificate_id",
						Type: schema.TypeString,
					},
					{
						Name: "created_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "default_route_settings_data_trace_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("DefaultRouteSettings.DataTraceEnabled"),
					},
					{
						Name:     "default_route_settings_detailed_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("DefaultRouteSettings.DetailedMetricsEnabled"),
					},
					{
						Name:     "default_route_settings_logging_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DefaultRouteSettings.LoggingLevel"),
					},
					{
						Name:     "default_route_settings_throttling_burst_limit",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("DefaultRouteSettings.ThrottlingBurstLimit"),
					},
					{
						Name:     "default_route_settings_throttling_rate_limit",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("DefaultRouteSettings.ThrottlingRateLimit"),
					},
					{
						Name: "deployment_id",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "last_deployment_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "last_updated_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "route_settings",
						Type: schema.TypeJSON,
					},
					{
						Name: "stage_variables",
						Type: schema.TypeJSON,
					},
					{
						Name: "tags",
						Type: schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiIntegrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func resolveApigatewayAPIModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchApigatewayApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
