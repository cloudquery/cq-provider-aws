package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamGroupPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_group_policies",
		Resolver:     fetchIamGroupPolicies,
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
				Name: "group_name",
				Type: schema.TypeString,
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamGroupPolicyPolicyDocument,
			},
			{
				Name: "policy_name",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM

	getGroupsPolicies := func(groupName *string) error {
		config := iam.ListGroupPoliciesInput{
			GroupName: groupName,
		}
		for {
			output, err := svc.ListGroupPolicies(ctx, &config)
			if err != nil {
				return err
			}
			for _, p := range output.PolicyNames {
				policyResult, err := svc.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{PolicyName: &p, GroupName: groupName})
				if err != nil {
					return err
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

	var config iam.ListGroupsInput
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		for _, group := range response.Groups {
			if err := getGroupsPolicies(group.GroupName); err != nil {
				return err
			}
		}
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveIamGroupPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetGroupPolicyOutput)
	if !ok {
		return fmt.Errorf("not group policy")
	}

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return err
	}

	var document map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	resource.Set(c.Name, document)
	return nil
}
