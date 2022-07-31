package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/plugin/source"
	"github.com/cloudquery/cq-provider-sdk/plugin/source/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/plugin/source/testing"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := `
max_goroutines: 100
configuration:
  regions: ["us-east-1"]
  accounts:
  - id: testAccount
    role_arn: ""
  aws_debug: false
  max_retries: 3
  max_backoff: 60
tables:
  - "*"
`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Plugin: &source.SourcePlugin{
			Name:    "aws_mock_test_provider",
			Version: "development",
			Configure: func(logger zerolog.Logger, i interface{}) (schema.ClientMeta, error) {
				c := NewAwsClient(logger)
				c.ServicesManager.InitServicesForPartitionAccountAndRegion("aws", "testAccount", "us-east-1", builder(t, ctrl))
				c.Partition = "aws"
				return &c, nil
			},
			Tables: []*schema.Table{
				table,
			},
			Config: func() interface{} {
				return &Config{}
			},
		},
		Config:           cfg,
		SkipIgnoreInTest: true,
	})
}
