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

func IamUserPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_user_policies",
		Resolver:     fetchIamUserPolicies,
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
				Resolver: resolveIamUserPolicyPolicyDocument,
			},
			{
				Name: "policy_name",
				Type: schema.TypeString,
			},
			{
				Name: "user_name",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamUserPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM
	_, err := GetCredentialReport(ctx, svc)
	if err != nil {
		return err
	}
	meta.(*client.Client).ReportUsers = nil

	getPoliciesForUser := func(userName *string) error {
		config := iam.ListUserPoliciesInput{UserName: userName}
		for {
			output, err := svc.ListUserPolicies(ctx, &config)
			if err != nil {
				return err
			}
			for _, p := range output.PolicyNames {
				policyCfg := &iam.GetUserPolicyInput{PolicyName: &p, UserName: userName}
				policyResult, err := svc.GetUserPolicy(ctx, policyCfg)
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

	var config iam.ListUsersInput
	for {
		output, err := svc.ListUsers(ctx, &config)
		if err != nil {
			return err
		}
		for _, u := range output.Users {
			if err := getPoliciesForUser(u.UserName); err != nil {
				return err
			}
		}
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
func resolveIamUserPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetUserPolicyOutput)
	if !ok {
		return fmt.Errorf("not user policy")
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
	return resource.Set(c.Name, document)
}
