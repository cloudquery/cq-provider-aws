package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource triggers --config triggers.hcl --output .
func Triggers() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_triggers",
		Description:  "Information about a specific trigger.",
		Resolver:     fetchGlueTriggers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the trigger.",
				Type:        schema.TypeString,
				Resolver:    resolveGlueTriggerArn,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueTriggerTags,
			},
			{
				Name:        "description",
				Description: "A description of this trigger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "event_batching_condition_size",
				Description: "Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.  This member is required.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("EventBatchingCondition.BatchSize"),
			},
			{
				Name:          "event_batching_condition_window",
				Description:   "Window of time in seconds after which EventBridge event trigger fires",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("EventBatchingCondition.BatchWindow"),
				IgnoreInTests: true,
			},
			{
				Name:          "id",
				Description:   "Reserved for future use.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "name",
				Description: "The name of the trigger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "predicate_logical",
				Description: "An optional field if only one condition is listed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Predicate.Logical"),
			},
			{
				Name:          "schedule",
				Description:   "A cron expression used to specify the schedule (see Time-Based Schedules for Jobs and Crawlers (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *).",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "The current state of the trigger.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of trigger that this is.",
				Type:        schema.TypeString,
			},
			{
				Name:          "workflow_name",
				Description:   "The name of the workflow associated with the trigger.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_trigger_actions",
				Description: "Defines an action to be initiated by a trigger.",
				Resolver:    fetchGlueTriggerActions,
				Columns: []schema.Column{
					{
						Name:        "trigger_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_triggers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "arguments",
						Description:   "The job arguments used when this trigger fires",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:          "crawler_name",
						Description:   "The name of the crawler to be used with this action.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "job_name",
						Description: "The name of a job to be run.",
						Type:        schema.TypeString,
					},
					{
						Name:          "notify_delay_after",
						Description:   "After a job run starts, the number of minutes to wait before sending a job run delay notification.",
						Type:          schema.TypeBigInt,
						Resolver:      schema.PathResolver("NotificationProperty.NotifyDelayAfter"),
						IgnoreInTests: true,
					},
					{
						Name:          "security_configuration",
						Description:   "The name of the SecurityConfiguration structure to be used with this action.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "timeout",
						Description:   "The JobRun timeout in minutes",
						Type:          schema.TypeBigInt,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "aws_glue_trigger_predicate_conditions",
				Description:   "Defines a condition under which a trigger fires.",
				Resolver:      fetchGlueTriggerPredicateConditions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "trigger_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_triggers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "crawl_state",
						Description: "The state of the crawler to which this condition applies.",
						Type:        schema.TypeString,
					},
					{
						Name:        "crawler_name",
						Description: "The name of the crawler to which this condition applies.",
						Type:        schema.TypeString,
					},
					{
						Name:        "job_name",
						Description: "The name of the job whose JobRuns this condition applies to, and on which this trigger waits.",
						Type:        schema.TypeString,
					},
					{
						Name:        "logical_operator",
						Description: "A logical operator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "The condition state",
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

func fetchGlueTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListTriggersInput{MaxResults: aws.Int32(200)}
	for {
		result, err := svc.ListTriggers(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, name := range result.TriggerNames {
			tr, err := svc.GetTrigger(ctx, &glue.GetTriggerInput{Name: aws.String(name)})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			if tr.Trigger != nil {
				res <- *tr.Trigger
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueTriggerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name)))
	return diag.WrapError(resource.Set(c.Name, arn))
}
func resolveGlueTriggerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
}
func fetchGlueTriggerActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	tr := parent.Item.(types.Trigger)
	res <- tr.Actions
	return nil
}
func fetchGlueTriggerPredicateConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	tr := parent.Item.(types.Trigger)
	if tr.Predicate != nil {
		res <- tr.Predicate.Conditions
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func triggerARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "trigger", name)
}
