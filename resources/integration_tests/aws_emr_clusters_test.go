package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEmrClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EmrClusters(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_emr_clusters",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("emr-cluster-%s%s", res.Prefix, res.Suffix),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{squirrel.Eq{"name": fmt.Sprintf("emr-cluster-%s%s", res.Prefix, res.Suffix)},
					squirrel.NotEq{"status_state": "TERMINATED_WITH_ERRORS"}})
			},
		}
	})
}
