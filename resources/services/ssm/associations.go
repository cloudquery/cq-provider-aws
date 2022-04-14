package ssm

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=associations.hcl -domain=ssm -resource=associations
func Associations() *schema.Table {
	return &schema.Table{
		Name:         "aws_ssm_associations",
		Description:  "Describes the parameters for a document.",
		Resolver:     fetchSsmAssociations,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ssm"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "ARN of the association",
				Type:        schema.TypeString,
				Resolver:    resolveAssociationARN,
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
				Name:        "apply_only_at_cron_interval",
				Description: "By default, when you create a new associations, the system runs it immediately after it is created and then according to the schedule you specified",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The association ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssociationId"),
			},
			{
				Name:        "association_name",
				Description: "The association name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "association_version",
				Description: "The association version.",
				Type:        schema.TypeString,
			},
			{
				Name:          "automation_target_parameter_name",
				Description:   "Choose the parameter that will define how your automation will branch out",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "calendar_names",
				Description:   "The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your associations are gated under",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "compliance_severity",
				Description: "The severity level that is assigned to the association.",
				Type:        schema.TypeString,
			},
			{
				Name:        "date",
				Description: "The date when the association was made.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "document_version",
				Description: "The document version.",
				Type:        schema.TypeString,
			},
			{
				Name:          "instance_id",
				Description:   "The instance ID.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "last_execution_date",
				Description: "The date on which the association was last run.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_successful_execution_date",
				Description: "The last date on which the association was successfully run.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_update_association_date",
				Description: "The date when the association was last updated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "max_concurrency",
				Description: "The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_errors",
				Description: "The number of errors that are allowed before the system stops sending requests to run the association on additional targets",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the SSM document.",
				Type:        schema.TypeString,
			},
			{
				Name:          "output_location_s3_bucket_name",
				Description:   "The name of the S3 bucket.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("OutputLocation.S3Location.OutputS3BucketName"),
				IgnoreInTests: true,
			},
			{
				Name:          "output_location_s3_key_prefix",
				Description:   "The S3 bucket subfolder.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("OutputLocation.S3Location.OutputS3KeyPrefix"),
				IgnoreInTests: true,
			},
			{
				Name:          "output_location_s3_region",
				Description:   "The Amazon Web Services Region of the S3 bucket.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("OutputLocation.S3Location.OutputS3Region"),
				IgnoreInTests: true,
			},
			{
				Name:        "overview_association_status_aggregated_count",
				Description: "Returns the number of targets for the association status",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Overview.AssociationStatusAggregatedCount"),
			},
			{
				Name:        "overview_detailed_status",
				Description: "A detailed status of the association.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Overview.DetailedStatus"),
			},
			{
				Name:        "overview_status",
				Description: "The status of the association",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Overview.Status"),
			},
			{
				Name:          "parameters",
				Description:   "A description of the parameters for a document.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "schedule_expression",
				Description:   "A cron expression that specifies a schedule when the association runs.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "status_date",
				Description:   "The date when the status changed.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("Status.Date"),
				IgnoreInTests: true,
			},
			{
				Name:          "status_message",
				Description:   "The reason for the status.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Status.Message"),
				IgnoreInTests: true,
			},
			{
				Name:          "status_name",
				Description:   "The status.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Status.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "status_additional_info",
				Description:   "A user-defined string.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Status.AdditionalInfo"),
				IgnoreInTests: true,
			},
			{
				Name:          "sync_compliance",
				Description:   "The mode for generating association compliance",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "targets",
				Description: "The managed nodes targeted by the request.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAssociationTargets,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ssm_association_target_locations",
				Description:   "The combination of Amazon Web Services Regions and Amazon Web Services accounts targeted by the current Automation execution.",
				Resolver:      fetchSsmAssociationTargetLocations,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "association_cq_id",
						Description: "Unique CloudQuery ID of aws_ssm_associations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "accounts",
						Description: "The Amazon Web Services accounts targeted by the current Automation execution.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "execution_role_name",
						Description: "The Automation execution role used by the currently running Automation",
						Type:        schema.TypeString,
					},
					{
						Name:        "regions",
						Description: "The Amazon Web Services Regions targeted by the current Automation execution.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "target_location_max_concurrency",
						Description: "The maximum number of Amazon Web Services Regions and Amazon Web Services accounts allowed to run the Automation concurrently.",
						Type:        schema.TypeString,
					},
					{
						Name:        "target_location_max_errors",
						Description: "The maximum number of errors allowed before the system stops queueing additional Automation executions for the currently running Automation.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSsmAssociations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM
	params := ssm.ListAssociationsInput{}
	for {
		result, err := svc.ListAssociations(ctx, &params, func(o *ssm.Options) { o.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
		}
		for _, a := range result.Associations {
			desc, err := svc.DescribeAssociation(
				ctx,
				&ssm.DescribeAssociationInput{
					AssociationId:      a.AssociationId,
					AssociationVersion: a.AssociationVersion,
				},
				func(o *ssm.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
			}
			if desc.AssociationDescription != nil {
				res <- *desc.AssociationDescription
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
func fetchSsmAssociationTargetLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	a := parent.Item.(types.AssociationDescription)
	res <- a.TargetLocations
	return nil
}
func resolveAssociationTargets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	a := resource.Item.(types.AssociationDescription)
	b, err := json.Marshal(a.Targets)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func resolveAssociationARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, associationARN(meta, *resource.Item.(types.AssociationDescription).AssociationId)))
}
func associationARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return client.MakeARN(client.SSMService, cl.AccountID, cl.Region, "association", id)
}
