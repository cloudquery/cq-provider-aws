package resources

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SsoAdminInstances() *schema.Table {
	return &schema.Table{
		Name:         "aws_sso_admin_instances",
		Description:  "Provides information about the SSO instance.",
		Resolver:     fetchSsoAdminInstances,
		Multiplex:    client.AccountMultiplex,
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
				Name:        "identity_store_id",
				Description: "The identifier of the identity store that is connected to the SSO instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The ARN of the SSO instance under which the operation will be executed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceArn"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_sso_admin_instance_groups",
				Description: "A group object, which contains a specified group’s metadata and attributes.",
				Resolver:    fetchSsoAdminInstanceGroups,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_sso_admin_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "display_name",
						Description: "Contains the group’s display name value",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for a group in the identity store.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("GroupId"),
					},
				},
			},
			{
				Name:        "aws_sso_admin_instance_users",
				Description: "A user object, which contains a specified user’s metadata and attributes.",
				Resolver:    fetchSsoAdminInstanceUsers,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_sso_admin_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The identifier for a user in the identity store.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UserId"),
					},
					{
						Name:        "name",
						Description: "Contains the user’s user name value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UserName"),
					},
				},
			},
			{
				Name:        "aws_sso_admin_instance_permission_sets",
				Description: "An entity that contains IAM policies.",
				Resolver:    fetchSsoAdminInstancePermissionSets,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_sso_admin_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "inline_policy",
						Type:     schema.TypeJSON,
						Resolver: ResolveSsoAdminInstancePermissionSetInlinePolicy,
					},
					{
						Name:        "tags",
						Description: "tags of the instance",
						Type:        schema.TypeJSON,
						Resolver:    ResolveSsoAdminInstancePermissionSetTags,
					},
					{
						Name:        "created_date",
						Description: "The date that the permission set was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description of the PermissionSet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the permission set.",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The ARN of the permission set",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PermissionSetArn"),
					},
					{
						Name:        "relay_state",
						Description: "Used to redirect users within the application during the federation authentication process.",
						Type:        schema.TypeString,
					},
					{
						Name:        "session_duration",
						Description: "The length of time that the application user sessions are valid for in the ISO-8601 standard.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_sso_admin_instance_permission_set_account_assignments",
						Description: "The assignment that indicates a principal's limited access to a specified Amazon Web Services account with a specified permission set",
						Resolver:    fetchSsoAdminInstancePermissionSetAccountAssignments,
						Columns: []schema.Column{
							{
								Name:        "instance_permission_set_cq_id",
								Description: "Unique CloudQuery ID of aws_sso_admin_instance_permission_sets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "account_id",
								Description: "The identifier of the Amazon Web Services account.",
								Type:        schema.TypeString,
							},
							{
								Name:        "permission_set_arn",
								Description: "The ARN of the permission set",
								Type:        schema.TypeString,
							},
							{
								Name:        "principal_id",
								Description: "An identifier for an object in Amazon Web Services SSO, such as a user or group. PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6)",
								Type:        schema.TypeString,
							},
							{
								Name:        "principal_type",
								Description: "The entity type for which the assignment will be created.",
								Type:        schema.TypeString,
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

func fetchSsoAdminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ssoadmin.ListInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().SSOAdmin
	for {
		response, err := svc.ListInstances(ctx, &config, func(o *ssoadmin.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Instances
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchSsoAdminInstanceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", parent.Item)
	}
	config := identitystore.ListGroupsInput{
		IdentityStoreId: r.IdentityStoreId,
	}
	awsClient := meta.(*client.Client)
	svc := awsClient.Services().IdentityStore
	for {
		response, err := svc.ListGroups(ctx, &config, func(options *identitystore.Options) {
			options.Region = awsClient.Region
		})
		if err != nil {
			return err
		}
		res <- response.Groups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
func fetchSsoAdminInstanceUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", parent.Item)
	}
	config := identitystore.ListUsersInput{
		IdentityStoreId: r.IdentityStoreId,
	}
	awsClient := meta.(*client.Client)
	svc := awsClient.Services().IdentityStore
	for {
		response, err := svc.ListUsers(ctx, &config, func(options *identitystore.Options) {
			options.Region = awsClient.Region
		})
		if err != nil {
			return err
		}
		res <- response.Users
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
func fetchSsoAdminInstancePermissionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", parent.Item)
	}

	awsClient := meta.(*client.Client)
	svc := awsClient.Services().SSOAdmin
	config := ssoadmin.ListPermissionSetsInput{
		InstanceArn: r.InstanceArn,
	}
	for {
		response, err := svc.ListPermissionSets(ctx, &config, func(o *ssoadmin.Options) {
			o.Region = awsClient.Region
		})
		if err != nil {
			return err
		}

		for _, p := range response.PermissionSets {
			ps, err := svc.DescribePermissionSet(ctx, &ssoadmin.DescribePermissionSetInput{
				InstanceArn:      r.InstanceArn,
				PermissionSetArn: &p,
			}, func(options *ssoadmin.Options) {
				options.Region = awsClient.Region
			})
			if err != nil {
				return err
			}

			res <- *ps.PermissionSet
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func ResolveSsoAdminInstancePermissionSetInlinePolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ps, ok := resource.Item.(types.PermissionSet)
	if !ok {
		return fmt.Errorf("expected to have types.PermissionSet but got %T", resource.Item)
	}
	im, ok := resource.Parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", resource.Parent.Item)
	}
	awsClient := meta.(*client.Client)
	svc := awsClient.Services().SSOAdmin

	config := ssoadmin.GetInlinePolicyForPermissionSetInput{
		PermissionSetArn: ps.PermissionSetArn,
		InstanceArn:      im.InstanceArn,
	}
	response, err := svc.GetInlinePolicyForPermissionSet(ctx, &config, func(o *ssoadmin.Options) {
		o.Region = awsClient.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.InlinePolicy)
}
func ResolveSsoAdminInstancePermissionSetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ps, ok := resource.Item.(types.PermissionSet)
	if !ok {
		return fmt.Errorf("expected to have types.PermissionSet but got %T", resource.Item)
	}
	im, ok := resource.Parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", resource.Parent.Item)
	}

	awsClient := meta.(*client.Client)
	svc := awsClient.Services().SSOAdmin
	config := ssoadmin.ListTagsForResourceInput{
		InstanceArn: im.InstanceArn,
		ResourceArn: ps.PermissionSetArn,
	}
	tags := make(map[string]string)
	for {
		response, err := svc.ListTagsForResource(ctx, &config, func(o *ssoadmin.Options) {
			o.Region = awsClient.Region
		})
		if err != nil {
			return err
		}

		for _, t := range response.Tags {
			tags[*t.Key] = *t.Value
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
func fetchSsoAdminInstancePermissionSetAccountAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	ps, ok := parent.Item.(types.PermissionSet)
	if !ok {
		return fmt.Errorf("expected to have types.PermissionSet but got %T", parent.Item)
	}
	im, ok := parent.Parent.Item.(types.InstanceMetadata)
	if !ok {
		return fmt.Errorf("expected to have types.InstanceMetadata but got %T", parent.Parent.Item)
	}
	awsClient := meta.(*client.Client)
	svc := awsClient.Services().SSOAdmin

	config := ssoadmin.ListAccountAssignmentsInput{
		AccountId:        &awsClient.AccountID,
		PermissionSetArn: ps.PermissionSetArn,
		InstanceArn:      im.InstanceArn,
	}
	for {
		response, err := svc.ListAccountAssignments(ctx, &config, func(o *ssoadmin.Options) {
			o.Region = awsClient.Region
		})
		if err != nil {
			return err
		}
		res <- response.AccountAssignments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
