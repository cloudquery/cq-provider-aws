package ssm

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

//go:generate cq-gen -config=patch_baselines.hcl -domain=ssm -resource=patch_baselines
func PatchBaselines() *schema.Table {
	return &schema.Table{
		Name:         "aws_ssm_patch_baselines",
		Resolver:     fetchSsmPatchBaselines,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ssm"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "ARN of the resource.",
				Type:        schema.TypeString,
				Resolver:    resolvePatchBaselineARN,
			},
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
				Name:        "approved_patches",
				Description: "A list of explicitly approved patches for the baseline.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "approved_patches_compliance_level",
				Description: "Returns the specified compliance severity level for approved patches in the patch baseline.",
				Type:        schema.TypeString,
			},
			{
				Name:        "approved_patches_enable_non_security",
				Description: "Indicates whether the list of approved patches includes non-security updates that should be applied to the instances",
				Type:        schema.TypeBool,
			},
			{
				Name:        "baseline_id",
				Description: "The ARN of the retrieved patch baseline.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_date",
				Description: "The date the patch baseline was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A description of the patch baseline.",
				Type:        schema.TypeString,
			},
			{
				Name:        "global_filters",
				Description: "A set of global filters used to exclude patches from the baseline.",
				Type:        schema.TypeJSON,
				Resolver:    resolvePatchBaselineGlobalFilters,
			},
			{
				Name:        "modified_date",
				Description: "The date the patch baseline was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the patch baseline.",
				Type:        schema.TypeString,
			},
			{
				Name:        "operating_system",
				Description: "Returns the operating system specified for the patch baseline.",
				Type:        schema.TypeString,
			},
			{
				Name:        "patch_groups",
				Description: "Patch groups included in the patch baseline.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "rejected_patches",
				Description: "A list of explicitly rejected patches for the baseline.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "rejected_patches_action",
				Description: "The action specified to take on patches included in the RejectedPatches list",
				Type:        schema.TypeString,
			},
			{
				Name:          "tags",
				Description:   "Resource tags.",
				Type:          schema.TypeJSON,
				Resolver:      resolvePatchBaselineTags,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ssm_patch_baseline_approval_rules",
				Description: "Defines an approval rule for a patch baseline.",
				Resolver:    fetchSsmPatchBaselineApprovalRules,
				Columns: []schema.Column{
					{
						Name:        "patch_baseline_cq_id",
						Description: "Unique CloudQuery ID of aws_ssm_patch_baselines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "approve_after_days",
						Description:   "The number of days after the release date of each patch matched by the rule that the patch is marked as approved in the patch baseline",
						Type:          schema.TypeInt,
						IgnoreInTests: true,
					},
					{
						Name:          "approve_until_date",
						Description:   "The cutoff date for auto approval of released patches",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "compliance_level",
						Description: "A compliance severity level for all approved patches in a patch baseline.",
						Type:        schema.TypeString,
					},
					{
						Name:        "enable_non_security",
						Description: "For instances identified by the approval rule filters, enables a patch baseline to apply non-security updates available in the specified repository",
						Type:        schema.TypeBool,
					},
					{
						Name:        "patch_filter_group",
						Description: "The patch filter group that defines the criteria for the rule.",
						Type:        schema.TypeJSON,
						Resolver:    resolveApprovalRulePatchFilterGroup,
					},
				},
			},
			{
				Name:          "aws_ssm_patch_baseline_sources",
				Description:   "Information about the patches to use to update the instances, including target operating systems and source repository",
				Resolver:      fetchSsmPatchBaselineSources,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "patch_baseline_cq_id",
						Description: "Unique CloudQuery ID of aws_ssm_patch_baselines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "configuration",
						Description: "The value of the yum repo configuration",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name specified to identify the patch source.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "products",
						Description: "The specific operating system versions a patch repository applies to, such as \"Ubuntu16.04\", \"AmazonLinux2016.09\", \"RedhatEnterpriseLinux7.2\" or \"Suse12.7\". For lists of supported product values, see PatchFilter.  This member is required.",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

const patchBaselinesConcurrency = 5

func fetchSsmPatchBaselines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM
	params := ssm.DescribePatchBaselinesInput{}
	g, ctx := errgroup.WithContext(ctx)
	s := semaphore.NewWeighted(patchBaselinesConcurrency)
	for {
		result, err := svc.DescribePatchBaselines(ctx, &params, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range result.BaselineIdentities {
			id := item.BaselineId
			g.Go(func() error {
				if err := s.Acquire(ctx, 1); err != nil {
					return diag.WrapError(err)
				}
				defer s.Release(1)
				b, err := svc.GetPatchBaseline(
					ctx,
					&ssm.GetPatchBaselineInput{
						BaselineId: id,
					},
					func(o *ssm.Options) {
						o.Region = cl.Region
					},
				)
				if err != nil {
					return diag.WrapError(err)
				}
				if b != nil {
					res <- *b
				}
				return nil
			})
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return g.Wait()
}
func fetchSsmPatchBaselineApprovalRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	b := parent.Item.(ssm.GetPatchBaselineOutput)
	if b.ApprovalRules != nil {
		res <- b.ApprovalRules.PatchRules
	}
	return nil
}
func fetchSsmPatchBaselineSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	b := parent.Item.(ssm.GetPatchBaselineOutput)
	res <- b.Sources
	return nil
}
func resolvePatchBaselineGlobalFilters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	b := resource.Item.(ssm.GetPatchBaselineOutput)
	if b.GlobalFilters == nil {
		return nil
	}
	data, err := json.Marshal(b.GlobalFilters.PatchFilters)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveApprovalRulePatchFilterGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.PatchRule)
	if rule.PatchFilterGroup == nil {
		return nil
	}
	data, err := json.Marshal(rule.PatchFilterGroup.PatchFilters)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolvePatchBaselineTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM
	baseline := resource.Item.(ssm.GetPatchBaselineOutput)
	if strings.HasPrefix(*baseline.BaselineId, "arn:") {
		// ListTagsForResource returns 400 InvalidResourceId for builtin baselines.
		// And builtin baselines have baseline id in the ARN format.
		return nil
	}
	input := ssm.ListTagsForResourceInput{
		ResourceId:   baseline.BaselineId,
		ResourceType: types.ResourceTypeForTaggingPatchBaseline,
	}
	result, err := svc.ListTagsForResource(ctx, &input, func(o *ssm.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	b, err := json.Marshal(client.TagsToMap(result.TagList))
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func resolvePatchBaselineARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	baseline := resource.Item.(ssm.GetPatchBaselineOutput)
	arn := *baseline.BaselineId
	if !strings.HasPrefix(*baseline.BaselineId, "arn:") {
		arn = patchBaselineARN(meta, *baseline.BaselineId)
	}
	return diag.WrapError(resource.Set(c.Name, arn))
}
func patchBaselineARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return client.MakeARN(client.SSMService, cl.AccountID, cl.Region, "patchbaseline", id)
}
