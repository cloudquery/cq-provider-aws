package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRedshiftClusters(t *testing.T) {
	const clusterFKName = "cluster_cq_id"
	awsTestIntegrationHelper(t, resources.RedshiftClusters(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_redshift_clusters",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"cluster_availability_status": "Available",
					"db_name":                     "mydb",
					"id":                          "tf-redshift-cluster",
					"master_username":             "foo",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_redshift_cluster_nodes",
					ForeignKeyName: clusterFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"node_role": "SHARED",
						},
					}},
				},
				{
					Name:           "aws_redshift_cluster_parameter_groups",
					ForeignKeyName: clusterFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"parameter_group_name": "parameter-group-test-terraform",
						},
					}},
				},
			},
		}
	})
	// TODO: relations not tested:
	// aws_redshift_cluster_security_groups
	// aws_redshift_cluster_deferred_maintenance_windows
	// aws_redshift_cluster_endpoint_vpc_endpoints
	// aws_redshift_cluster_iam_roles
	// aws_redshift_cluster_vpc_security_groups
}
