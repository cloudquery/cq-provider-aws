package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		client.Logger().Info("account id", "account_id", accountID)
		l = append(l, client.withAccountID(accountID))
	}
	return l
}

func ServiceAccountRegionMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for i, account := range client.Accounts {
			for _, region := range account.Regions {
				// if !isSupportedServiceForRegion(service, region) {
				// 	meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
				// 	continue
				// }
				l = append(l, client.withAccountIDAndRegion(account.AccountID, region, client.Cfgs[i]))
			}
		}
		return l
	}
}

// func ServiceAccountRegionNamespaceMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
// 	return func(meta schema.ClientMeta) []schema.ClientMeta {
// 		var l = make([]schema.ClientMeta, 0)
// 		client := meta.(*Client)
// 		for accountID := range client.ServicesManager.services {
// 			for region := range client.ServicesManager.services[accountID] {
// 				if !isSupportedServiceForRegion(service, region) {
// 					meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
// 					continue
// 				}
// 				for _, ns := range AllNamespaces {
// 					l = append(l, client.withAccountIDRegionAndNamespace(accountID, region, ns))
// 				}
// 			}
// 		}
// 		return l
// 	}
// }

// func ServiceAccountRegionScopeMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
// 	return func(meta schema.ClientMeta) []schema.ClientMeta {
// 		var l = make([]schema.ClientMeta, 0)
// 		client := meta.(*Client)
// 		for accountID := range client.ServicesManager.services {
// 			// always fetch cloudfront related resources
// 			l = append(l, client.withAccountIDRegionAndScope(accountID, cloudfrontScopeRegion, wafv2types.ScopeCloudfront))
// 			for region := range client.ServicesManager.services[accountID] {
// 				if !isSupportedServiceForRegion(service, region) {
// 					meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
// 					continue
// 				}
// 				l = append(l, client.withAccountIDRegionAndScope(accountID, region, wafv2types.ScopeRegional))
// 			}
// 		}
// 		return l
// 	}
// }
