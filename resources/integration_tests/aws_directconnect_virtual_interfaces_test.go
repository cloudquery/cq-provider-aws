package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDirectconnectVirtualInterfaces(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectVirtualInterfaces(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_directconnect_virtual_interfaces",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{squirrel.Eq{"virtual_interface_name": fmt.Sprintf("vif-%s-%s", res.Prefix, res.Suffix)}, squirrel.NotEq{"virtual_interface_state": "deleted"}})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"virtual_interface_name": fmt.Sprintf("vif-%s-%s", res.Prefix, res.Suffix),
						"address_family":         "ipv4",
						"amazon_address":         "175.45.176.2/30",
						"asn":                    float64(65352),
						"customer_address":       "175.45.176.1/30",
						//"route_filter_prefixes":  []interface{}{"210.52.109.0/24", "175.45.176.0/22"},
						"tags": map[string]interface{}{"Type": "integration_test", "TestId": "windowsfifl5fe"},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_directconnect_virtual_interface_bgp_peers",
					ForeignKeyName: "virtual_interface_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"address_family":   "ipv4",
								"amazon_address":   "175.45.176.2/30",
								"asn":              float64(65352),
								"customer_address": "175.45.176.1/30",
							},
						},
					},
				},
			}}
	})
}
