package resources

import (
	"context"
	"errors"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_groups",
		Resolver:     fetchIamGroups,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: resolveIamGroupPolicies,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "create_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "group_id",
				Type: schema.TypeString,
			},
			{
				Name: "group_name",
				Type: schema.TypeString,
			},
			{
				Name: "path",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			iamGroupPolicies(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Groups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Group)
	svc := meta.(*client.Client).Services().IAM
	config := iam.ListGroupPoliciesInput{
		GroupName: r.GroupName,
	}
	var policies []string
	for {
		response, err := svc.ListGroupPolicies(ctx, &config)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchEntity" {
				return nil
			}
			return err
		}
		policies = append(policies, response.PolicyNames...)
		if response.Marker == nil {
			break
		}
		config.Marker = response.Marker
	}
	return resource.Set(c.Name, policies)
}
