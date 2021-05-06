package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53HostedZones() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_hosted_zones",
		Resolver:     fetchRoute53HostedZones,
		Multiplex:    client.AccountRegionMultiplex,
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
				Resolver: schema.PathResolver("Id"),
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
	resource.Set("tags", tags)
	return nil
}
