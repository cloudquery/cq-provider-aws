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

func WafRuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_rule_groups",
		Resolver:     fetchWafRuleGroups,
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
				Name:     "rule_ids",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafRuleGroupRuleIds,
			},
			{
				Name: "rule_group_id",
				Type: schema.TypeString,
			},
			{
				Name: "metric_name",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRuleGroupsInput{}
	for {
		output, err := service.ListRuleGroups(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, rG := range output.RuleGroups {
			ruleGroup, err := service.GetRuleGroup(ctx, &waf.GetRuleGroupInput{RuleGroupId: rG.RuleGroupId}, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- ruleGroup.RuleGroup
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafRuleGroupRuleIds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not an RuleGroupInstance: %#v", resource.Item)
	}

	// Resolves rule group rules
	client := meta.(*client.Client)
	service := client.Services().Waf
	listActivatedRulesConfig := waf.ListActivatedRulesInRuleGroupInput{RuleGroupId: ruleGroup.RuleGroupId}
	var ruleIDs []string
	for {
		rules, err := service.ListActivatedRulesInRuleGroup(ctx, &listActivatedRulesConfig, func(options *waf.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}
		for _, rule := range rules.ActivatedRules {
			ruleIDs = append(ruleIDs, aws.ToString(rule.RuleId))
		}

		if aws.ToString(rules.NextMarker) == "" {
			break
		}
		listActivatedRulesConfig.NextMarker = rules.NextMarker
	}
	return resource.Set("rule_ids", ruleIDs)
}
