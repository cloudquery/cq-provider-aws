package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource resource_groups --config gen.hcl --output .
func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:                 "aws_resourcegroups_resource_groups",
		Description:          "The unique identifiers for a resource group",
		Resolver:             fetchResourcegroupsResourceGroups,
		Multiplex:            client.ServiceAccountRegionMultiplexer("resourcegroups"),
		IgnoreError:          client.IgnoreCommonErrors,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveGroupQuery,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcegroupsResourceGroupTags,
			},
			{
				Name:        "resource_query_type",
				Description: "The type of the query.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_query",
				Description: "The query that defines a group or a search.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The ARN of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupArn"),
			},
			{
				Name:        "name",
				Description: "The name of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupName"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchResourcegroupsResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config resourcegroups.ListGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().ResourceGroups
	for {
		output, err := svc.ListGroups(ctx, &config, func(options *resourcegroups.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.GroupIdentifiers
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveGroupQuery(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	group := resource.Item.(types.GroupIdentifier)
	cl := meta.(*client.Client)
	svc := cl.Services().ResourceGroups
	input := resourcegroups.GetGroupQueryInput{
		Group: group.GroupArn,
	}
	output, err := svc.GetGroupQuery(ctx, &input, func(options *resourcegroups.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("resource_query_type", output.GroupQuery.ResourceQuery.Type); err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set("resource_query", output.GroupQuery.ResourceQuery.Query))
}

func resolveResourcegroupsResourceGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ResourceGroups
	group := resource.Item.(types.GroupIdentifier)
	input := resourcegroups.GetTagsInput{
		Arn: group.GroupArn,
	}
	output, err := svc.GetTags(ctx, &input, func(options *resourcegroups.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, output.Tags))
}
