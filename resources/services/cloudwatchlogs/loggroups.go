package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Loggroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudwatchlogs_loggroups",
		Description:  "Represents a log group.",
		Resolver:     fetchCloudwatchlogsLoggroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("logs"),
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
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:          "kms_key_id",
				Description:   "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "log_group_name",
				Description: "The name of the log group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "metric_filter_count",
				Description: "The number of metric filters.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:          "retention_in_days",
				Description:   "The number of days to retain the log events in the specified log group",
				Type:          schema.TypeBigInt,
				IgnoreInTests: true,
			},
			{
				Name:        "stored_bytes",
				Description: "The number of bytes stored.",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_cloudwatchlogs_logstreams",
				Description: "Represents a log stream, which is a sequence of log events from a single emitter of logs.",
				Resolver:    fetchCloudwatchlogsLogstreams,
				Columns: []schema.Column{
					{
						Name:        "loggroup_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudwatchlogs_loggroups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the log stream.",
						Type:        schema.TypeString,
					},
					{
						Name:        "creation_time",
						Description: "The creation time of the stream, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "first_event_timestamp",
						Description: "The time of the first event, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "last_event_timestamp",
						Description: "The time of the most recent log event in the log stream in CloudWatch Logs",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "last_ingestion_time",
						Description: "The ingestion time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "log_stream_name",
						Description: "The name of the log stream.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stored_bytes",
						Description: "The number of bytes stored",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "upload_sequence_token",
						Description: "The sequence token.",
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

func fetchCloudwatchlogsLoggroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs
	for {
		response, err := svc.DescribeLogGroups(ctx, &config, func(options *cloudwatchlogs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.LogGroups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchCloudwatchlogsLogstreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs
	logGroup := parent.Item.(types.LogGroup)
	config := cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: logGroup.LogGroupName,
	}
	for {
		response, err := svc.DescribeLogStreams(ctx, &config, func(options *cloudwatchlogs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.LogStreams
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
