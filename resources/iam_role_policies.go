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

func IamRolePolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_role_policies",
		Resolver:     fetchIamRolePolicies,
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
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamRolePolicyPolicyDocument,
			},
			{
				Name: "policy_name",
				Type: schema.TypeString,
			},
			{
				Name: "role_name",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamRolePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM

	getPoliciesForRole := func(roleName *string) error {
		config := iam.ListRolePoliciesInput{
			RoleName: roleName,
		}
		for {
			output, err := svc.ListRolePolicies(ctx, &config)
			if err != nil {
				return err
			}
			for _, p := range output.PolicyNames {
				policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: roleName})
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

	var config iam.ListRolesInput
	for {
		response, err := svc.ListRoles(ctx, &config)
		if err != nil {
			return err
		}
		for _, role := range response.Roles {
			if err := getPoliciesForRole(role.RoleName); err != nil {
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
func resolveIamRolePolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetRolePolicyOutput)
	if !ok {
		return fmt.Errorf("not role policy")
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
