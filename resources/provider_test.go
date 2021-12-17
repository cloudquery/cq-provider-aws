package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func awsTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) client.Services, options TestOptions) {
	ctrl := gomock.NewController(t)

	cfg := client.Config{
		Regions:    []string{"us-east-1"},
		Accounts:   []client.Account{{ID: "testAccount", RoleARN: ""}},
		AWSDebug:   false,
		MaxRetries: 3,
		MaxBackoff: 60,
	}
	providertest.TestResource(t, Provider, providertest.ResourceTestData{
		Table:  table,
		Config: cfg,
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
			c := client.NewAwsClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), cfg.Accounts, []string{"us-east-1"})
			c.ServicesManager.InitServicesForAccountAndRegion("testAccount", "us-east-1", builder(t, ctrl))
			return &c, nil
		},
		SkipEmptyJsonB: options.SkipEmptyJsonB,
	})
}

func TestTableIdentifiers(t *testing.T) {
	t.Parallel()
	for _, res := range Provider().ResourceMap {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}

func testTableIdentifiers(t *testing.T, table *schema.Table) {
	const maxIdentifierLength = 63 // maximum allowed identifier length is 63 bytes https://www.postgresql.org/docs/13/limits.html

	assert.NotEmpty(t, table.Name)
	assert.LessOrEqual(t, len(table.Name), maxIdentifierLength, "Table name too long")

	for _, c := range table.Columns {
		assert.NotEmpty(t, c.Name)
		assert.LessOrEqual(t, len(c.Name), maxIdentifierLength, "Column name too long:", c.Name)
	}

	for _, res := range table.Relations {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}
