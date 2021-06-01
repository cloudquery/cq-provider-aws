package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafWebacls() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_webacls",
		Resolver:     fetchWafWebacls,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:     "default_action_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultAction.Type"),
			},
			{
				Name:     "web_acl_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLId"),
			},
			{
				Name: "metric_name",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "web_acl_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLArn"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_waf_webacl_rules",
				Resolver: fetchWafWebaclRules,
				Columns: []schema.Column{
					{
						Name:     "webacl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name: "rule_id",
						Type: schema.TypeString,
					},
					{
						Name:     "action_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Action.Type"),
					},
					{
						Name:     "excluded_rules",
						Type:     schema.TypeStringArray,
						Resolver: resolveWafWebaclRuleExcludedRules,
					},
					{
						Name:     "override_action_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("OverrideAction.Type"),
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafWebacls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := waf.GetWebACLInput{WebACLId: webAcl.WebACLId}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- webAclOutput.WebACL
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func fetchWafWebaclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.Rules
	return nil
}

func resolveWafWebaclRuleExcludedRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.ActivatedRule)
	if !ok {
		return fmt.Errorf("not an ActivatedRule instance")
	}
	excludedRules := make([]string, len(rule.ExcludedRules))
	for i := range rule.ExcludedRules {
		excludedRules[i] = aws.ToString(rule.ExcludedRules[i].RuleId)
	}
	return resource.Set(c.Name, excludedRules)
}
