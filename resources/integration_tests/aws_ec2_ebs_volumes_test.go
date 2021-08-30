package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2EbsVolumes(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2EbsVolumes(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2EbsVolumes().Name,
			//Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
			//	return sq.Where(squirrel.And{
			//		squirrel.Eq{"tags->>'TestId'": res.Suffix},
			//		squirrel.NotEq{"state_name": "terminated"},
			//	})
			//},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"volume_type":          "gp2",
						"multi_attach_enabled": false,
						"encrypted":            false,
						"fast_restored":        false,
						"size":                 float64(5),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ebs-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
