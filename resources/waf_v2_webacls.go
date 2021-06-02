package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafV2Webacls() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_v2_webacls",
		Resolver:     fetchWafV2Webacls,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name:     "default_action",
				Type:     schema.TypeJSON,
				Resolver: resolveWafV2webaclDefaultAction,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "visibility_config_cloud_watch_metrics_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
			},
			{
				Name:     "visibility_config_metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
			},
			{
				Name:     "visibility_config_sampled_requests_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
			},
			{
				Name: "capacity",
				Type: schema.TypeBigInt,
			},
			{
				Name: "custom_response_bodies",
				Type: schema.TypeJSON,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "label_namespace",
				Type: schema.TypeString,
			},
			{
				Name: "managed_by_firewall_manager",
				Type: schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_waf_v2_webacl_rules",
				Resolver: fetchWafV2WebaclRules,
				Columns: []schema.Column{
					{
						Name:     "webacl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclRuleStatement,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
					{
						Name:     "action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclRuleAction,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclRuleOverrideAction,
					},
					{
						Name:     "labels",
						Type:     schema.TypeStringArray,
						Resolver: resolveWafV2webaclRuleLabels,
					},
				},
			},
			{
				Name:     "aws_waf_v2_webacl_post_process_firewall_manager_rule_groups",
				Resolver: fetchWafV2WebaclPostProcessFirewallManagerRuleGroups,
				Columns: []schema.Column{
					{
						Name:     "webacl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclPostProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclPostProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
			{
				Name:     "aws_waf_v2_webacl_pre_process_firewall_manager_rule_groups",
				Resolver: fetchWafV2WebaclPreProcessFirewallManagerRuleGroups,
				Columns: []schema.Column{
					{
						Name:     "webacl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclPreProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafV2webaclPreProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafV2Webacls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2
	config := wafv2.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(options *wafv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
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

func resolveWafV2webaclDefaultAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL, ok := resource.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", resource.Item)
	}
	if webACL.DefaultAction == nil {
		return nil
	}
	data, err := json.Marshal(webACL.DefaultAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchWafV2WebaclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.Rules
	return nil
}

func resolveWafV2webaclRuleStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.Statement == nil {
		return nil
	}
	data, err := json.Marshal(rule.Statement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafV2webaclRuleAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.Action == nil {
		return nil
	}
	data, err := json.Marshal(rule.Action)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafV2webaclRuleOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(rule.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafV2webaclRuleLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	labels := make([]string, len(rule.RuleLabels))
	for i := range rule.RuleLabels {
		labels[i] = aws.ToString(rule.RuleLabels[i].Name)
	}
	return resource.Set(c.Name, labels)
}

func fetchWafV2WebaclPostProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.PostProcessFirewallManagerRuleGroups
	return nil
}

func resolveWafV2webaclPostProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafV2webaclPostProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchWafV2WebaclPreProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.PreProcessFirewallManagerRuleGroups
	return nil
}

func resolveWafV2webaclPreProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafV2webaclPreProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
