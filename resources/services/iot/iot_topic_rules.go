package iot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotTopicRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_topic_rules",
		Description:  "The output from the GetTopicRule operation.",
		Resolver:     fetchIotTopicRules,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotTopicRuleTags,
			},
			{
				Name:        "aws_iot_sql_version",
				Description: "The version of the SQL rules engine to use when evaluating the rule.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.AwsIotSqlVersion"),
			},
			{
				Name:        "created_at",
				Description: "The date and time the rule was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Rule.CreatedAt"),
			},
			{
				Name:        "description",
				Description: "The description of the rule.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.Description"),
			},
			{
				Name:        "error_action_cloudwatch_alarm_name",
				Description: "The CloudWatch alarm name.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchAlarm.AlarmName"),
			},
			{
				Name:        "error_action_cloudwatch_alarm_role_arn",
				Description: "The IAM role that allows access to the CloudWatch alarm.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchAlarm.RoleArn"),
			},
			{
				Name:        "error_action_cloudwatch_alarm_state_reason",
				Description: "The reason for the alarm change.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchAlarm.StateReason"),
			},
			{
				Name:        "error_action_cloudwatch_alarm_state_value",
				Description: "The value of the alarm state",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchAlarm.StateValue"),
			},
			{
				Name:        "error_action_cloudwatch_logs_log_group_name",
				Description: "The CloudWatch log group to which the action sends data.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchLogs.LogGroupName"),
			},
			{
				Name:        "error_action_cloudwatch_logs_role_arn",
				Description: "The IAM role that allows access to the CloudWatch log.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchLogs.RoleArn"),
			},
			{
				Name:        "error_action_cloudwatch_metric_metric_name",
				Description: "The CloudWatch metric name.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.MetricName"),
			},
			{
				Name:        "error_action_cloudwatch_metric_metric_namespace",
				Description: "The CloudWatch metric namespace name.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.MetricNamespace"),
			},
			{
				Name:        "error_action_cloudwatch_metric_unit",
				Description: "The metric unit (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.MetricUnit"),
			},
			{
				Name:        "error_action_cloudwatch_metric_value",
				Description: "The CloudWatch metric value.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.MetricValue"),
			},
			{
				Name:        "error_action_cloudwatch_metric_role_arn",
				Description: "The IAM role that allows access to the CloudWatch metric.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.RoleArn"),
			},
			{
				Name:        "error_action_cloudwatch_metric_timestamp",
				Description: "An optional Unix timestamp (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.CloudwatchMetric.MetricTimestamp"),
			},
			{
				Name:        "error_action_dynamo_db_hash_key_field",
				Description: "The hash key name.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.HashKeyField"),
			},
			{
				Name:        "error_action_dynamo_db_hash_key_value",
				Description: "The hash key value.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.HashKeyValue"),
			},
			{
				Name:        "error_action_dynamo_db_role_arn",
				Description: "The ARN of the IAM role that grants access to the DynamoDB table.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.RoleArn"),
			},
			{
				Name:        "error_action_dynamo_db_table_name",
				Description: "The name of the DynamoDB table.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.TableName"),
			},
			{
				Name:        "error_action_dynamo_db_hash_key_type",
				Description: "The hash key type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.HashKeyType"),
			},
			{
				Name:        "error_action_dynamo_db_operation",
				Description: "The type of operation to be performed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.Operation"),
			},
			{
				Name:        "error_action_dynamo_db_payload_field",
				Description: "The action payload",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.PayloadField"),
			},
			{
				Name:        "error_action_dynamo_db_range_key_field",
				Description: "The range key name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.RangeKeyField"),
			},
			{
				Name:        "error_action_dynamo_db_range_key_type",
				Description: "The range key type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.RangeKeyType"),
			},
			{
				Name:        "error_action_dynamo_db_range_key_value",
				Description: "The range key value.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDB.RangeKeyValue"),
			},
			{
				Name:        "error_action_dynamo_db_v2_put_item_table_name",
				Description: "The table where the message data will be written.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDBv2.PutItem.TableName"),
			},
			{
				Name:        "error_action_dynamo_db_v2_role_arn",
				Description: "The ARN of the IAM role that grants access to the DynamoDB table.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.DynamoDBv2.RoleArn"),
			},
			{
				Name:        "error_action_elasticsearch_endpoint",
				Description: "The endpoint of your OpenSearch domain.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Elasticsearch.Endpoint"),
			},
			{
				Name:        "error_action_elasticsearch_id",
				Description: "The unique identifier for the document you are storing.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Elasticsearch.Id"),
			},
			{
				Name:        "error_action_elasticsearch_index",
				Description: "The index where you want to store your data.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Elasticsearch.Index"),
			},
			{
				Name:        "error_action_elasticsearch_role_arn",
				Description: "The IAM role ARN that has access to OpenSearch.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Elasticsearch.RoleArn"),
			},
			{
				Name:        "error_action_elasticsearch_type",
				Description: "The type of document you are storing.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Elasticsearch.Type"),
			},
			{
				Name:        "error_action_firehose_delivery_stream_name",
				Description: "The delivery stream name.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Firehose.DeliveryStreamName"),
			},
			{
				Name:        "error_action_firehose_role_arn",
				Description: "The IAM role that grants access to the Amazon Kinesis Firehose stream.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Firehose.RoleArn"),
			},
			{
				Name:        "error_action_firehose_batch_mode",
				Description: "Whether to deliver the Kinesis Data Firehose stream as a batch by using PutRecordBatch (https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html). The default value is false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Firehose.BatchMode"),
			},
			{
				Name:        "error_action_firehose_separator",
				Description: "A character separator that will be used to separate records written to the Firehose stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Firehose.Separator"),
			},
			{
				Name:        "error_action_http_url",
				Description: "The endpoint URL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Http.Url"),
			},
			{
				Name:        "error_action_http_auth_sigv4_role_arn",
				Description: "The ARN of the signing role.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Http.Auth.Sigv4.RoleArn"),
			},
			{
				Name:        "error_action_http_auth_sigv4_service_name",
				Description: "The service name to use while signing with Sig V4.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Http.Auth.Sigv4.ServiceName"),
			},
			{
				Name:        "error_action_http_auth_sigv4_signing_region",
				Description: "The signing region.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Http.Auth.Sigv4.SigningRegion"),
			},
			{
				Name:        "error_action_http_confirmation_url",
				Description: "The URL to which IoT sends a confirmation message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Http.ConfirmationUrl"),
			},
			{
				Name:        "error_action_http_headers",
				Description: "The HTTP headers to send with the message data.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIotTopicRulesErrorActionHttpHeaders,
			},
			{
				Name:        "error_action_iot_analytics_batch_mode",
				Description: "Whether to process the action as a batch",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotAnalytics.BatchMode"),
			},
			{
				Name:        "error_action_iot_analytics_channel_arn",
				Description: "(deprecated) The ARN of the IoT Analytics channel to which message data will be sent.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotAnalytics.ChannelArn"),
			},
			{
				Name:        "error_action_iot_analytics_channel_name",
				Description: "The name of the IoT Analytics channel to which message data will be sent.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotAnalytics.ChannelName"),
			},
			{
				Name:        "error_action_iot_analytics_role_arn",
				Description: "The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotAnalytics.RoleArn"),
			},
			{
				Name:        "error_action_iot_events_input_name",
				Description: "The name of the IoT Events input.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotEvents.InputName"),
			},
			{
				Name:        "error_action_iot_events_role_arn",
				Description: "The ARN of the role that grants IoT permission to send an input to an IoT Events detector",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotEvents.RoleArn"),
			},
			{
				Name:        "error_action_iot_events_batch_mode",
				Description: "Whether to process the event actions as a batch",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotEvents.BatchMode"),
			},
			{
				Name:        "error_action_iot_events_message_id",
				Description: "The ID of the message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.IotEvents.MessageId"),
			},
			{
				Name:        "error_action_iot_site_wise",
				Description: "Sends data from the MQTT message that triggered the rule to IoT SiteWise asset properties.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIotTopicRulesErrorActionIotSiteWise,
			},
			{
				Name:        "error_action_kafka_client_properties",
				Description: "Properties of the Apache Kafka producer client.  This member is required.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kafka.ClientProperties"),
			},
			{
				Name:        "error_action_kafka_destination_arn",
				Description: "The ARN of Kafka action's VPC TopicRuleDestination.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kafka.DestinationArn"),
			},
			{
				Name:        "error_action_kafka_topic",
				Description: "The Kafka topic for messages to be sent to the Kafka broker.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kafka.Topic"),
			},
			{
				Name:        "error_action_kafka_key",
				Description: "The Kafka message key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kafka.Key"),
			},
			{
				Name:        "error_action_kafka_partition",
				Description: "The Kafka message partition.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kafka.Partition"),
			},
			{
				Name:        "error_action_kinesis_role_arn",
				Description: "The ARN of the IAM role that grants access to the Amazon Kinesis stream.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kinesis.RoleArn"),
			},
			{
				Name:        "error_action_kinesis_stream_name",
				Description: "The name of the Amazon Kinesis stream.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kinesis.StreamName"),
			},
			{
				Name:        "error_action_kinesis_partition_key",
				Description: "The partition key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Kinesis.PartitionKey"),
			},
			{
				Name:        "error_action_lambda_function_arn",
				Description: "The ARN of the Lambda function.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Lambda.FunctionArn"),
			},
			{
				Name:        "error_action_open_search_endpoint",
				Description: "The endpoint of your OpenSearch domain.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.OpenSearch.Endpoint"),
			},
			{
				Name:        "error_action_open_search_id",
				Description: "The unique identifier for the document you are storing.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.OpenSearch.Id"),
			},
			{
				Name:        "error_action_open_search_index",
				Description: "The OpenSearch index where you want to store your data.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.OpenSearch.Index"),
			},
			{
				Name:        "error_action_open_search_role_arn",
				Description: "The IAM role ARN that has access to OpenSearch.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.OpenSearch.RoleArn"),
			},
			{
				Name:        "error_action_open_search_type",
				Description: "The type of document you are storing.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.OpenSearch.Type"),
			},
			{
				Name:        "error_action_republish_role_arn",
				Description: "The ARN of the IAM role that grants access.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Republish.RoleArn"),
			},
			{
				Name:        "error_action_republish_topic",
				Description: "The name of the MQTT topic.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Republish.Topic"),
			},
			{
				Name:        "error_action_republish_qos",
				Description: "The Quality of Service (QoS) level to use when republishing messages",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Republish.Qos"),
			},
			{
				Name:        "error_action_s3_bucket_name",
				Description: "The Amazon S3 bucket.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.S3.BucketName"),
			},
			{
				Name:        "error_action_s3_key",
				Description: "The object key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.S3.Key"),
			},
			{
				Name:        "error_action_s3_role_arn",
				Description: "The ARN of the IAM role that grants access.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.S3.RoleArn"),
			},
			{
				Name:        "error_action_s3_canned_acl",
				Description: "The Amazon S3 canned ACL that controls access to the object identified by the object key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.S3.CannedAcl"),
			},
			{
				Name:        "error_action_salesforce_token",
				Description: "The token used to authenticate access to the Salesforce IoT Cloud Input Stream. The token is available from the Salesforce IoT Cloud platform after creation of the Input Stream.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Salesforce.Token"),
			},
			{
				Name:        "error_action_salesforce_url",
				Description: "The URL exposed by the Salesforce IoT Cloud Input Stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Salesforce.Url"),
			},
			{
				Name:        "error_action_sns_role_arn",
				Description: "The ARN of the IAM role that grants access.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sns.RoleArn"),
			},
			{
				Name:        "error_action_sns_target_arn",
				Description: "The ARN of the SNS topic.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sns.TargetArn"),
			},
			{
				Name:        "error_action_sns_message_format",
				Description: "(Optional) The message format of the message to publish",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sns.MessageFormat"),
			},
			{
				Name:        "error_action_sqs_queue_url",
				Description: "The URL of the Amazon SQS queue.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sqs.QueueUrl"),
			},
			{
				Name:        "error_action_sqs_role_arn",
				Description: "The ARN of the IAM role that grants access.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sqs.RoleArn"),
			},
			{
				Name:        "error_action_sqs_use_base64",
				Description: "Specifies whether to use Base64 encoding.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Sqs.UseBase64"),
			},
			{
				Name:        "error_action_step_functions_role_arn",
				Description: "The ARN of the role that grants IoT permission to start execution of a state machine (\"Action\":\"states:StartExecution\").  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.StepFunctions.RoleArn"),
			},
			{
				Name:        "error_action_step_functions_state_machine_name",
				Description: "The name of the Step Functions state machine whose execution will be started.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.StepFunctions.StateMachineName"),
			},
			{
				Name:        "error_action_step_functions_execution_name_prefix",
				Description: "(Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.StepFunctions.ExecutionNamePrefix"),
			},
			{
				Name:        "error_action_timestream_database_name",
				Description: "The name of an Amazon Timestream database.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Timestream.DatabaseName"),
			},
			{
				Name:        "error_action_timestream_dimensions",
				Description: "Metadata attributes of the time series that are written in each measure record.  This member is required.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIotTopicRulesErrorActionTimestreamDimensions,
			},
			{
				Name:        "error_action_timestream_role_arn",
				Description: "The ARN of the role that grants permission to write to the Amazon Timestream database table.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Timestream.RoleArn"),
			},
			{
				Name:        "error_action_timestream_table_name",
				Description: "The name of the database table into which to write the measure records.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Timestream.TableName"),
			},
			{
				Name:        "error_action_timestream_timestamp_unit",
				Description: "The precision of the timestamp value that results from the expression described in value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Timestream.Timestamp.Unit"),
			},
			{
				Name:        "error_action_timestream_timestamp_value",
				Description: "An expression that returns a long epoch time value.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.ErrorAction.Timestream.Timestamp.Value"),
			},
			{
				Name:        "rule_disabled",
				Description: "Specifies whether the rule is disabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Rule.RuleDisabled"),
			},
			{
				Name:        "rule_name",
				Description: "The name of the rule.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.RuleName"),
			},
			{
				Name:        "sql",
				Description: "The SQL statement used to query the topic",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rule.Sql"),
			},
			{
				Name:        "arn",
				Description: "The rule ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleArn"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iot_topic_rule_actions",
				Description: "Describes the actions associated with a rule.",
				Resolver:    fetchIotTopicRuleActions,
				Columns: []schema.Column{
					{
						Name:        "topic_rule_cq_id",
						Description: "Unique CloudQuery ID of aws_iot_topic_rules table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cloudwatch_alarm_alarm_name",
						Description: "The CloudWatch alarm name.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchAlarm.AlarmName"),
					},
					{
						Name:        "cloudwatch_alarm_role_arn",
						Description: "The IAM role that allows access to the CloudWatch alarm.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchAlarm.RoleArn"),
					},
					{
						Name:        "cloudwatch_alarm_state_reason",
						Description: "The reason for the alarm change.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchAlarm.StateReason"),
					},
					{
						Name:        "cloudwatch_alarm_state_value",
						Description: "The value of the alarm state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchAlarm.StateValue"),
					},
					{
						Name:        "cloudwatch_logs_log_group_name",
						Description: "The CloudWatch log group to which the action sends data.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchLogs.LogGroupName"),
					},
					{
						Name:        "cloudwatch_logs_role_arn",
						Description: "The IAM role that allows access to the CloudWatch log.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchLogs.RoleArn"),
					},
					{
						Name:        "cloudwatch_metric_metric_name",
						Description: "The CloudWatch metric name.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.MetricName"),
					},
					{
						Name:        "cloudwatch_metric_metric_namespace",
						Description: "The CloudWatch metric namespace name.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.MetricNamespace"),
					},
					{
						Name:        "cloudwatch_metric_metric_unit",
						Description: "The metric unit (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.MetricUnit"),
					},
					{
						Name:        "cloudwatch_metric_metric_value",
						Description: "The CloudWatch metric value.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.MetricValue"),
					},
					{
						Name:        "cloudwatch_metric_role_arn",
						Description: "The IAM role that allows access to the CloudWatch metric.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.RoleArn"),
					},
					{
						Name:        "cloudwatch_metric_metric_timestamp",
						Description: "An optional Unix timestamp (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudwatchMetric.MetricTimestamp"),
					},
					{
						Name:        "dynamo_db_hash_key_field",
						Description: "The hash key name.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.HashKeyField"),
					},
					{
						Name:        "dynamo_db_hash_key_value",
						Description: "The hash key value.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.HashKeyValue"),
					},
					{
						Name:        "dynamo_db_role_arn",
						Description: "The ARN of the IAM role that grants access to the DynamoDB table.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.RoleArn"),
					},
					{
						Name:        "dynamo_db_table_name",
						Description: "The name of the DynamoDB table.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.TableName"),
					},
					{
						Name:        "dynamo_db_hash_key_type",
						Description: "The hash key type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.HashKeyType"),
					},
					{
						Name:        "dynamo_db_operation",
						Description: "The type of operation to be performed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.Operation"),
					},
					{
						Name:        "dynamo_db_payload_field",
						Description: "The action payload",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.PayloadField"),
					},
					{
						Name:        "dynamo_db_range_key_field",
						Description: "The range key name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.RangeKeyField"),
					},
					{
						Name:        "dynamo_db_range_key_type",
						Description: "The range key type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.RangeKeyType"),
					},
					{
						Name:        "dynamo_db_range_key_value",
						Description: "The range key value.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDB.RangeKeyValue"),
					},
					{
						Name:        "dynamo_db_v2_put_item_table_name",
						Description: "The table where the message data will be written.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDBv2.PutItem.TableName"),
					},
					{
						Name:        "dynamo_db_v2_role_arn",
						Description: "The ARN of the IAM role that grants access to the DynamoDB table.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DynamoDBv2.RoleArn"),
					},
					{
						Name:        "elasticsearch_endpoint",
						Description: "The endpoint of your OpenSearch domain.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Elasticsearch.Endpoint"),
					},
					{
						Name:        "elasticsearch_id",
						Description: "The unique identifier for the document you are storing.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Elasticsearch.Id"),
					},
					{
						Name:        "elasticsearch_index",
						Description: "The index where you want to store your data.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Elasticsearch.Index"),
					},
					{
						Name:        "elasticsearch_role_arn",
						Description: "The IAM role ARN that has access to OpenSearch.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Elasticsearch.RoleArn"),
					},
					{
						Name:        "elasticsearch_type",
						Description: "The type of document you are storing.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Elasticsearch.Type"),
					},
					{
						Name:        "firehose_delivery_stream_name",
						Description: "The delivery stream name.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Firehose.DeliveryStreamName"),
					},
					{
						Name:        "firehose_role_arn",
						Description: "The IAM role that grants access to the Amazon Kinesis Firehose stream.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Firehose.RoleArn"),
					},
					{
						Name:        "firehose_batch_mode",
						Description: "Whether to deliver the Kinesis Data Firehose stream as a batch by using PutRecordBatch (https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html). The default value is false",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Firehose.BatchMode"),
					},
					{
						Name:        "firehose_separator",
						Description: "A character separator that will be used to separate records written to the Firehose stream",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Firehose.Separator"),
					},
					{
						Name:        "http_url",
						Description: "The endpoint URL",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Http.Url"),
					},
					{
						Name:        "http_auth_sigv4_role_arn",
						Description: "The ARN of the signing role.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Http.Auth.Sigv4.RoleArn"),
					},
					{
						Name:        "http_auth_sigv4_service_name",
						Description: "The service name to use while signing with Sig V4.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Http.Auth.Sigv4.ServiceName"),
					},
					{
						Name:        "http_auth_sigv4_signing_region",
						Description: "The signing region.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Http.Auth.Sigv4.SigningRegion"),
					},
					{
						Name:        "http_confirmation_url",
						Description: "The URL to which IoT sends a confirmation message",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Http.ConfirmationUrl"),
					},
					{
						Name:        "http_headers",
						Description: "The HTTP headers to send with the message data.",
						Type:        schema.TypeJSON,
						Resolver:    resolveIotTopicRuleActionsHttpHeaders,
					},
					{
						Name:        "iot_analytics_batch_mode",
						Description: "Whether to process the action as a batch",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("IotAnalytics.BatchMode"),
					},
					{
						Name:        "iot_analytics_channel_arn",
						Description: "(deprecated) The ARN of the IoT Analytics channel to which message data will be sent.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotAnalytics.ChannelArn"),
					},
					{
						Name:        "iot_analytics_channel_name",
						Description: "The name of the IoT Analytics channel to which message data will be sent.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotAnalytics.ChannelName"),
					},
					{
						Name:        "iot_analytics_role_arn",
						Description: "The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotAnalytics.RoleArn"),
					},
					{
						Name:        "iot_events_input_name",
						Description: "The name of the IoT Events input.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotEvents.InputName"),
					},
					{
						Name:        "iot_events_role_arn",
						Description: "The ARN of the role that grants IoT permission to send an input to an IoT Events detector",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotEvents.RoleArn"),
					},
					{
						Name:        "iot_events_batch_mode",
						Description: "Whether to process the event actions as a batch",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("IotEvents.BatchMode"),
					},
					{
						Name:        "iot_events_message_id",
						Description: "The ID of the message",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IotEvents.MessageId"),
					},
					{
						Name:        "iot_site_wise",
						Description: "Sends data from the MQTT message that triggered the rule to IoT SiteWise asset properties.",
						Type:        schema.TypeJSON,
						Resolver:    resolveIotTopicRuleActionsIotSiteWise,
					},
					{
						Name:        "kafka_client_properties",
						Description: "Properties of the Apache Kafka producer client.  This member is required.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Kafka.ClientProperties"),
					},
					{
						Name:        "kafka_destination_arn",
						Description: "The ARN of Kafka action's VPC TopicRuleDestination.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kafka.DestinationArn"),
					},
					{
						Name:        "kafka_topic",
						Description: "The Kafka topic for messages to be sent to the Kafka broker.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kafka.Topic"),
					},
					{
						Name:        "kafka_key",
						Description: "The Kafka message key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kafka.Key"),
					},
					{
						Name:        "kafka_partition",
						Description: "The Kafka message partition.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kafka.Partition"),
					},
					{
						Name:        "kinesis_role_arn",
						Description: "The ARN of the IAM role that grants access to the Amazon Kinesis stream.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kinesis.RoleArn"),
					},
					{
						Name:        "kinesis_stream_name",
						Description: "The name of the Amazon Kinesis stream.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kinesis.StreamName"),
					},
					{
						Name:        "kinesis_partition_key",
						Description: "The partition key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Kinesis.PartitionKey"),
					},
					{
						Name:        "lambda_function_arn",
						Description: "The ARN of the Lambda function.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Lambda.FunctionArn"),
					},
					{
						Name:        "open_search_endpoint",
						Description: "The endpoint of your OpenSearch domain.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenSearch.Endpoint"),
					},
					{
						Name:        "open_search_id",
						Description: "The unique identifier for the document you are storing.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenSearch.Id"),
					},
					{
						Name:        "open_search_index",
						Description: "The OpenSearch index where you want to store your data.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenSearch.Index"),
					},
					{
						Name:        "open_search_role_arn",
						Description: "The IAM role ARN that has access to OpenSearch.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenSearch.RoleArn"),
					},
					{
						Name:        "open_search_type",
						Description: "The type of document you are storing.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenSearch.Type"),
					},
					{
						Name:        "republish_role_arn",
						Description: "The ARN of the IAM role that grants access.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Republish.RoleArn"),
					},
					{
						Name:        "republish_topic",
						Description: "The name of the MQTT topic.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Republish.Topic"),
					},
					{
						Name:        "republish_qos",
						Description: "The Quality of Service (QoS) level to use when republishing messages",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Republish.Qos"),
					},
					{
						Name:        "s3_bucket_name",
						Description: "The Amazon S3 bucket.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3.BucketName"),
					},
					{
						Name:        "s3_key",
						Description: "The object key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3.Key"),
					},
					{
						Name:        "s3_role_arn",
						Description: "The ARN of the IAM role that grants access.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3.RoleArn"),
					},
					{
						Name:        "s3_canned_acl",
						Description: "The Amazon S3 canned ACL that controls access to the object identified by the object key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3.CannedAcl"),
					},
					{
						Name:        "salesforce_token",
						Description: "The token used to authenticate access to the Salesforce IoT Cloud Input Stream. The token is available from the Salesforce IoT Cloud platform after creation of the Input Stream.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Salesforce.Token"),
					},
					{
						Name:        "salesforce_url",
						Description: "The URL exposed by the Salesforce IoT Cloud Input Stream",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Salesforce.Url"),
					},
					{
						Name:        "sns_role_arn",
						Description: "The ARN of the IAM role that grants access.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Sns.RoleArn"),
					},
					{
						Name:        "sns_target_arn",
						Description: "The ARN of the SNS topic.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Sns.TargetArn"),
					},
					{
						Name:        "sns_message_format",
						Description: "(Optional) The message format of the message to publish",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Sns.MessageFormat"),
					},
					{
						Name:        "sqs_queue_url",
						Description: "The URL of the Amazon SQS queue.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Sqs.QueueUrl"),
					},
					{
						Name:        "sqs_role_arn",
						Description: "The ARN of the IAM role that grants access.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Sqs.RoleArn"),
					},
					{
						Name:        "sqs_use_base64",
						Description: "Specifies whether to use Base64 encoding.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Sqs.UseBase64"),
					},
					{
						Name:        "step_functions_role_arn",
						Description: "The ARN of the role that grants IoT permission to start execution of a state machine (\"Action\":\"states:StartExecution\").  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StepFunctions.RoleArn"),
					},
					{
						Name:        "step_functions_state_machine_name",
						Description: "The name of the Step Functions state machine whose execution will be started.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StepFunctions.StateMachineName"),
					},
					{
						Name:        "step_functions_execution_name_prefix",
						Description: "(Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StepFunctions.ExecutionNamePrefix"),
					},
					{
						Name:        "timestream_database_name",
						Description: "The name of an Amazon Timestream database.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Timestream.DatabaseName"),
					},
					{
						Name:        "timestream_dimensions",
						Description: "Metadata attributes of the time series that are written in each measure record.  This member is required.",
						Type:        schema.TypeJSON,
						Resolver:    resolveIotTopicRuleActionsTimestreamDimensions,
					},
					{
						Name:        "timestream_role_arn",
						Description: "The ARN of the role that grants permission to write to the Amazon Timestream database table.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Timestream.RoleArn"),
					},
					{
						Name:        "timestream_table_name",
						Description: "The name of the database table into which to write the measure records.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Timestream.TableName"),
					},
					{
						Name:        "timestream_timestamp_unit",
						Description: "The precision of the timestamp value that results from the expression described in value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Timestream.Timestamp.Unit"),
					},
					{
						Name:        "timestream_timestamp_value",
						Description: "An expression that returns a long epoch time value.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Timestream.Timestamp.Value"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotTopicRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTopicRulesInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListTopicRules(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}

		for _, s := range response.Rules {
			rule, err := svc.GetTopicRule(ctx, &iot.GetTopicRuleInput{
				RuleName: s.RuleName,
			}, func(options *iot.Options) {
				options.Region = client.Region
			})
			if err != nil {
				return err
			}
			res <- rule
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotTopicRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.GetTopicRuleOutput)
	if !ok {
		return fmt.Errorf("expected  *iot.GetTopicRuleOutput but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.RuleArn,
	}
	tags := make(map[string]interface{})

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})

		if err != nil {
			return err
		}
		for _, t := range response.Tags {
			tags[*t.Key] = t.Value
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
func resolveIotTopicRulesErrorActionHttpHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.GetTopicRuleOutput)
	if !ok {
		return fmt.Errorf("expected  *iot.GetTopicRuleOutput but got %T", resource.Item)
	}
	if i.Rule == nil || i.Rule.ErrorAction == nil || i.Rule.ErrorAction.Http == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Rule.ErrorAction.Http.Headers {
		j[*h.Key] = *h.Value
	}
	return resource.Set(c.Name, j)
}
func resolveIotTopicRulesErrorActionIotSiteWise(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.GetTopicRuleOutput)
	if !ok {
		return fmt.Errorf("expected  *types.Action but got %T", resource.Item)
	}
	if i.Rule == nil || i.Rule.ErrorAction == nil || i.Rule.ErrorAction.IotSiteWise == nil {
		return nil
	}
	b, err := json.Marshal(i.Rule.ErrorAction.IotSiteWise)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveIotTopicRulesErrorActionTimestreamDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*iot.GetTopicRuleOutput)
	if !ok {
		return fmt.Errorf("expected  *iot.GetTopicRuleOutput but got %T", resource.Item)
	}
	if i.Rule == nil || i.Rule.ErrorAction == nil || i.Rule.ErrorAction.Timestream == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Rule.ErrorAction.Timestream.Dimensions {
		j[*h.Name] = *h.Value
	}
	return resource.Set(c.Name, j)
}
func fetchIotTopicRuleActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i, ok := parent.Item.(*iot.GetTopicRuleOutput)
	if !ok {
		return fmt.Errorf("expected  *iot.GetTopicRuleOutput but got %T", parent.Item)
	}
	if i.Rule == nil {
		return nil
	}
	res <- i.Rule.Actions
	return nil
}
func resolveIotTopicRuleActionsHttpHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(types.Action)
	if !ok {
		return fmt.Errorf("expected  *types.Action but got %T", resource.Item)
	}
	if i.Http == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Http.Headers {
		j[*h.Key] = *h.Value
	}
	return resource.Set(c.Name, j)
}
func resolveIotTopicRuleActionsIotSiteWise(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(types.Action)
	if !ok {
		return fmt.Errorf("expected  *types.Action but got %T", resource.Item)
	}
	if i.IotSiteWise == nil {
		return nil
	}
	b, err := json.Marshal(i.IotSiteWise)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveIotTopicRuleActionsTimestreamDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(types.Action)
	if !ok {
		return fmt.Errorf("expected  *types.Action but got %T", resource.Item)
	}
	if i.Timestream == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Timestream.Dimensions {
		j[*h.Name] = *h.Value
	}
	return resource.Set(c.Name, j)
}
