package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotBillingGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_billing_groups",
		Resolver:     fetchIotBillingGroups,
		Multiplex:    client.AccountRegionMultiplex,
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
				Name:        "things_in_group",
				Description: "Lists the things in the specified group",
				Type:        schema.TypeStringArray,
				Resolver:    resolveIotBillingGroupThingsInGroup,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIotBillingGroupTags,
			},
			{
				Name:        "arn",
				Description: "The ARN of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupArn"),
			},
			{
				Name:        "id",
				Description: "The ID of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupId"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("BillingGroupMetadata.CreationDate"),
			},
			{
				Name:        "name",
				Description: "The name of the billing group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingGroupName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BillingGroupProperties.BillingGroupDescription"),
			},
			{
				Name:        "version",
				Description: "The version of the billing group.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIotBillingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input iot.ListBillingGroupsInput
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListBillingGroups(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, g := range response.BillingGroups {
			group, err := svc.DescribeBillingGroup(ctx, &iot.DescribeBillingGroupInput{
				BillingGroupName: g.GroupName,
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
func resolveIotBillingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeBillingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeBillingGroupOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListThingsInBillingGroupInput{
		BillingGroupName: i.BillingGroupName,
	}

	var things []string
	for {
		response, err := svc.ListThingsInBillingGroup(ctx, &input, func(options *iot.Options) {
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
func resolveIotBillingGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.DescribeBillingGroupOutput)
	if !ok {
		return fmt.Errorf("expected *iot.DescribeBillingGroupOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.BillingGroupArn,
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
