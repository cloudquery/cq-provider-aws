package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotThingGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_thing_groups",
		Resolver:     fetchIotThingGroups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "things_in_group",
				Description: "Lists the things in the specified group",
				Type:        schema.TypeStringArray,
				Resolver:    resolveIotThingGroupThingsInGroup,
			},
			{
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: resolveIotThingGroupPolicies,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIotThingGroupTags,
			},
			{
				Name:        "index_name",
				Description: "The dynamic thing group index name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "query_string",
				Description: "The dynamic thing group search query string.",
				Type:        schema.TypeString,
			},
			{
				Name:        "query_version",
				Description: "The dynamic thing group query version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The dynamic thing group status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The thing group ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupArn"),
			},
			{
				Name:        "id",
				Description: "The thing group ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupId"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ThingGroupMetadata.CreationDate"),
			},
			{
				Name:     "parent_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupMetadata.ParentGroupName"),
			},
			{
				Name:     "root_to_parent_thing_groups",
				Type:     schema.TypeJSON,
				Resolver: resolveIotThingGroupRootToParentThingGroups,
			},
			{
				Name:        "name",
				Description: "The name of the thing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingGroupName"),
			},
			{
				Name:     "attribute_payload_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingGroupProperties.AttributePayload.Attributes"),
			},
			{
				Name:     "attribute_payload_merge",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ThingGroupProperties.AttributePayload.Merge"),
			},
			{
				Name:     "thing_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupProperties.ThingGroupDescription"),
			},
			{
				Name:        "version",
				Description: "The version of the thing group.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIotThingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input iot.ListThingGroupsInput
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThingGroups(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, g := range response.ThingGroups {
			group, err := svc.DescribeThingGroup(ctx, &iot.DescribeThingGroupInput{
				ThingGroupName: g.GroupName,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- group
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func resolveIotThingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeThingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeThingGroupOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListThingsInThingGroupInput{
		ThingGroupName: i.ThingGroupName,
	}

	var things []string
	for {
		response, err := svc.ListThingsInThingGroup(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}

		things = append(things, response.Things...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, things)
}
func resolveIotThingGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeThingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeThingGroupOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListAttachedPoliciesInput{
		Target: i.ThingGroupArn,
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
func resolveIotThingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeThingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeThingGroupOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.ThingGroupArn,
	}
	tags := make(map[string]interface{})

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})

		if err != nil {
			return err
		}
		for _, t := range response.Tags {
			tags[*t.Key] = t.Value
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
func resolveIotThingGroupRootToParentThingGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeThingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeThingGroupOutput but got %T", resource.Item)
	}
	if i.ThingGroupMetadata == nil {
		return nil
	}

	data, err := json.Marshal(i.ThingGroupMetadata.RootToParentThingGroups)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, data)
}
