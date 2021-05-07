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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRoute53hostedZoneTags,
			},
			{
				Name: "caller_reference",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: resolveRoute53hostedZoneResourceID,
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
						Name:     "hosted_zone_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_records",
						Type:     schema.TypeStringArray,
						Resolver: resolveRoute53hostedZoneResourceRecordSetResourceRecords,
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
			},
			{
				Name:     "aws_route53_hosted_zone_traffic_policy_instances",
				Resolver: fetchRoute53HostedZoneTrafficPolicyInstances,
				Columns: []schema.Column{
					{
						Name:     "hosted_zone_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "policy_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "message",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "ttl",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("TTL"),
					},
					{
						Name: "traffic_policy_id",
						Type: schema.TypeString,
					},
					{
						Name: "traffic_policy_type",
						Type: schema.TypeString,
					},
					{
						Name: "traffic_policy_version",
						Type: schema.TypeInt,
					},
				},
			},
			{
				Name:     "aws_route53_hosted_zone_vpc_association_authorizations",
				Resolver: fetchRoute53HostedZoneVpcAssociationAuthorizations,
				Columns: []schema.Column{
					{
						Name:     "hosted_zone_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "vpc_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VPCId"),
					},
					{
						Name:     "vpc_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VPCRegion"),
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
	r := resource.Item.(types.HostedZone)
	svc := meta.(*client.Client).Services().Route53
	resourceId := strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
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
	resource.Set(c.Name, tags)
	return nil
}
func resolveRoute53hostedZoneResourceID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.HostedZone)
	resource.Set(c.Name, strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1))
	return nil
}
func fetchRoute53HostedZoneQueryLoggingConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.HostedZone)
	svc := meta.(*client.Client).Services().Route53
	resourceId := strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: &resourceId}
	for {
		response, err := svc.ListQueryLoggingConfigs(ctx, &config, func(options *route53.Options) {})
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
	r := parent.Item.(types.HostedZone)
	svc := meta.(*client.Client).Services().Route53
	resourceId := strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
	config := route53.ListResourceRecordSetsInput{HostedZoneId: &resourceId}

	response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {})
	if err != nil {
		return err
	}
	res <- response.ResourceRecordSets

	return nil
}
func resolveRoute53hostedZoneResourceRecordSetResourceRecords(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ResourceRecordSet)
	var recordSets []string
	for _, t := range r.ResourceRecords {
		recordSets = append(recordSets, *t.Value)
	}
	resource.Set(c.Name, recordSets)
	return nil
}
func fetchRoute53HostedZoneTrafficPolicyInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.HostedZone)
	resourceId := strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
	config := route53.ListTrafficPolicyInstancesByHostedZoneInput{HostedZoneId: &resourceId}
	svc := meta.(*client.Client).Services().Route53
	for {
		response, err := svc.ListTrafficPolicyInstancesByHostedZone(ctx, &config, func(o *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.TrafficPolicyInstances
		if aws.ToString(response.TrafficPolicyInstanceNameMarker) == "" {
			break
		}
		config.TrafficPolicyInstanceNameMarker = response.TrafficPolicyInstanceNameMarker
	}
	return nil
}

func fetchRoute53HostedZoneVpcAssociationAuthorizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.HostedZone)
	resourceId := strings.Replace(*r.Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
	config := route53.ListVPCAssociationAuthorizationsInput{HostedZoneId: &resourceId}
	svc := meta.(*client.Client).Services().Route53
	for {
		response, err := svc.ListVPCAssociationAuthorizations(ctx, &config, func(o *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.VPCs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
