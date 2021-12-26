package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func awsTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) client.Services, options TestOptions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := `
		regions = ["us-east-1"]
		accounts "testAccount" {
			role_arn = ""
		}
		aws_debug = false
		max_retries = 3
		max_backoff = 60
	`
	accounts := []client.Account{
		{ID: "testAccount"},
	}

	providertest.TestResource(t, Provider, providertest.ResourceTestData{
		Table:  table,
		Config: cfg,
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
			c := client.NewAwsClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), accounts, []string{"us-east-1"})
			c.ServicesManager.InitServicesForAccountAndRegion("testAccount", "us-east-1", builder(t, ctrl))
			return &c, nil
		},
		SkipEmptyJsonB: options.SkipEmptyJsonB,
	})
}
