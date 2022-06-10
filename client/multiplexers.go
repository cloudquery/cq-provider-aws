package client

import (
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for partition := range client.ServicesManager.services {
		for accountID := range client.ServicesManager.services[partition] {
			l = append(l, client.withPartitionAccountID(partition, accountID))
		}
	}
	return l
}

func ServiceAccountRegionMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.Logger().Trace("region is not supported for service", "service", service, "region", region, "partition", partition)
						continue
					}
					l = append(l, client.withPartitionAccountIDAndRegion(partition, accountID, region))
				}
			}
		}
		return l
	}
}

func ServiceAccountRegionNamespaceMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
						continue
					}
					for _, ns := range AllNamespaces {
						l = append(l, client.withPartitionAccountIDRegionAndNamespace(partition, accountID, region, ns))
					}
				}
			}
		}
		return l
	}
}

func ServiceAccountRegionScopeMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				// always fetch cloudfront related resources
				l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, cloudfrontScopeRegion, wafv2types.ScopeCloudfront))
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
						continue
					}
					l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, region, wafv2types.ScopeRegional))
				}
			}
		}
		return l
	}
}
