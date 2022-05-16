package iam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource groups --config gen.hcl --output .
func Groups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_groups",
		Description:  "Contains information about an IAM group entity",
		Resolver:     fetchIamGroups,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "policies",
				Description: "List of policies attached to group.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIamGroupPolicies,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the group",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when the group was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupId"),
			},
			{
				Name:        "name",
				Description: "The friendly name that identifies the group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupName"),
			},
			{
				Name:        "path",
				Description: "The path to the group",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iam_group_policies",
				Description: "Contains the response to a successful GetGroupPolicy request",
				Resolver:    fetchIamGroupPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"group_cq_id", "policy_name"}},
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "group_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "account_id",
						Description: "The AWS Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSAccount,
					},
					{
						Name:        "group_name",
						Description: "The group the policy is associated with",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_document",
						Description: "The policy document",
						Type:        schema.TypeJSON,
						Resolver:    resolveGroupPoliciesPolicyDocument,
					},
					{
						Name:        "policy_name",
						Description: "The name of the policy",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_iam_group_accessed_details",
				Resolver: fetchIamGroupAccessedDetails,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_iam_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "aws_iam_group_accessed_detail_tracked_actions_last_accessed",
						Description: "Contains details about the most recent attempt to access an action within the service",
						Resolver:    fetchIamGroupAccessedDetailTrackedActionsLastAccesseds,
						Columns: []schema.Column{
							{
								Name:        "group_accessed_detail_cq_id",
								Description: "Unique CloudQuery ID of aws_iam_group_accessed_details table (FK)",
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
						Name:        "aws_iam_group_accessed_detail_entities",
						Description: "An object that contains details about when the IAM entities (users or roles) were last used in an attempt to access the specified AWS service",
						Resolver:    fetchIamGroupAccessedDetailEntities,
						Columns: []schema.Column{
							{
								Name:        "group_accessed_detail_cq_id",
								Description: "Unique CloudQuery ID of aws_iam_group_accessed_details table (FK)",
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
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Groups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func ResolveIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Group)
	svc := meta.(*client.Client).Services().IAM
	config := iam.ListAttachedGroupPoliciesInput{
		GroupName: r.GroupName,
	}
	response, err := svc.ListAttachedGroupPolicies(ctx, &config)
	if err != nil {
		return diag.WrapError(err)
	}
	policyMap := map[string]*string{}
	for _, p := range response.AttachedPolicies {
		policyMap[*p.PolicyArn] = p.PolicyName
	}
	return resource.Set(c.Name, policyMap)
}
func fetchIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	group := parent.Item.(types.Group)
	config := iam.ListGroupPoliciesInput{
		GroupName: group.GroupName,
	}
	for {
		output, err := svc.ListGroupPolicies(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		for _, p := range output.PolicyNames {
			policyResult, err := svc.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{PolicyName: &p, GroupName: group.GroupName})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- policyResult
		}
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
func resolveGroupPoliciesPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetGroupPolicyOutput)

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return diag.WrapError(err)
	}

	var document map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, document)
}
func fetchIamGroupAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	group := parent.Item.(types.Group)
	config := iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         group.Arn,
		Granularity: types.AccessAdvisorUsageGranularityTypeActionLevel,
	}
	output, err := svc.GenerateServiceLastAccessedDetails(ctx, &config)
	if err != nil {
		return diag.WrapError(err)
	}

	getDetails := iam.GetServiceLastAccessedDetailsInput{
		JobId: output.JobId,
	}
	for {
		details, err := svc.GetServiceLastAccessedDetails(ctx, &getDetails)
		if err != nil {
			return diag.WrapError(err)
		}

		switch details.JobStatus {
		case types.JobStatusTypeInProgress:
			time.Sleep(time.Millisecond * 200)
			continue
		case types.JobStatusTypeFailed:
			return diag.WrapError(fmt.Errorf("failed to get last acessed details with error: %s - %s", *details.Error.Code, *details.Error.Message))
		case types.JobStatusTypeCompleted:
			for _, s := range details.ServicesLastAccessed {
				if *s.TotalAuthenticatedEntities > 0 {
					res <- AccessedDetails{
						JobId:               output.JobId,
						ServiceLastAccessed: s,
					}
				}
			}
			if details.Marker == nil {
				return nil
			}
			if details.Marker != nil {
				getDetails.Marker = details.Marker
			}
		}
	}
}
func fetchIamGroupAccessedDetailTrackedActionsLastAccesseds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	serviceLastAccessed := parent.Item.(AccessedDetails)
	res <- serviceLastAccessed.TrackedActionsLastAccessed
	return nil
}
func fetchIamGroupAccessedDetailEntities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().IAM
	serviceLastAccessed := parent.Item.(AccessedDetails)
	config := iam.GetServiceLastAccessedDetailsWithEntitiesInput{
		JobId:            serviceLastAccessed.JobId,
		ServiceNamespace: serviceLastAccessed.ServiceNamespace,
	}
	for {
		output, err := svc.GetServiceLastAccessedDetailsWithEntities(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.EntityDetailsList
		if output.Marker == nil {
			break
		}
		if output.Marker != nil {
			config.Marker = output.Marker
		}
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type AccessedDetails struct {
	types.ServiceLastAccessed
	JobId *string
}
