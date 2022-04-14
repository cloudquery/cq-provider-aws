package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=instance_patches.hcl -domain=ssm -resource=instance_patches
func InstancePatches() *schema.Table {
	return &schema.Table{
		Name:          "aws_ssm_instance_patches",
		Description:   "Information about the state of a patch on a particular instance as it relates to the patch baseline used to patch the instance.",
		Resolver:      fetchSsmInstancePatches,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "kb_id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "instance_cq_id",
				Description: "Unique CloudQuery ID of aws_ssm_instances table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "classification",
				Description: "The classification of the patch, such as SecurityUpdates, Updates, and CriticalUpdates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "installed_time",
				Description: "The date/time the patch was installed on the instance",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "kb_id",
				Description: "The operating system-specific ID of the patch.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KBId"),
			},
			{
				Name:        "severity",
				Description: "The severity of the patchsuch as Critical, Important, and Moderate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the patch on the instance, such as INSTALLED or FAILED",
				Type:        schema.TypeString,
			},
			{
				Name:        "title",
				Description: "The title of the patch.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cve_ids",
				Description: "The IDs of one or more Common Vulnerabilities and Exposure (CVE) issues that are resolved by the patch.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CVEIds"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSsmInstancePatches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM
	instance := parent.Item.(types.InstanceInformation)
	params := ssm.DescribeInstancePatchesInput{InstanceId: instance.InstanceId}
	for {
		result, err := svc.DescribeInstancePatches(ctx, &params, func(o *ssm.Options) { o.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Patches
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
