package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DaxClusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_dax_clusters",
		Description: "Contains all of the attributes of a specific DAX cluster.",
		Resolver:    fetchDaxClusters,
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
				Name:        "active_nodes",
				Description: "The number of nodes in the cluster that are active (i.e., capable of serving requests).",
				Type:        schema.TypeInt,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterArn"),
			},
			{
				Name:        "cluster_discovery_endpoint_address",
				Description: "The DNS hostname of the endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterDiscoveryEndpoint.Address"),
			},
			{
				Name:        "cluster_discovery_endpoint_port",
				Description: "The port number that applications should use to connect to the endpoint.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ClusterDiscoveryEndpoint.Port"),
			},
			{
				Name:        "cluster_discovery_endpoint_url",
				Description: "The URL that applications should use to connect to the endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterDiscoveryEndpoint.URL"),
			},
			{
				Name:        "cluster_endpoint_encryption_type",
				Description: "The type of encryption supported by the cluster's endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the DAX cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterName"),
			},
			{
				Name:        "description",
				Description: "The description of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_role_arn",
				Description: "A valid Amazon Resource Name (ARN) that identifies an IAM role",
				Type:        schema.TypeString,
			},
			{
				Name:        "node_ids_to_remove",
				Description: "A list of nodes to be removed from the cluster.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "node_type",
				Description: "The node type for the nodes in the cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "notification_configuration_topic_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the topic.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotificationConfiguration.TopicArn"),
			},
			{
				Name:        "notification_configuration_topic_status",
				Description: "The current state of the topic",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotificationConfiguration.TopicStatus"),
			},
			{
				Name:        "node_ids_to_reboot",
				Description: "The node IDs of one or more nodes to be rebooted.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ParameterGroup.NodeIdsToReboot"),
			},
			{
				Name:        "parameter_apply_status",
				Description: "The status of parameter updates.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ParameterGroup.ParameterApplyStatus"),
			},
			{
				Name:        "parameter_group_name",
				Description: "The name of the parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ParameterGroup.ParameterGroupName"),
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "A range of time when maintenance of DAX cluster software will be performed",
				Type:        schema.TypeString,
			},
			{
				Name:        "sse_description_status",
				Description: "The current state of server-side encryption:  * ENABLING - Server-side encryption is being enabled.  * ENABLED - Server-side encryption is enabled.  * DISABLING - Server-side encryption is being disabled.  * DISABLED - Server-side encryption is disabled.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SSEDescription.Status"),
			},
			{
				Name:        "security_groups",
				Description: "A list of security groups, and the status of each, for the nodes in the cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDaxClusterSecurityGroups,
			},
			{
				Name:        "status",
				Description: "The current status of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_group",
				Description: "The subnet group where the DAX cluster is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "total_nodes",
				Description: "The total number of nodes in the cluster.",
				Type:        schema.TypeInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_dax_cluster_nodes",
				Description: "Represents an individual node within a DAX cluster.",
				Resolver:    fetchDaxClusterNodes,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_dax_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "availability_zone",
						Description: "The Availability Zone (AZ) in which the node has been deployed.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_address",
						Description: "The DNS hostname of the endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoint.Address"),
					},
					{
						Name:        "endpoint_port",
						Description: "The port number that applications should use to connect to the endpoint.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Endpoint.Port"),
					},
					{
						Name:        "endpoint_url",
						Description: "The URL that applications should use to connect to the endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoint.URL"),
					},
					{
						Name:        "node_create_time",
						Description: "The date and time (in UNIX epoch format) when the node was launched.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "node_id",
						Description: "A system-generated identifier for the node.",
						Type:        schema.TypeString,
					},
					{
						Name:        "node_status",
						Description: "The current status of the node",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_group_status",
						Description: "The status of the parameter group associated with this node",
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
func fetchDaxClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DAX

	config := dax.DescribeClustersInput{}
	for {
		output, err := svc.DescribeClusters(ctx, &config, func(o *dax.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}

		for i := range output.Clusters {
			res <- output.Clusters[i]
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}

func resolveDaxClusterSecurityGroups(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Cluster)
	return resource.Set(c.Name, r.SecurityGroups)
}

func fetchDaxClusterNodes(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Cluster)
	for i := range r.Nodes {
		res <- r.Nodes[i]
	}
	return nil
}
