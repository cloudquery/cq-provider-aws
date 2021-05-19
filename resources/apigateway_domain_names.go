package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayDomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_domain_names",
		Resolver:     fetchApigatewayDomainNames,
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
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "api_mapping_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name:     "mutual_tls_authentication_truststore_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreUri"),
			},
			{
				Name:     "mutual_tls_authentication_truststore_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreVersion"),
			},
			{
				Name:     "mutual_tls_authentication_truststore_warnings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreWarnings"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigateway_domain_name_configurations",
				Resolver: fetchApigatewayDomainNameConfigurations,
				Columns: []schema.Column{
					{
						Name:     "domain_name_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_gateway_domain_name",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_arn",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_name",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_upload_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "domain_name_status",
						Type: schema.TypeString,
					},
					{
						Name: "domain_name_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "endpoint_type",
						Type: schema.TypeString,
					},
					{
						Name: "hosted_zone_id",
						Type: schema.TypeString,
					},
					{
						Name: "security_policy",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayDomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchApigatewayDomainNameConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
