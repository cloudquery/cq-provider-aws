package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DynamodbTables() *schema.Table {
	return &schema.Table{
		Name:        "aws_dynamodb_tables",
		Description: "Represents the properties of a table.",
		Resolver:    fetchDynamodbTables,
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
				Name:        "archival_summary",
				Description: "Contains information about the table archive.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableArchivalSummary,
			},
			{
				Name:        "attribute_definitions",
				Description: "An array of AttributeDefinition objects",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableAttributeDefinitions,
			},
			{
				Name:        "billing_mode_summary",
				Description: "Contains the details for the read/write capacity mode.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableBillingModeSummary,
			},
			{
				Name:        "creation_date_time",
				Description: "The date and time when the table was created, in UNIX epoch time (http://www.epochconverter.com/) format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "global_table_version",
				Description: "Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html) in use, if the table is replicated across Amazon Web Services Regions.",
				Type:        schema.TypeString,
			},
			{
				Name:        "item_count",
				Description: "The number of items in the specified table",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "key_schema",
				Description: "The primary key structure for the table",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableKeySchema,
			},
			{
				Name:        "latest_stream_arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_stream_label",
				Description: "A timestamp, in ISO 8601 format, for this stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioned_throughput_last_decrease_date_time",
				Description: "The date and time of the last provisioned throughput decrease for this table.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ProvisionedThroughput.LastDecreaseDateTime"),
			},
			{
				Name:        "provisioned_throughput_last_increase_date_time",
				Description: "The date and time of the last provisioned throughput increase for this table.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ProvisionedThroughput.LastIncreaseDateTime"),
			},
			{
				Name:        "provisioned_throughput_number_of_decreases_today",
				Description: "The number of provisioned throughput decreases for this table during this UTC calendar day",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ProvisionedThroughput.NumberOfDecreasesToday"),
			},
			{
				Name:        "provisioned_throughput_read_capacity_units",
				Description: "The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ProvisionedThroughput.ReadCapacityUnits"),
			},
			{
				Name:        "provisioned_throughput_write_capacity_units",
				Description: "The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ProvisionedThroughput.WriteCapacityUnits"),
			},
			{
				Name:        "restore_summary",
				Description: "Contains details for the restore.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableRestoreSummary,
			},
			{
				Name:        "inaccessible_encryption_date_time",
				Description: "Indicates the time, in UNIX epoch date format, when DynamoDB detected that the table's KMS key was inaccessible",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("SSEDescription.InaccessibleEncryptionDateTime"),
			},
			{
				Name:        "kms_master_key_arn",
				Description: "The KMS key ARN used for the KMS encryption.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SSEDescription.KMSMasterKeyArn"),
			},
			{
				Name:        "sse_type",
				Description: "Server-side encryption type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SSEDescription.SSEType"),
			},
			{
				Name:        "sse_status",
				Description: "Represents the current state of server-side encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SSEDescription.Status"),
			},
			{
				Name:        "stream_specification",
				Description: "The current DynamoDB Streams configuration for the table.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableStreamSpecification,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the table.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableArn"),
			},
			{
				Name:        "table_class_summary_last_update_date_time",
				Description: "The date and time at which the table class was last updated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("TableClassSummary.LastUpdateDateTime"),
			},
			{
				Name:        "table_class_summary_table_class",
				Description: "The table class of the specified table",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableClassSummary.TableClass"),
			},
			{
				Name:        "id",
				Description: "Unique identifier for the table for which the backup was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableId"),
			},
			{
				Name:        "name",
				Description: "The name of the table.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableName"),
			},
			{
				Name:        "table_size_bytes",
				Description: "The total size of the specified table, in bytes",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "table_status",
				Description: "The current state of the table:  * CREATING - The table is being created.  * UPDATING - The table is being updated.  * DELETING - The table is being deleted.  * ACTIVE - The table is ready for use.  * INACCESSIBLE_ENCRYPTION_CREDENTIALS - The KMS key used to encrypt the table in inaccessible",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_dynamodb_table_global_secondary_indexes",
				Description: "Represents the properties of a global secondary index.",
				Resolver:    fetchDynamodbTableGlobalSecondaryIndexes,
				Columns: []schema.Column{
					{
						Name:        "table_cq_id",
						Description: "Unique CloudQuery ID of aws_dynamodb_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "backfilling",
						Description: "Indicates whether the index is currently backfilling",
						Type:        schema.TypeBool,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) that uniquely identifies the index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IndexArn"),
					},
					{
						Name:        "name",
						Description: "The name of the global secondary index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IndexName"),
					},
					{
						Name:        "index_size_bytes",
						Description: "The total size of the specified index, in bytes",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "status",
						Description: "The current state of the global secondary index:  * CREATING - The index is being created.  * UPDATING - The index is being updated.  * DELETING - The index is being deleted.  * ACTIVE - The index is ready for use.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IndexStatus"),
					},
					{
						Name:        "item_count",
						Description: "The number of items in the specified index",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "key_schema",
						Description: "The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  * HASH - partition key  * RANGE - sort key  The partition key of an item is also known as its hash attribute",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableGlobalSecondaryIndexKeySchema,
					},
					{
						Name:        "projection_non_key_attributes",
						Description: "Represents the non-key attribute names which will be projected into the index. For local secondary indexes, the total count of NonKeyAttributes summed across all of the local secondary indexes, must not exceed 20",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Projection.NonKeyAttributes"),
					},
					{
						Name:        "projection_type",
						Description: "The set of attributes that are projected into the index:  * KEYS_ONLY - Only the index and primary keys are projected into the index.  * INCLUDE - In addition to the attributes described in KEYS_ONLY, the secondary index will include other non-key attributes that you specify.  * ALL - All of the table attributes are projected into the index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Projection.ProjectionType"),
					},
					{
						Name:        "provisioned_throughput_last_decrease_date_time",
						Description: "The date and time of the last provisioned throughput decrease for this table.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ProvisionedThroughput.LastDecreaseDateTime"),
					},
					{
						Name:        "provisioned_throughput_last_increase_date_time",
						Description: "The date and time of the last provisioned throughput increase for this table.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ProvisionedThroughput.LastIncreaseDateTime"),
					},
					{
						Name:        "provisioned_throughput_number_of_decreases_today",
						Description: "The number of provisioned throughput decreases for this table during this UTC calendar day",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ProvisionedThroughput.NumberOfDecreasesToday"),
					},
					{
						Name:        "provisioned_throughput_read_capacity_units",
						Description: "The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ProvisionedThroughput.ReadCapacityUnits"),
					},
					{
						Name:        "provisioned_throughput_write_capacity_units",
						Description: "The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ProvisionedThroughput.WriteCapacityUnits"),
					},
				},
			},
			{
				Name:        "aws_dynamodb_table_local_secondary_indexes",
				Description: "Represents the properties of a local secondary index.",
				Resolver:    fetchDynamodbTableLocalSecondaryIndexes,
				Columns: []schema.Column{
					{
						Name:        "table_cq_id",
						Description: "Unique CloudQuery ID of aws_dynamodb_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) that uniquely identifies the index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IndexArn"),
					},
					{
						Name:        "name",
						Description: "Represents the name of the local secondary index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IndexName"),
					},
					{
						Name:        "index_size_bytes",
						Description: "The total size of the specified index, in bytes",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "item_count",
						Description: "The number of items in the specified index",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "key_schema",
						Description: "The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:  * HASH - partition key  * RANGE - sort key  The partition key of an item is also known as its hash attribute",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableLocalSecondaryIndexKeySchema,
					},
					{
						Name:        "projection_non_key_attributes",
						Description: "Represents the non-key attribute names which will be projected into the index. For local secondary indexes, the total count of NonKeyAttributes summed across all of the local secondary indexes, must not exceed 20",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Projection.NonKeyAttributes"),
					},
					{
						Name:        "projection_type",
						Description: "The set of attributes that are projected into the index:  * KEYS_ONLY - Only the index and primary keys are projected into the index.  * INCLUDE - In addition to the attributes described in KEYS_ONLY, the secondary index will include other non-key attributes that you specify.  * ALL - All of the table attributes are projected into the index.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Projection.ProjectionType"),
					},
				},
			},
			{
				Name:        "aws_dynamodb_table_replicas",
				Description: "Contains the details of the replica.",
				Resolver:    fetchDynamodbTableReplicas,
				Columns: []schema.Column{
					{
						Name:        "table_cq_id",
						Description: "Unique CloudQuery ID of aws_dynamodb_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "global_secondary_indexes",
						Description: "Replica-specific global secondary index settings.",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableReplicaGlobalSecondaryIndexes,
					},
					{
						Name:        "kms_master_key_id",
						Description: "The KMS key of the replica that will be used for KMS encryption.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KMSMasterKeyId"),
					},
					{
						Name:        "provisioned_throughput_override_read_capacity_units",
						Description: "Replica-specific read capacity units",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ProvisionedThroughputOverride.ReadCapacityUnits"),
					},
					{
						Name:        "region_name",
						Description: "The name of the Region.",
						Type:        schema.TypeString,
					},
					{
						Name:        "replica_inaccessible_date_time",
						Description: "The time at which the replica was first detected as inaccessible",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "replica_status",
						Description: "The current state of the replica:  * CREATING - The replica is being created.  * UPDATING - The replica is being updated.  * DELETING - The replica is being deleted.  * ACTIVE - The replica is ready for use.  * REGION_DISABLED - The replica is inaccessible because the Amazon Web Services Region has been disabled",
						Type:        schema.TypeString,
					},
					{
						Name:        "replica_status_description",
						Description: "Detailed information about the replica status.",
						Type:        schema.TypeString,
					},
					{
						Name:        "replica_status_percent_progress",
						Description: "Specifies the progress of a Create, Update, or Delete action on the replica as a percentage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "replica_table_class_summary_last_update_date_time",
						Description: "The date and time at which the table class was last updated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ReplicaTableClassSummary.LastUpdateDateTime"),
					},
					{
						Name:        "replica_table_class_summary_table_class",
						Description: "The table class of the specified table",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReplicaTableClassSummary.TableClass"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDynamodbTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DynamoDB

	config := dynamodb.ListTablesInput{}
	for {
		output, err := svc.ListTables(ctx, &config, func(o *dynamodb.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}

		for i := range output.TableNames {
			response, err := svc.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: &output.TableNames[i]}, func(o *dynamodb.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- response.Table
		}

		if aws.ToString(output.LastEvaluatedTableName) == "" {
			break
		}
		config.ExclusiveStartTableName = output.LastEvaluatedTableName
	}

	return nil
}

func resolveDynamodbTableArchivalSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.ArchivalSummary == nil {
		return nil
	}
	return resource.Set(c.Name, map[string]interface{}{
		"date_time":  r.ArchivalSummary.ArchivalDateTime,
		"backup_arn": r.ArchivalSummary.ArchivalBackupArn,
		"reason":     r.ArchivalSummary.ArchivalReason,
	})
}

