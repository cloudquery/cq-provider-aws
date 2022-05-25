package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource groups_access_advisor_details --config gen.hcl --output .
func GroupsAccessAdvisorDetails() *schema.Table {
	return &schema.Table{
		Name:          "aws_iam_access_advisor_details",
		Resolver:      fetchIamGroupsAccessAdvisorDetails,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		IgnoreInTests: true,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "service_namespace"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name: "parent_type",
				Type: schema.TypeString,
			},
			{
				Name:        "service_name",
				Description: "The name of the service in which access was attempted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceLastAccessed.ServiceName"),
			},
			{
				Name:        "service_namespace",
				Description: "The namespace of the service in which access was attempted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceLastAccessed.ServiceNamespace"),
			},
			{
				Name:        "last_authenticated",
				Description: "The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when an authenticated entity most recently attempted to access the service",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServiceLastAccessed.LastAuthenticated"),
			},
			{
				Name:        "last_authenticated_entity",
				Description: "The ARN of the authenticated entity (user or role) that last attempted to access the service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceLastAccessed.LastAuthenticatedEntity"),
			},
			{
				Name:        "last_authenticated_region",
				Description: "The Region from which the authenticated entity (user or role) last attempted to access the service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceLastAccessed.LastAuthenticatedRegion"),
			},
			{
				Name:        "total_authenticated_entities",
				Description: "The total number of authenticated principals (root user, IAM users, or IAM roles) that have attempted to access the service",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServiceLastAccessed.TotalAuthenticatedEntities"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iam_groups_access_advisor_detail_tracked_actions_last_accessed",
				Description: "Contains details about the most recent attempt to access an action within the service",
				Resolver:    fetchIamAccessAdvisorTrackedActionsLastAccesseds,
				Columns: []schema.Column{
					{
						Name:        "groups_access_advisor_detail_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_groups_access_advisor_details table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action_name",
						Description: "The name of the tracked action to which access was attempted",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_accessed_entity",
						Description: "The Amazon Resource Name (ARN)",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_accessed_region",
						Description: "The Region from which the authenticated entity (user or role) last attempted to access the tracked action",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_accessed_time",
						Description: "The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when an authenticated entity most recently attempted to access the tracked service",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "aws_iam_groups_access_advisor_detail_entities",
				Description: "An object that contains details about when the IAM entities (users or roles) were last used in an attempt to access the specified AWS service",
				Resolver:    fetchIamAccessAdvisorEntities,
				Columns: []schema.Column{
					{
						Name:        "groups_access_advisor_detail_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_groups_access_advisor_details table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EntityInfo.Arn"),
					},
					{
						Name:        "id",
						Description: "The identifier of the entity (user or role)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EntityInfo.Id"),
					},
					{
						Name:        "name",
						Description: "The name of the entity (user or role)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EntityInfo.Name"),
					},
					{
						Name:        "type",
						Description: "The type of entity (user or role)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EntityInfo.Type"),
					},
					{
						Name:        "path",
						Description: "The path to the entity (user or role)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EntityInfo.Path"),
					},
					{
						Name:        "last_authenticated",
						Description: "The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when the authenticated entity last attempted to access AWS",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIamGroupsAccessAdvisorDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, g := range response.Groups {
			err := fetchIamAccessDetails(ctx, res, svc, *g.Arn, GROUP)
			if err != nil {
				return diag.WrapError(err)
			}
			return nil
		}
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
