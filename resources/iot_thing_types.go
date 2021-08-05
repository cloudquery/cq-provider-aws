package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotThingTypes() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_thing_types",
		Description:  "The definition of the thing type, including thing type name and description.",
		Resolver:     fetchIotThingTypes,
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
				Name:        "arn",
				Description: "The thing type ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingTypeArn"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time when the thing type was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ThingTypeMetadata.CreationDate"),
			},
			{
				Name:        "deprecated",
				Description: "Whether the thing type is deprecated",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ThingTypeMetadata.Deprecated"),
			},
			{
				Name:        "deprecation_date",
				Description: "The date and time when the thing type was deprecated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ThingTypeMetadata.DeprecationDate"),
			},
			{
				Name:        "name",
				Description: "The name of the thing type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingTypeName"),
			},
			{
				Name:        "searchable_attributes",
				Description: "A list of searchable thing attribute names.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ThingTypeProperties.SearchableAttributes"),
			},
			{
				Name:        "description",
				Description: "The description of the thing type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingTypeProperties.ThingTypeDescription"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIotThingTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input iot.ListThingTypesInput
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThingTypes(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		res <- response.ThingTypes

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
