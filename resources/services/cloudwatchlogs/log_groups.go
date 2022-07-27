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

func CloudwatchlogsLogGroups() *schema.Table {
	return &schema.Table{
		Name:          "aws_cloudwatchlogs_log_groups",
		Description:   "CloudWatch Logs enables you to centralize the logs from all of your systems, applications, and AWS services that you use, in a single, highly scalable service. Log groups define groups of log streams that share the same retention, monitoring, and access control settings. Each log stream has to belong to one log group. ",
		Resolver:      fetchCloudwatchlogsLogGroups,
		Multiplex:     client.ServiceAccountRegionMultiplexer("logs"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
		IgnoreInTests: true,
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
				Description: "The ARN (Amazon Resource Name) for the group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KmsKeyId"),
			},
			{
				Name:        "name",
				Description: "The name of the log group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogGroupName"),
			},
			{
				Name:        "metric_filter_count",
				Description: "The number of metric filters",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "retention_in_days",
				Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, and 3653.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "stored_bytes",
				Description: "The number of bytes stored",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cloudwatchlogs_log_group_streams",
				Description:   "The Log Streams connected to a particular Log Group, each of which is a sequence of log events from a single emitter of logs",
				Resolver:      fetchCloudwatchlogsLogStreams,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "log_group_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudwatchlogs_log_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The ARN (Amazon Resource Name) for the stream.",
						Type:        schema.TypeString,
					},
					{
						Name:        "creation_time",
						Description: "The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "first_event_timestamp",
						Description: "The time of the first event, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "last_event_timestamp",
						Description: "The time of the most recent log event in the log stream in CloudWatch Logs. This number is expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC. ",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "last_ingestion_time",
						Description: "The ingestion time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "log_stream_name",
						Description: "The name of the CloudWatch metric.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stored_bytes",
						Description: "The number of bytes stored.",
						Type:        schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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

func fetchCloudwatchlogsLogStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.LogGroup)
	config := cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: r.LogGroupName,
	}

	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs

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
