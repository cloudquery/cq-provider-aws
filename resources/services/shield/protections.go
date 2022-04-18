package shield

import (
	"context"
	"github.com/cloudquery/cq-provider-aws/client"
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
				Name:        "protection_arn",
				Description: "The ARN (Amazon Resource Name) of the protection.",
				Type:        schema.TypeString,
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
	panic("not implemented")
}
