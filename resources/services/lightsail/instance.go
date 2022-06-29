package lightsail

import (
	"context"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource instance --config resources/services/lightsail/gen.hcl --output .
func Instances() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_instance",
		Description:  "Describes an instance (a virtual private server).",
		Resolver:     fetchLightsailInstances,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Description: "The Amazon Resource Name (ARN) of the instance (e.g., arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE).",
				Type:        schema.TypeString,
			},
			{
				Name:        "blueprint_id",
				Description: "The blueprint ID (e.g., os_amlinux_2016_03).",
				Type:        schema.TypeString,
			},
			{
				Name:        "blueprint_name",
				Description: "The friendly name of the blueprint (e.g., Amazon Linux).",
				Type:        schema.TypeString,
			},
			{
				Name:        "bundle_id",
				Description: "The bundle for the instance (e.g., micro_1_0).",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the instance was created (e.g., 1479734909.17) in Unix time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "hardware_cpu_count",
				Description: "The number of vCPUs the instance has.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Hardware.CpuCount"),
			},
			{
				Name:        "hardware_ram_size_in_gb",
				Description: "The amount of RAM in GB on the instance (e.g., 1.0).",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Hardware.RamSizeInGb"),
			},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_addresses",
				Description: "The IPv6 addresses of the instance.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "is_static_ip",
				Description: "A Boolean value indicating whether this instance has a static IP assigned to it.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "location_availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "location_region_name",
				Description: "The AWS Region name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.RegionName"),
			},
			{
				Name:        "name",
				Description: "The name the user gave the instance (e.g., Amazon_Linux-1GB-Ohio-1).",
				Type:        schema.TypeString,
			},
			{
				Name:        "networking_monthly_transfer_gb_per_month_allocated",
				Description: "The amount allocated per month (in GB).",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Networking.MonthlyTransfer.GbPerMonthAllocated"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The type of resource (usually Instance).",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssh_key_name",
				Description: "The name of the SSH key being used to connect to the instance (e.g., LightsailDefaultKeyPair).",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_code",
				Description: "The status code for the instance.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:        "state_name",
				Description: "The state of the instance (e.g., running or pending).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Name"),
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "username",
				Description: "The user name for connecting to the instance (e.g., ec2-user).",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_instance_add_ons",
				Description: "Describes an add-on that is enabled for an Amazon Lightsail resource.",
				Resolver:    fetchLightsailInstanceAddOns,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instance table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the add-on.",
						Type:        schema.TypeString,
					},
					{
						Name:        "next_snapshot_time_of_day",
						Description: "The next daily time an automatic snapshot will be created",
						Type:        schema.TypeString,
					},
					{
						Name:        "snapshot_time_of_day",
						Description: "The daily time when an automatic snapshot is created",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the add-on.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_hardware_disks",
				Description: "Describes a block storage disk.",
				Resolver:    fetchLightsailInstanceHardwareDisks,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instance table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the disk.",
						Type:        schema.TypeString,
					},
					{
						Name:        "attached_to",
						Description: "The resources to which the disk is attached.",
						Type:        schema.TypeString,
					},
					{
						Name:        "attachment_state",
						Description: "(Deprecated) The attachment state of the disk",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The date when the disk was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "gb_in_use",
						Description: "(Deprecated) The number of GB in use by the disk",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "iops",
						Description: "The input/output operations per second (IOPS) of the disk.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "is_attached",
						Description: "A Boolean value indicating whether the disk is attached.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "is_system_disk",
						Description: "A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it).",
						Type:        schema.TypeBool,
					},
					{
						Name:        "location_availability_zone",
						Description: "The Availability Zone",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.AvailabilityZone"),
					},
					{
						Name:        "location_region_name",
						Description: "The AWS Region name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.RegionName"),
					},
					{
						Name:        "name",
						Description: "The unique name of the disk.",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The disk path.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The Lightsail resource type (e.g., Disk).",
						Type:        schema.TypeString,
					},
					{
						Name:        "size_in_gb",
						Description: "The size of the disk in GB.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "state",
						Description: "Describes the status of the disk.",
						Type:        schema.TypeString,
					},
					{
						Name:        "support_code",
						Description: "The support code",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_lightsail_instance_hardware_disk_add_ons",
						Description: "Describes an add-on that is enabled for an Amazon Lightsail resource.",
						Resolver:    fetchLightsailInstanceHardwareDiskAddOns,
						Columns: []schema.Column{
							{
								Name:        "instance_hardware_disk_cq_id",
								Description: "Unique CloudQuery ID of aws_lightsail_instance_hardware_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of the add-on.",
								Type:        schema.TypeString,
							},
							{
								Name:        "next_snapshot_time_of_day",
								Description: "The next daily time an automatic snapshot will be created",
								Type:        schema.TypeString,
							},
							{
								Name:        "snapshot_time_of_day",
								Description: "The daily time when an automatic snapshot is created",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the add-on.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_lightsail_instance_hardware_disk_tags",
						Description: "Describes a tag key and optional value assigned to an Amazon Lightsail resource. For more information about tags in Lightsail, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-tags).",
						Resolver:    fetchLightsailInstanceHardwareDiskTags,
						Columns: []schema.Column{
							{
								Name:        "instance_hardware_disk_cq_id",
								Description: "Unique CloudQuery ID of aws_lightsail_instance_hardware_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key",
								Description: "The key of the tag",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The value of the tag",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_networking_ports",
				Description: "Describes information about ports for an Amazon Lightsail instance.",
				Resolver:    fetchLightsailInstanceNetworkingPorts,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instance table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "access_direction",
						Description: "The access direction (inbound or outbound)",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_from",
						Description: "The location from which access is allowed",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_type",
						Description: "The type of access (Public or Private).",
						Type:        schema.TypeString,
					},
					{
						Name:        "cidr_list_aliases",
						Description: "An alias that defines access for a preconfigured range of IP addresses",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "cidrs",
						Description: "The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "common_name",
						Description: "The common name of the port information.",
						Type:        schema.TypeString,
					},
					{
						Name:        "from_port",
						Description: "The first port in a range of open ports on an instance",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "ipv6_cidrs",
						Description: "The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "protocol",
						Description: "The IP protocol name",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_port",
						Description: "The last port in a range of open ports on an instance",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_tags",
				Description: "Describes a tag key and optional value assigned to an Amazon Lightsail resource. For more information about tags in Lightsail, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-tags).",
				Resolver:    fetchLightsailInstanceTags,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instance table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "The key of the tag",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The value of the tag",
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

func fetchLightsailInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceAddOns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceHardwareDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceHardwareDiskAddOns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceHardwareDiskTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceNetworkingPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchLightsailInstanceTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
