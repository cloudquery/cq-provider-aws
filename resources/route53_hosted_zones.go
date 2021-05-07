package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53HostedZones() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_hosted_zones",
		Resolver:     fetchRoute53HostedZones,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: resolveRoute53hostedZoneResourceID,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRoute53hostedZoneTags,
			},
			{
				Name: "caller_reference",
				Type: schema.TypeString,
			},

			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "config_comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Config.Comment"),
			},
			{
				Name:     "config_private_zone",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Config.PrivateZone"),
			},
			{
				Name:     "linked_service_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LinkedService.Description"),
			},
			{
				Name:     "linked_service_principal",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LinkedService.ServicePrincipal"),
			},
			{
				Name: "resource_record_set_count",
				Type: schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_route53_hosted_zone_query_logging_configs",
				Resolver: fetchRoute53HostedZoneQueryLoggingConfigs,
				Columns: []schema.Column{
					{
						Name:     "hosted_zone_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "cloud_watch_logs_log_group_arn",
						Type: schema.TypeString,
					},
					{
						Name:     "query_logging_config_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
				},
			},
			{
				Name:     "aws_route53_hosted_zone_resource_record_sets",
				Resolver: fetchRoute53HostedZoneResourceRecordSets,
				Columns: []schema.Column{
					{
						Name:     "hosted_zone_id2",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name:     "dns_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AliasTarget.DNSName"),
					},
					{
						Name:     "evaluate_target_health",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("AliasTarget.EvaluateTargetHealth"),
					},
					{
						Name:     "hosted_zone_id1",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AliasTarget.HostedZoneId"),
					},
					{
						Name: "failover",
						Type: schema.TypeString,
					},
					{
						Name:     "geo_location_continent_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GeoLocation.ContinentCode"),
					},
					{
						Name:     "geo_location_country_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GeoLocation.CountryCode"),
					},
					{
						Name:     "geo_location_subdivision_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GeoLocation.SubdivisionCode"),
					},
					{
						Name: "health_check_id",
						Type: schema.TypeString,
					},
					{
						Name: "multi_value_answer",
						Type: schema.TypeBool,
					},
					{
						Name: "region",
						Type: schema.TypeString,
					},
					{
						Name: "set_identifier",
						Type: schema.TypeString,
					},
					{
						Name:     "ttl",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("TTL"),
					},
					{
						Name: "traffic_policy_instance_id",
						Type: schema.TypeString,
					},
					{
						Name: "weight",
						Type: schema.TypeBigInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_route53_hosted_zone_resource_record_set_resource_records",
						Resolver: fetchRoute53HostedZoneResourceRecordSetResourceRecords,
						Columns: []schema.Column{
							{
								Name:     "hosted_zone_resource_record_set_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "value",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53HostedZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config route53.ListHostedZonesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListHostedZones(ctx, &config, func(o *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.HostedZones
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveRoute53hostedZoneTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Route53
	resourceId := resource.Get("resource_id").(string)
	tagsOutput, err := svc.ListTagsForResource(ctx, &route53.ListTagsForResourceInput{ResourceId: &resourceId, ResourceType: types.TagResourceTypeHostedzone}, func(options *route53.Options) {})
	if err != nil {
		return err
	}

	if tagsOutput.ResourceTagSet == nil {
		return nil
	}

	tags := map[string]*string{}
	for _, t := range tagsOutput.ResourceTagSet.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}

func resolveRoute53hostedZoneResourceID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.HostedZone)
	resource.Set(c.Name, strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1))
	return nil
}

func fetchRoute53HostedZoneQueryLoggingConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Route53
	resourceId := parent.Get("resource_id").(string)
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: &resourceId}
	for {
		response, err := svc.ListQueryLoggingConfigs(ctx, &route53.ListQueryLoggingConfigsInput{HostedZoneId: &resourceId}, func(options *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.QueryLoggingConfigs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchRoute53HostedZoneResourceRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}

func fetchRoute53HostedZoneResourceRecordSetResourceRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
