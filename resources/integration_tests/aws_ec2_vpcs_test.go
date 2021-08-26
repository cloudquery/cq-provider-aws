package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Vpcs(), []string{"aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpcs",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"is_default": false,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("vpc%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcPeeringConnections(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpc_peering_connections",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"tags->>'TestId'": res.Suffix},
					squirrel.NotEq{"status_code": "deleted"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"accepter_allow_dns_resolution_from_remote_vpc":           true,
						"accepter_allow_egress_local_classic_link_to_remote_vpc":  false,
						"accepter_allow_egress_local_vpc_to_remote_classic_link":  false,
						"requester_allow_dns_resolution_from_remote_vpc":          true,
						"requester_allow_egress_local_classic_link_to_remote_vpc": false,
						"requester_allow_egress_local_vpc_to_remote_classic_link": false,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