func resolveDynamodbTableAttributeDefinitions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, r.AttributeDefinitions)
}

func resolveDynamodbTableBillingModeSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, r.BillingModeSummary)
}

func resolveDynamodbTableKeySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, r.KeySchema)
}

func resolveDynamodbTableRestoreSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, r.RestoreSummary)
}

func resolveDynamodbTableStreamSpecification(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, r.StreamSpecification)
}

func fetchDynamodbTableGlobalSecondaryIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(*types.TableDescription)
	for i := range p.GlobalSecondaryIndexes {
		res <- p.GlobalSecondaryIndexes[i]
	}
	return nil
}

func resolveDynamodbTableGlobalSecondaryIndexKeySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.GlobalSecondaryIndexDescription)
	return resource.Set(c.Name, r.KeySchema)
}

func fetchDynamodbTableLocalSecondaryIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(*types.TableDescription)
	for i := range p.LocalSecondaryIndexes {
		res <- p.LocalSecondaryIndexes[i]
	}
	return nil
}

func resolveDynamodbTableLocalSecondaryIndexKeySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.LocalSecondaryIndexDescription)
	return resource.Set(c.Name, r.KeySchema)
}

func fetchDynamodbTableReplicas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(*types.TableDescription)
	for i := range p.Replicas {
		res <- p.Replicas[i]
	}
	return nil
}

func resolveDynamodbTableReplicaGlobalSecondaryIndexes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ReplicaDescription)
	return resource.Set(c.Name, r.GlobalSecondaryIndexes)
}
