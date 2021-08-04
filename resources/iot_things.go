package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotThings() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_things",
		Description:  "The properties of the thing, including thing name, thing type name, and a list of thing attributes.",
		Resolver:     fetchIotThings,
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
				Name:     "principals",
				Type:     schema.TypeStringArray,
				Resolver: resolveIotThingPrincipals,
			},
			{
				Name:        "attributes",
				Description: "A list of thing attributes which are name-value pairs.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "thing_arn",
				Description: "The thing ARN.",
				Type:        schema.TypeString,
			},
			{
				Name:        "thing_name",
				Description: "The name of the thing.",
				Type:        schema.TypeString,
			},
			{
				Name:        "thing_type_name",
				Description: "The name of the thing type, if the thing has been associated with a type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "version",
				Description: "The version of the thing record in the registry.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIotThings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input iot.ListThingsInput
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThings(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Things
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func resolveIotThingPrincipals(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(types.ThingAttribute)
	if !ok {
		return fmt.Errorf("expected *types.ThingAttribute but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListThingPrincipalsInput{
		ThingName: i.ThingName,
	}
	var principals []string

	for {
		response, err := svc.ListThingPrincipals(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})

		if err != nil {
			return err
		}
		principals = append(principals, response.Principals...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, principals)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveIotThingTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(types.ThingAttribute)
	if !ok {
		return fmt.Errorf("expected *types.ThingAttribute but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.ThingArn,
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
