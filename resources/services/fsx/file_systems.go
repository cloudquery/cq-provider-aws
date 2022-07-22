package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func FsxFileSystems() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_file_systems",
		Description:  "An Amazon FSx file system.",
		Resolver:     fetchFsxFileSystems,
		Multiplex:    client.ServiceAccountRegionMultiplexer("fsx"),
		IgnoreError:  client.IgnoreCommonErrors,
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
				Name:        "id",
				Description: "The system-generated, unique 17-digit ID of the file system.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemId"),
			},
			{
				Name:        "creation_time",
				Description: "The time that the file system was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account that created the file system.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemId"),
			},
			{
				Name:        "file_system_type",
				Description: "The type of Amazon FSx file system.",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the file system.",
				Type:        schema.TypeString,
			},
			{
				Name:          "failure_details_message",
				Description:   "A message describing the data repository association failure.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("FailureDetails.Message"),
				IgnoreInTests: true,
			},
			{
				Name:        "storage_capacity",
				Description: "The storage capacity of the file system in gibibytes (GiB).",
				Type:        schema.TypeInt,
			},
			{
				Name:        "storage_type",
				Description: "The type of storage the file system is using.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the primary virtual private cloud (VPC) for the file system.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "Specifies the IDs of the subnets that the file system is accessible from.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "network_interface_ids",
				Description: "The IDs of the elastic network interfaces from which a specific file system is accessible.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "dns_name",
				Description: "The Domain Name System (DNS) name for the file system.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the Key Management Service (KMS) key used to encrypt Amazon FSx file system data.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the file system resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags associated with a particular file system.",
				Type:        schema.TypeJSON,
				Resolver:    resolveFsxFileSystemTags,
			},
			{
				Name:        "file_system_type_version",
				Description: "The Lustre version of the Amazon FSx for Lustre file system, either 2.10 or 2.12 .",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_fsx_file_system_windows_configs",
				Description:   "The configuration for this Amazon FSx for Windows File Server file system.",
				Resolver:      fetchFsxFileSystemWindowsConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "active_directory_id",
						Description: "The ID for an existing Amazon Web Services Managed Microsoft Active Directory instance that the file system is joined to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the file system deployment type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "remote_administration_endpoint",
						Description: "Specifies the file system remote administration endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "preferred_subnet_id",
						Description: "Specifies the file system preferred subnet ID.",
						Type:        schema.TypeString,
					},
					{
						Name:        "preferred_file_server_ip",
						Description: "Specifies the file system preferred file server IP.",
						Type:        schema.TypeString,
					},
					{
						Name:        "throughput_capacity",
						Description: "The throughput of the Amazon FSx file system, measured in megabytes per second.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "maintenance_operations_in_progress",
						Description: "The list of maintenance operations in progress for this file system.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone. d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.",
						Type:        schema.TypeString,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "The preferred time to take daily automatic backups, in the UTC time zone.",
						Type:        schema.TypeString,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A boolean flag indicating whether tags on the file system should be copied to backups.",
						Type:        schema.TypeBool,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_fsx_file_system_self_managed_active_directory_configs",
						Description:   "The configuration of the self-managed Microsoft Active Directory (AD) directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.",
						Resolver:      fetchFsxFileSystemSelfManagedActiveDirectoryConfigs,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "windows_config_cq_id",
								Description: "Unique CloudQuery ID of aws_fsx_file_system_windows_configs table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "active_directory_id",
								Description: "The Active Directory ID.",
								Type:        schema.TypeString,
								Resolver:    schema.ParentResourceFieldResolver("active_directory_id"),
							},
							{
								Name:        "domain_name",
								Description: "The fully qualified domain name of the self-managed AD directory.",
								Type:        schema.TypeString,
							},
							{
								Name:        "organizational_unit_distinguished_id",
								Description: "The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.",
								Type:        schema.TypeString,
							},
							{
								Name:        "file_system_administrators_group",
								Description: "The name of the domain group whose members have administrative privileges for the FSx file system.",
								Type:        schema.TypeString,
							},
							{
								Name:        "user_name",
								Description: "The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain.",
								Type:        schema.TypeString,
							},
							{
								Name:        "dns_ips",
								Description: "A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.",
								Type:        schema.TypeStringArray,
							},
						},
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_aliases",
				Description:   "DNS aliases that are currently associated with the Amazon FSx file system.",
				Resolver:      fetchFsxFileSystemAliases,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the DNS alias.",
						Type:        schema.TypeString,
					},
					{
						Name:        "lifecycle",
						Description: "The lifecycle status of DNS alias.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_audit_log_configs",
				Description:   "The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.",
				Resolver:      fetchFsxFileSystemAuditLogConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "file_access_audit_log_level",
						Description: "Sets which attempt type is logged by Amazon FSx for file and folder accesses.",
						Type:        schema.TypeString,
					},
					{
						Name:        "file_share_access_audit_log_level",
						Description: "Sets which attempt type is logged by Amazon FSx for file share accesses.",
						Type:        schema.TypeString,
					},
					{
						Name:        "audit_log_destination",
						Description: "The Amazon Resource Name (ARN) for the destination of the audit logs.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_lustre_configs",
				Description:   "The configuration for the Amazon FSx for Lustre file system.",
				Resolver:      fetchFsxFileSystemLustreConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "deployment_type",
						Description: "The deployment type of the FSx for Lustre file system.",
						Type:        schema.TypeString,
					},
					{
						Name:        "per_unit_storage_throughput",
						Description: "Per unit storage throughput represents the megabytes per second of read or write throughput per 1 tebibyte of storage provisioned.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "mount_name",
						Description: "Mount name of the FSx Lustre file system.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM.",
						Type:        schema.TypeString,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A boolean flag indicating whether tags on the file system are copied to backups.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "drive_cache_type",
						Description: "The type of drive cache used by PERSISTENT_1 file systems that are provisioned with HDD storage devices.",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_compression_type",
						Description: "The data compression configuration for the file system.",
						Type:        schema.TypeString,
					},
					{
						Name:        "log_level",
						Description: "The data repository events that are logged by Amazon FSx.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogConfiguration.Level"),
						// IgnoreInTests: true,
					},
					{
						Name:        "log_destination",
						Description: "The Amazon Resource Name (ARN) that specifies the destination of the logs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogConfiguration.Destination"),
						// IgnoreInTests: true,
					},
					{
						Name:        "root_squash",
						Description: "You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534 ).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RootSquashConfiguration.RootSquash"),
						// IgnoreInTests: true,
					},
					{
						Name:        "no_squash_nids",
						Description: "When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply.",
						Type:        schema.TypeIntArray,
						Resolver:    schema.PathResolver("RootSquashConfiguration.NoSquashNids"),
						// IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_fsx_file_system_lustre_data_repository_associations",
						Description:   "The data repository configuration object for Lustre file systems returned in the response of the CreateFileSystem operation.",
						Resolver:      fetchFsxFileSystemLustreDataRepositoryAssociations,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "lustre_config_cq_id",
								Description: "Unique CloudQuery ID of aws_fsx_file_system_lustre_configs table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "import_path",
								Description: "The import path to the Amazon S3 bucket (and optional prefix) that you're using as the data repository for your FSx for Lustre file system.",
								Type:        schema.TypeString,
							},
							{
								Name:        "export_path",
								Description: "The export path to the Amazon S3 bucket (and prefix) that you are using to store new and changed Lustre file system files in S3.",
								Type:        schema.TypeString,
							},
							{
								Name:        "imported_file_chunk_size",
								Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "auto_import_policy",
								Description: "Describes the file system's linked S3 data repository's AutoImportPolicy. The AutoImportPolicy configures how Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket.",
								Type:        schema.TypeString,
							},
							{
								Name:          "failure_details_message",
								Description:   "A message describing the data repository association failure.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("FailureDetails.Message"),
								IgnoreInTests: true,
							},
						},
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_administrative_actions",
				Description:   "A list of administrative actions for the file system that are in process or waiting to be processed.",
				Resolver:      fetchFsxFileSystemAdministrativeActions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action_type",
						Description: "Describes the type of administrative action.",
						Type:        schema.TypeString,
					},
					{
						Name:        "progress_percent",
						Description: "The percentage-complete status of a STORAGE_OPTIMIZATION administrative action.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "request_time",
						Description: "The time that the administrative action request was received.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "status",
						Description: "Describes the status of the administrative action.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_ontap_configs",
				Description:   "The configuration for this Amazon FSx for NetApp ONTAP file system.",
				Resolver:      fetchFsxFileSystemOntapConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the FSx for ONTAP file system deployment type in use in the file system.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_ip_address_range",
						Description: "(Multi-AZ only) The IP address range in which the endpoints to access your file system are created.",
						Type:        schema.TypeString,
					},
					{
						Name:        "intercluster_dns_name",
						Description: "The Domain Name Service (DNS) name for the file system intercluster endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Intercluster.DNSName"),
					},
					{
						Name:        "intercluster_ip_addresses",
						Description: "IP addresses of the file system intercluster endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Intercluster.IpAddresses"),
					},
					{
						Name:        "management_dns_name",
						Description: "The Domain Name Service (DNS) name for the file system management endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Management.DNSName"),
					},
					{
						Name:        "management_ip_addresses",
						Description: "IP addresses of the file system management endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Management.IpAddresses"),
					},
					{
						Name:        "disk_iops_mode",
						Description: "Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Mode"),
					},
					{
						Name:        "disk_iops",
						Description: "The total number of SSD IOPS provisioned for the file system.",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Iops"),
					},
					{
						Name:        "preferred_subnet_id",
						Description: "The ID for a subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "route_table_ids",
						Description: "(Multi-AZ only) The VPC route tables in which your file system's endpoints are created.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "throughput_capacity",
						Description: "The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps).",
						Type:        schema.TypeInt,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "A recurring weekly time, in the format D:HH:MM .",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_fsx_file_system_openzfs_configs",
				Description:   "The configuration for this Amazon FSx for OpenZFS file system.",
				Resolver:      fetchFsxFileSystemOpenZFSConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "fsx_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_file_systems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A boolean flag indicating whether tags on the file system should be copied to backups.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "copy_tags_to_volumes",
						Description: "A Boolean value indicating whether tags for the volume should be copied to snapshots.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the file-system deployment type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "throughput_capacity",
						Description: "The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps).",
						Type:        schema.TypeInt,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "A recurring weekly time, in the format D:HH:MM .",
						Type:        schema.TypeString,
					},
					{
						Name:        "disk_iops_mode",
						Description: "Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Mode"),
					},
					{
						Name:        "disk_iops",
						Description: "The total number of SSD IOPS provisioned for the file system.",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Iops"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchFsxFileSystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config fsx.DescribeFileSystemsInput
	c := meta.(*client.Client)
	svc := c.Services().FSX
	for {
		response, err := svc.DescribeFileSystems(ctx, &config, func(options *fsx.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.FileSystems
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchFsxFileSystemWindowsConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.WindowsConfiguration
	return nil
}

func fetchFsxFileSystemSelfManagedActiveDirectoryConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	windowsCfg := parent.Item.(types.WindowsFileSystemConfiguration)
	res <- windowsCfg.SelfManagedActiveDirectoryConfiguration
	return nil
}

func fetchFsxFileSystemAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config fsx.DescribeFileSystemAliasesInput
	c := meta.(*client.Client)
	svc := c.Services().FSX
	for {
		response, err := svc.DescribeFileSystemAliases(ctx, &config, func(options *fsx.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Aliases
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchFsxFileSystemAuditLogConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.WindowsConfiguration.AuditLogConfiguration
	return nil
}

func fetchFsxFileSystemLustreConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.LustreConfiguration
	return nil
}

func fetchFsxFileSystemLustreDataRepositoryAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lustreCfg := parent.Item.(types.LustreFileSystemConfiguration)
	res <- lustreCfg.DataRepositoryConfiguration
	return nil
}

func fetchFsxFileSystemAdministrativeActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.AdministrativeActions
	return nil
}

func fetchFsxFileSystemOntapConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.OntapConfiguration
	return nil
}

func fetchFsxFileSystemOpenZFSConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fs := parent.Item.(types.FileSystem)
	res <- fs.OpenZFSConfiguration
	return nil
}
