package kinesis

import (
	"context"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Streams() *schema.Table {
	return &schema.Table{
		Name:         "aws_kinesis_streams",
		Description:  "Represents the output for DescribeStreamSummary",
		Resolver:     fetchKinesisStreams,
		Multiplex:    client.ServiceAccountRegionMultiplexer("logs"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "open_shard_count",
				Description: "The number of open shards in the stream.  This member is required.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "retention_period_hours",
				Description: "The current retention period, in hours.  This member is required.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) for the stream being described.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamARN"),
			},
			{
				Name:        "stream_creation_timestamp",
				Description: "The approximate time that the stream was created.  This member is required.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "stream_name",
				Description: "The name of the stream being described.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "stream_status",
				Description: "The current status of the stream being described",
				Type:        schema.TypeString,
			},
			{
				Name:        "consumer_count",
				Description: "The number of enhanced fan-out consumers registered with the stream.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "encryption_type",
				Description: "The encryption type used",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_id",
				Description: "The GUID for the customer-managed Amazon Web Services KMS key to use for encryption",
				Type:        schema.TypeString,
			},
			{
				Name:        "stream_mode_details_stream_mode",
				Description: "Specifies the capacity mode to which you want to set your data stream. Currently, in Kinesis Data Streams, you can choose between an on-demand capacity mode and a provisioned capacity mode for your data streams.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamModeDetails.StreamMode"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_kinesis_stream_enhanced_monitoring",
				Description: "Represents enhanced metrics types.",
				Resolver:    fetchKinesisStreamEnhancedMonitorings,
				Columns: []schema.Column{
					{
						Name:        "stream_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_streams table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "shard_level_metrics",
						Description: "List of shard-level metrics",
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

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisStreamEnhancedMonitorings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchKinesisStreamShards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
