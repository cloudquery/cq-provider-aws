package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func S3Accounts() *schema.Table {
	return &schema.Table{
		Name:                 "aws_s3_accounts",
		Description:          "An Amazon S3 bucket is a public cloud storage resource available in Amazon Web Services' (AWS) Simple Storage Service (S3)",
		Resolver:             fetchS3Account,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: resolveS3BucketsAttributes,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "name"}},
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
			},
			{
				Name: "logging_target_prefix",
				Type: schema.TypeString,
			},
			{
				Name: "logging_target_bucket",
				Type: schema.TypeString,
			},
			{
				Name: "versioning_status",
				Type: schema.TypeString,
			},
			{
				Name: "versioning_mfa_delete",
				Type: schema.TypeString,
			},
			{
				Name: "policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:        "creation_date",
				Description: "Date the bucket was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "block_public_acls",
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:        "block_public_policy",
				Description: "Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ignore_public_acls",
				Description: "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:        "restrict_public_buckets",
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:        "replication_role",
				Description: "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Role"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the s3 bucket",
				Type:        schema.TypeString,
				Resolver:    resolveS3BucketsArn,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_s3_bucket_grants",
				Description: "Container for grant information.",
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Resolver:    fetchS3BucketGrants,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of grantee",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.Type"),
					},
					{
						Name:        "display_name",
						Description: "Screen name of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.DisplayName"),
					},
					{
						Name:        "email_address",
						Description: "Email address of the grantee",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.EmailAddress"),
					},
					{
						Name:        "grantee_id",
						Description: "The canonical user ID of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.ID"),
					},
					{
						Name:        "uri",
						Description: "URI of the grantee group.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.URI"),
					},
					{
						Name:        "permission",
						Description: "Specifies the permission given to the grantee.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_s3_bucket_cors_rules",
				Description: "Specifies a cross-origin access rule for an Amazon S3 bucket.",
				Resolver:    fetchS3BucketCorsRules,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_methods",
						Description: "An HTTP method that you allow the origin to execute",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_origins",
						Description: "One or more origins you want customers to be able to access the bucket from.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_headers",
						Description: "Headers that are specified in the Access-Control-Request-Headers header",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "expose_headers",
						Description: "One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "id",
						Description: "Unique identifier for the rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "max_age_seconds",
						Description: "The time in seconds that your browser is to cache the preflight response for the specified resource.",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:        "aws_s3_bucket_encryption_rules",
				Description: "Specifies the default server-side encryption configuration.",
				Resolver:    fetchS3BucketEncryptionRules,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "sse_algorithm",
						Description: "Server-side encryption algorithm to use for the default encryption.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplyServerSideEncryptionByDefault.SSEAlgorithm"),
					},
					{
						Name:        "kms_master_key_id",
						Description: "AWS Key Management Service (KMS) customer master key ID to use for the default encryption",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplyServerSideEncryptionByDefault.KMSMasterKeyID"),
					},
					{
						Name:        "bucket_key_enabled",
						Description: "Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_s3_bucket_replication_rules",
				Description: "Specifies which Amazon S3 objects to replicate and where to store the replicas.",
				Resolver:    fetchS3BucketReplicationRules,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"bucket_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "destination_bucket",
						Description: "The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.Bucket"),
					},
					{
						Name:        "destination_access_control_translation_owner",
						Description: "Specifies the replica ownership",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.AccessControlTranslation.Owner"),
					},
					{
						Name:        "destination_account",
						Description: "Destination bucket owner account ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.Account"),
					},
					{
						Name:        "destination_encryption_configuration_replica_kms_key_id",
						Description: "Specifies the ID (Key ARN or Alias ARN) of the customer managed customer master key (CMK) stored in AWS Key Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.EncryptionConfiguration.ReplicaKmsKeyID"),
					},
					{
						Name:        "destination_metrics_status",
						Description: "Specifies whether the replication metrics are enabled.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.Metrics.Status"),
					},
					{
						Name:        "destination_metrics_event_threshold_minutes",
						Description: "Contains an integer specifying time in minutes",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Destination.Metrics.EventThreshold.Minutes"),
					},
					{
						Name:        "destination_replication_time_status",
						Description: "Specifies whether the replication time is enabled.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.ReplicationTime.Status"),
					},
					{
						Name:        "destination_replication_time_minutes",
						Description: "Contains an integer specifying time in minutes",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Destination.ReplicationTime.Time.Minutes"),
					},
					{
						Name:        "destination_storage_class",
						Description: "The storage class to use when replicating objects, such as S3 Standard or reduced redundancy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Destination.StorageClass"),
					},
					{
						Name:        "status",
						Description: "Specifies whether the rule is enabled.",
						Type:        schema.TypeString,
					},
					{
						Name:        "delete_marker_replication_status",
						Description: "Indicates whether to replicate delete markers",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DeleteMarkerReplication.Status"),
					},
					{
						Name:     "existing_object_replication_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ExistingObjectReplication.Status"),
					},
					{
						Name:        "filter",
						Description: "A filter that identifies the subset of objects to which the replication rule applies",
						Type:        schema.TypeJSON,
						Resolver:    resolveS3BucketReplicationRuleFilter,
					},
					{
						Name:        "id",
						Description: "A unique identifier for the rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "prefix",
						Description: "An object key name prefix that identifies the object or objects to which the rule applies",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "The priority indicates which rule has precedence whenever two or more replication rules conflict",
						Type:        schema.TypeInt,
					},
					{
						Name:        "source_replica_modifications_status",
						Description: "Specifies whether Amazon S3 replicates modifications on replicas.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceSelectionCriteria.ReplicaModifications.Status"),
					},
					{
						Name:        "source_sse_kms_encrypted_objects_status",
						Description: "Specifies whether Amazon S3 replicates objects created with server-side encryption using a customer master key (CMK) stored in AWS Key Management Service.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceSelectionCriteria.SseKmsEncryptedObjects.Status"),
					},
				},
			},
			{
				Name:        "aws_s3_bucket_lifecycles",
				Description: "A lifecycle rule for individual objects in an Amazon S3 bucket.",
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Resolver:    fetchS3BucketLifecycles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"bucket_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "status",
						Description: "If 'Enabled', the rule is currently being applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "abort_incomplete_multipart_upload_days_after_initiation",
						Description: "Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("AbortIncompleteMultipartUpload.DaysAfterInitiation"),
					},
					{
						Name:        "expiration_date",
						Description: "Indicates at what date the object is to be moved or deleted",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Expiration.Date"),
					},
					{
						Name:        "expiration_days",
						Description: "Indicates the lifetime, in days, of the objects that are subject to the rule. The value must be a non-zero positive integer.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Expiration.Days"),
					},
					{
						Name:        "expiration_expired_object_delete_marker",
						Description: "Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Expiration.ExpiredObjectDeleteMarker"),
					},
					{
						Name:        "filter",
						Description: "The Filter is used to identify objects that a Lifecycle Rule applies to",
						Type:        schema.TypeJSON,
						Resolver:    resolveS3BucketLifecycleFilter,
					},
					{
						Name:        "id",
						Description: "Unique identifier for the rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "noncurrent_version_expiration_days",
						Description: "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("NoncurrentVersionExpiration.NoncurrentDays"),
					},
					{
						Name:        "noncurrent_version_transitions",
						Description: "Specifies the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class",
						Type:        schema.TypeJSON,
						Resolver:    resolveS3BucketLifecycleNoncurrentVersionTransitions,
					},
					{
						Name:        "prefix",
						Description: "Prefix identifying one or more objects to which the rule applies",
						Type:        schema.TypeString,
					},
					{
						Name:        "transitions",
						Description: "Specifies when an Amazon S3 object transitions to a specified storage class.",
						Type:        schema.TypeJSON,
						Resolver:    resolveS3BucketLifecycleTransitions,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchS3Account(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().S3
	response, err := svc.ListBuckets(ctx, nil)
	if err != nil {
		return err
	}
	wb := make([]*WrappedBucket, len(response.Buckets))
	for i, b := range response.Buckets {
		wb[i] = &WrappedBucket{b, nil, nil}
	}

	res <- wb
	return nil
}
