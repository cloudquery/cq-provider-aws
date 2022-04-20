package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource protection_groups --config gen.hcl --output .
func ProtectionGroups() *schema.Table {
	return &schema.Table{
		Name:          "aws_shield_protection_groups",
		Description:   "A grouping of protected resources that you and Shield Advanced can monitor as a collective",
		Resolver:      fetchShieldProtectionGroups,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "aggregation",
				Description: "Defines how Shield combines resource data for the group in order to detect, mitigate, and report events  * Sum - Use the total traffic across the group This is a good choice for most cases",
				Type:        schema.TypeString,
			},
			{
				Name:        "members",
				Description: "The Amazon Resource Names (ARNs) of the resources to include in the protection group",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "pattern",
				Description: "The criteria to use to choose the protected resources for inclusion in the group",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The name of the protection group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtectionGroupId"),
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the protection group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtectionGroupArn"),
			},
			{
				Name:        "resource_type",
				Description: "The resource type to include in the protection group",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchShieldProtectionGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.ListProtectionGroupsInput{}
	for {
		output, err := svc.ListProtectionGroups(ctx, &config, func(o *shield.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.ProtectionGroups

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
