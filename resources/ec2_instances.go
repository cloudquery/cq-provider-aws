package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Instances() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_instances",
		Description:  "Describes an instance.",
		Resolver:     fetchEc2Instances,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Description: "The Amazon Resource Name (ARN) for the ec2 instance",
				Type:        schema.TypeString,
				Resolver:    resolveEc2instanceTagsArn,
			},
			{
				Name:        "id",
				Description: "The ID of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceId"),
			},
			{
				Name:        "ami_launch_index",
				Description: "The AMI launch index, which can be used to find this instance in the launch group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "capacity_reservation_id",
				Description: "The ID of the Capacity Reservation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cap_reservation_preference",
				Description: "Describes the instance's Capacity Reservation preferences.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CapacityReservationSpecification.CapacityReservationPreference"),
			},
			{
				Name:        "cap_reservation_target_capacity_reservation_id",
				Description: "The ID of the targeted Capacity Reservation.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationId"),
			},
			{
				Name:        "cap_reservation_target_capacity_reservation_rg_arn",
				Description: "The ARN of the targeted Capacity Reservation group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupArn"),
			},
			{
				Name:        "client_token",
				Description: "The idempotency token you provided when you launched the instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cpu_options_core_count",
				Description: "The number of CPU cores for the instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CpuOptions.CoreCount"),
			},
			{
				Name:        "cpu_options_threads_per_core",
				Description: "The number of threads per CPU core.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CpuOptions.ThreadsPerCore"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the instance is optimized for Amazon EBS I/O.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enclave_options_enabled",
				Description: "If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves; otherwise, it is not enabled for AWS Nitro Enclaves.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EnclaveOptions.Enabled"),
			},
			{
				Name:        "hibernation_options_configured",
				Description: "If this parameter is set to true, your instance is enabled for hibernation; otherwise, it is not enabled for hibernation.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HibernationOptions.Configured"),
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_instance_profile_arn",
				Description: "The Amazon Resource Name (ARN) of the instance profile.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IamInstanceProfile.Arn"),
			},
			{
				Name:        "iam_instance_profile_id",
				Description: "The ID of the instance profile.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IamInstanceProfile.Id"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the AMI used to launch the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_lifecycle",
				Description: "Indicates whether this is a Spot Instance or a Scheduled Instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_type",
				Description: "The instance type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kernel_id",
				Description: "The kernel associated with this instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_name",
				Description: "The name of the key pair, if this instance was launched with an associated key pair.",
				Type:        schema.TypeString,
			},
			{
				Name:        "launch_time",
				Description: "The time the instance was launched.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "metadata_options_http_endpoint",
				Description: "This parameter enables or disables the HTTP metadata endpoint on your instances.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpEndpoint"),
			},
			{
				Name:        "metadata_options_http_put_response_hop_limit",
				Description: "The desired HTTP PUT response hop limit for instance metadata requests.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MetadataOptions.HttpPutResponseHopLimit"),
			},
			{
				Name:        "metadata_options_http_tokens",
				Description: "The state of token usage for your instance metadata requests.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpTokens"),
			},
			{
				Name:        "metadata_options_state",
				Description: "The state of the metadata option changes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.State"),
			},
			{
				Name:        "monitoring_state",
				Description: "Indicates whether detailed monitoring is enabled.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Monitoring.State"),
			},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost.",
				Type:        schema.TypeString,
			},
			{
				Name:        "placement_affinity",
				Description: "The affinity setting for the instance on the Dedicated Host.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.Affinity"),
			},
			{
				Name:        "placement_availability_zone",
				Description: "The Availability Zone of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.AvailabilityZone"),
			},
			{
				Name:        "placement_group_name",
				Description: "The name of the placement group the instance is in.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.GroupName"),
			},
			{
				Name:        "placement_host_id",
				Description: "The ID of the Dedicated Host on which the instance resides.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.HostId"),
			},
			{
				Name:        "placement_host_resource_group_arn",
				Description: "The ARN of the host resource group in which to launch the instances.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.HostResourceGroupArn"),
			},
			{
				Name:        "placement_partition_number",
				Description: "The number of the partition the instance is in.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Placement.PartitionNumber"),
			},
			{
				Name:        "placement_spread_domain",
				Description: "Reserved for future use.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.SpreadDomain"),
			},
			{
				Name:        "placement_tenancy",
				Description: "The tenancy of the instance (if the instance is running in a VPC).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.Tenancy"),
			},
			{
				Name:        "platform",
				Description: "The value is Windows for Windows instances; otherwise blank.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_name",
				Description: "(IPv4 only) The private DNS hostname name assigned to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_address",
				Description: "The private IPv4 address assigned to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_dns_name",
				Description: "(IPv4 only) The public DNS name assigned to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_ip_address",
				Description: "The public IPv4 address, or the Carrier IP address assigned to the instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ramdisk_id",
				Description: "The RAM disk associated with this instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_type",
				Description: "The root device type used by the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_dest_check",
				Description: "Specifies whether to enable an instance launched in a VPC to perform NAT.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "spot_instance_request_id",
				Description: "If the request is a Spot Instance request, the ID of the request.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sriov_net_support",
				Description: "Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_code",
				Description: "The state of the instance as a 16-bit unsigned integer.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:        "state_name",
				Description: "The current state of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Name"),
			},
			{
				Name:        "state_reason_code",
				Description: "The reason code for the state change.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StateReason.Code"),
			},
			{
				Name:        "state_reason_message",
				Description: "The message for the state change.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StateReason.Message"),
			},
			{
				Name:        "state_transition_reason",
				Description: "The reason for the most recent state transition.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet in which the instance is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the instance.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2instanceTags,
			},
			{
				Name:        "virtualization_type",
				Description: "The virtualization type of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the instance is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "licenses",
				Description: "The license configurations.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveEc2InstanceLicenses,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_instance_block_device_mappings",
				Description: "Describes a block device mapping.",
				Resolver:    fetchEc2InstanceBlockDeviceMappings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "ebs_volume_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "device_name",
						Description: "The device name (for example, /dev/sdh or xvdh).",
						Type:        schema.TypeString,
					},
					{
						Name:        "ebs_attach_time",
						Description: "The time stamp when the attachment initiated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Ebs.AttachTime"),
					},
					{
						Name:        "ebs_delete_on_termination",
						Description: "Indicates whether the volume is deleted on instance termination.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:        "ebs_status",
						Description: "The attachment state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.Status"),
					},
					{
						Name:        "ebs_volume_id",
						Description: "The ID of the EBS volume.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.VolumeId"),
					},
				},
			},
			{
				Name:        "aws_ec2_instance_elastic_gpu_associations",
				Description: "Describes the association between an instance and an Elastic Graphics accelerator.",
				Resolver:    fetchEc2InstanceElasticGpuAssociations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "elastic_gpu_association_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "elastic_gpu_association_id",
						Description: "The ID of the association.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_association_state",
						Description: "The state of the association between the instance and the Elastic Graphics accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_association_time",
						Description: "The time the Elastic Graphics accelerator was associated with the instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_id",
						Description: "The ID of the Elastic Graphics accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_arn",
						Description: "The Amazon Resource Name (ARN) for the ec2 elastic gpu",
						Type:        schema.TypeString,
						Resolver:    resolveEc2InstanceElasticGpuAssociationsArn,
					},
				},
			},
			{
				Name:        "aws_ec2_instance_elastic_inference_accelerator_associations",
				Description: "Describes the association between an instance and an elastic inference accelerator.",
				Resolver:    fetchEc2InstanceElasticInferenceAcceleratorAssociations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "elastic_inference_accelerator_association_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "elastic_inference_accelerator_arn",
						Description: "The Amazon Resource Name (ARN) of the elastic inference accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_id",
						Description: "The ID of the association.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_state",
						Description: "The state of the elastic inference accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_time",
						Description: "The time at which the elastic inference accelerator is associated with an instance.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "aws_ec2_instance_network_interfaces",
				Description: "Describes a network interface.",
				Resolver:    fetchEc2InstanceNetworkInterfaces,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "network_interface_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "association_carrier_ip",
						Description: "The carrier IP address associated with the network interface.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.CarrierIp"),
					},
					{
						Name:        "association_ip_owner_id",
						Description: "The ID of the owner of the Elastic IP address.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.IpOwnerId"),
					},
					{
						Name:        "association_public_dns_name",
						Description: "The public DNS name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.PublicDnsName"),
					},
					{
						Name:        "association_public_ip",
						Description: "The public IP address or Elastic IP address bound to the network interface.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.PublicIp"),
					},
					{
						Name:        "attachment_attach_time",
						Description: "The time stamp when the attachment initiated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Attachment.AttachTime"),
					},
					{
						Name:        "attachment_id",
						Description: "The ID of the network interface attachment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attachment.AttachmentId"),
					},
					{
						Name:        "attachment_delete_on_termination",
						Description: "Indicates whether the network interface is deleted when the instance is terminated.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Attachment.DeleteOnTermination"),
					},
					{
						Name:        "attachment_device_index",
						Description: "The index of the device on the instance for the network interface attachment.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attachment.DeviceIndex"),
					},
					{
						Name:        "attachment_network_card_index",
						Description: "The index of the network card.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attachment.NetworkCardIndex"),
					},
					{
						Name:        "attachment_status",
						Description: "The attachment state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attachment.Status"),
					},
					{
						Name:        "description",
						Description: "The description.",
						Type:        schema.TypeString,
					},
					{
						Name:        "interface_type",
						Description: "Describes the type of network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "mac_address",
						Description: "The MAC address.",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_interface_id",
						Description: "The ID of the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner_id",
						Description: "The ID of the AWS account that created the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_dns_name",
						Description: "The private DNS name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_ip_address",
						Description: "The IPv4 address of the network interface within the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source_dest_check",
						Description: "Indicates whether to validate network traffic to or from this network interface.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "status",
						Description: "The status of the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_id",
						Description: "The ID of the VPC.",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the ec2 network interface",
						Type:        schema.TypeString,
						Resolver:    resolveEc2InstanceNetworkInterfacesArn,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ec2_instance_network_interface_groups",
						Description: "Describes a security group.",
						Resolver:    fetchEc2InstanceNetworkInterfaceGroups,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_network_interface_cq_id", "group_id"}},
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "network_interface_id",
								Description: "The ID of the network interface.",
								Type:        schema.TypeString,
								Resolver:    schema.ParentPathResolver("NetworkInterfaceId"),
							},
							{
								Name:        "group_id",
								Description: "The ID of the security group.",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_name",
								Description: "The name of the security group.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_instance_network_interface_ipv6_addresses",
						Description: "Describes an IPv6 address.",
						Resolver:    fetchEc2InstanceNetworkInterfaceIpv6Addresses,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "ipv6_address",
								Description: "The IPv6 address.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_instance_network_interface_private_ip_addresses",
						Description: "Describes a private IPv4 address.",
						Resolver:    fetchEc2InstanceNetworkInterfacePrivateIpAddresses,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "association_carrier_ip",
								Description: "The carrier IP address associated with the network interface.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Association.CarrierIp"),
							},
							{
								Name:        "association_ip_owner_id",
								Description: "The ID of the owner of the Elastic IP address.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Association.IpOwnerId"),
							},
							{
								Name:        "association_public_dns_name",
								Description: "The public DNS name.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Association.PublicDnsName"),
							},
							{
								Name:        "association_public_ip",
								Description: "The public IP address or Elastic IP address bound to the network interface.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Association.PublicIp"),
							},
							{
								Name:        "is_primary",
								Description: "Indicates whether this IPv4 address is the primary private IP address of the network interface.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Primary"),
							},
							{
								Name:        "private_dns_name",
								Description: "The private IPv4 DNS name.",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_ip_address",
								Description: "The private IPv4 address of the network interface.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_ec2_instance_product_codes",
				Description: "Describes a product code.",
				Resolver:    fetchEc2InstanceProductCodes,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "product_code_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "product_code_id",
						Description: "The product code.",
						Type:        schema.TypeString,
					},
					{
						Name:        "product_code_type",
						Description: "The type of product code.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_ec2_instance_security_groups",
				Description: "Describes a security group.",
				Resolver:    fetchEc2InstanceSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "group_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "group_id",
						Description: "The ID of the security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_name",
						Description: "The name of the security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the ec2 security group",
						Type:        schema.TypeString,
						Resolver:    resolveEc2InstanceSecurityGroupsArn,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2

	response, err := svc.DescribeInstances(ctx, &ec2.DescribeInstancesInput{}, func(o *ec2.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}

	for _, reservation := range response.Reservations {
		res <- reservation.Instances
	}

	return nil
}
func resolveEc2instanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2InstanceBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.BlockDeviceMappings
	return nil
}
func fetchEc2InstanceElasticGpuAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ElasticGpuAssociations
	return nil
}
func fetchEc2InstanceElasticInferenceAcceleratorAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ElasticInferenceAcceleratorAssociations
	return nil
}
func resolveEc2InstanceLicenses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance, ok := resource.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	licenses := make([]string, len(instance.Licenses))
	for i, l := range instance.Licenses {
		licenses[i] = *l.LicenseConfigurationArn
	}
	return resource.Set(c.Name, licenses)
}
func fetchEc2InstanceNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.NetworkInterfaces
	return nil
}
func fetchEc2InstanceNetworkInterfaceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.Groups
	return nil
}
func fetchEc2InstanceNetworkInterfaceIpv6Addresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.Ipv6Addresses
	return nil
}
func fetchEc2InstanceNetworkInterfacePrivateIpAddresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.PrivateIpAddresses
	return nil
}
func fetchEc2InstanceProductCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ProductCodes
	return nil
}
func fetchEc2InstanceSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.SecurityGroups
	return nil
}
func resolveEc2instanceTagsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ins := resource.Item.(types.Instance)
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "instance", *ins.InstanceId, cl.Region, cl.AccountID))
}
func resolveEc2InstanceElasticGpuAssociationsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	gpu := resource.Item.(types.ElasticGpuAssociation)
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "elastic-gpu", *gpu.ElasticGpuId, cl.Region, cl.AccountID))
}
func resolveEc2InstanceSecurityGroupsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	sg := resource.Item.(types.SecurityGroup)
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "security-group", *sg.GroupId, cl.Region, cl.AccountID))
}
func resolveEc2InstanceNetworkInterfacesArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ni := resource.Item.(types.NetworkInterface)
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "network-interface", *ni.NetworkInterfaceId, cl.Region, cl.AccountID))
}
