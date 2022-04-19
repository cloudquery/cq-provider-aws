package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource protections --config gen.hcl --output .
func Protections() *schema.Table {
	return &schema.Table{
		Name:         "aws_shield_protections",
		Description:  "An object that represents a resource that is under DDoS protection.",
		Resolver:     fetchShieldProtections,
		Multiplex:    client.ServiceAccountRegionMultiplexer("shield"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The AWS tags of the resource.",
				Type:        schema.TypeString,
				Resolver:    ResolveShieldProtectionTags,
			},
			{
				Name:        "application_layer_automatic_response_configuration_status",
				Description: "Indicates whether automatic application layer DDoS mitigation is enabled for the protection.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationLayerAutomaticResponseConfiguration.Status"),
			},
			{
				Name:        "health_check_ids",
				Description: "The unique identifier (ID) for the Route 53 health check that's associated with the protection.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the protection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the protection",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the protection.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtectionArn"),
			},
			{
				Name:        "resource_arn",
				Description: "The ARN (Amazon Resource Name) of the Amazon Web Services resource that is protected.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchShieldProtections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.ListProtectionsInput{}
	for {
		output, err := svc.ListProtections(ctx, &config, func(o *shield.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Protections

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func ResolveShieldProtectionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Protection)
	client := meta.(*client.Client)
	svc := client.Services().Shield
	config := shield.ListTagsForResourceInput{ResourceARN: r.ProtectionArn}

	output, err := svc.ListTagsForResource(ctx, &config, func(o *shield.Options) {
		o.Region = client.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	tags := map[string]*string{}
	for _, t := range output.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
