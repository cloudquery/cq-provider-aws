package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticsearchDomains() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticsearch_domains",
		Description:  "The current status of an Elasticsearch domain. ",
		Resolver:     fetchElasticsearchDomains,
		Multiplex:    client.AccountRegionMultiplex,
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
				Name:        "arn",
				Description: "The Amazon resource name (ARN) of an Elasticsearch domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "domain_id",
				Description: "The unique identifier for the specified Elasticsearch domain.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name",
				Description: "The name of an Elasticsearch domain",
				Type:        schema.TypeString,
			},
			{
				Name:        "elasticsearch_cluster_config",
				Description: "The type and number of instances in the domain cluster.  This member is required.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainElasticsearchClusterConfig,
			},
			{
				Name:        "access_policies",
				Description: "IAM access policy as a JSON-formatted string.",
				Type:        schema.TypeString,
			},
			{
				Name:        "advanced_options",
				Description: "Specifies the status of the AdvancedOptions",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "advanced_security_options",
				Description: "The current status of the Elasticsearch domain's advanced security options.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainAdvancedSecurityOptions,
			},
			{
				Name:        "auto_tune_options",
				Description: "The current status of the Elasticsearch domain's Auto-Tune options.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainAutoTuneOptions,
			},
			{
				Name:        "cognito_options",
				Description: "The CognitoOptions for the specified domain",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainCognitoOptions,
			},
			{
				Name:        "created",
				Description: "The domain creation status",
				Type:        schema.TypeBool,
			},
			{
				Name:        "deleted",
				Description: "The domain deletion status",
				Type:        schema.TypeBool,
			},
			{
				Name:        "domain_endpoint_options",
				Description: "The current status of the Elasticsearch domain's endpoint options.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainDomainEndpointOptions,
			},
			{
				Name:        "ebs_options",
				Description: "The EBSOptions for the specified domain",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainEbsOptions,
			},
			{
				Name: "elasticsearch_version",
				Type: schema.TypeString,
			},
			{
				Name:        "encryption_at_rest_options",
				Description: "Specifies the status of the EncryptionAtRestOptions.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainEncryptionAtRestOptions,
			},
			{
				Name:        "endpoint",
				Description: "The Elasticsearch domain endpoint that you use to submit index and search requests.",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoints",
				Description: "Map containing the Elasticsearch domain endpoints used to submit index and search requests",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "log_publishing_options",
				Description: "Log publishing options for the given domain.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "node_to_node_encryption_enabled",
				Description: "Specify true to enable node-to-node encryption.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NodeToNodeEncryptionOptions.Enabled"),
			},
			{
				Name:        "processing",
				Description: "The status of the Elasticsearch domain configuration",
				Type:        schema.TypeBool,
			},
			{
				Name:        "service_software_options",
				Description: "The current status of the Elasticsearch domain's service software.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainServiceSoftwareOptions,
			},
			{
				Name:        "snapshot_options_automated_snapshot_start_hour",
				Description: "Specifies the time, in UTC format, when the service takes a daily automated snapshot of the specified Elasticsearch domain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SnapshotOptions.AutomatedSnapshotStartHour"),
			},
			{
				Name:        "upgrade_processing",
				Description: "The status of an Elasticsearch domain version upgrade",
				Type:        schema.TypeBool,
			},
			{
				Name:        "vpc_options",
				Description: "The VPCOptions for the specified domain",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticsearchDomainVpcOptions,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	optsFunc := func(options *elasticsearchservice.Options) { options.Region = c.Region }
	svc := c.Services().ElasticSearch
	out, err := svc.ListDomainNames(ctx, &elasticsearchservice.ListDomainNamesInput{}, optsFunc)
	if err != nil {
		return err
	}
	for _, info := range out.DomainNames {
		domainOutput, err := svc.DescribeElasticsearchDomain(ctx, &elasticsearchservice.DescribeElasticsearchDomainInput{DomainName: info.DomainName}, optsFunc)
		if err != nil {
			return nil
		}
		res <- domainOutput.DomainStatus
	}
	return nil
}

func jsonEncoder(valueFetcher func(*types.ElasticsearchDomainStatus) interface{}) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		ds, ok := resource.Item.(*types.ElasticsearchDomainStatus)
		if !ok {
			return fmt.Errorf("not an ElasticsearchDomainStatus instance: %#v", resource.Item)
		}
		data, err := json.Marshal(valueFetcher(ds))
		if err != nil {
			return err
		}
		return resource.Set(c.Name, data)
	}
}

var (
	resolveElasticsearchDomainElasticsearchClusterConfig = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.ElasticsearchClusterConfig })
	resolveElasticsearchDomainAdvancedSecurityOptions    = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.AdvancedSecurityOptions })
	resolveElasticsearchDomainAutoTuneOptions            = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.AutoTuneOptions })
	resolveElasticsearchDomainCognitoOptions             = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.CognitoOptions })
	resolveElasticsearchDomainDomainEndpointOptions      = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.DomainEndpointOptions })
	resolveElasticsearchDomainEbsOptions                 = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.EBSOptions })
	resolveElasticsearchDomainEncryptionAtRestOptions    = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.EncryptionAtRestOptions })
	resolveElasticsearchDomainServiceSoftwareOptions     = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.ServiceSoftwareOptions })
	resolveElasticsearchDomainVpcOptions                 = jsonEncoder(func(ds *types.ElasticsearchDomainStatus) interface{} { return ds.VPCOptions })
)
