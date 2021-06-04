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

func WafRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_rules",
		Resolver:     fetchWafRules,
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
				Name: "rule_id",
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
		Relations: []*schema.Table{
			{
				Name:     "aws_waf_rule_predicates",
				Resolver: fetchWafRulePredicates,
				Columns: []schema.Column{
					{
						Name:     "rule_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "data_id",
						Type: schema.TypeString,
					},
					{
						Name: "negated",
						Type: schema.TypeBool,
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
func fetchWafRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRulesInput{}
	for {
		output, err := service.ListRules(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, ruleSum := range output.Rules {
			rule, err := service.GetRule(ctx, &waf.GetRuleInput{RuleId: ruleSum.RuleId}, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- rule.Rule
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func fetchWafRulePredicates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	rule, ok := parent.Item.(*types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", rule)
	}
	res <- rule.Predicates
	return nil
}
