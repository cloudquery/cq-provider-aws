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

func TestMultiplexerPKMatch(t *testing.T) {
	t.Parallel()
	for res, table := range Provider().ResourceMap {
		if table.Multiplex == nil {
			continue
		}

		table := table
		t.Run(res, func(t *testing.T) {
			t.Parallel()
			cfgRegions := []string{"us-east-1", "eu-west-1"}

			c := client.NewAwsClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []client.Account{{ID: "test1"}, {ID: "test2"}}, cfgRegions)
			for _, a := range c.Accounts {
				for _, r := range cfgRegions {
					c.ServicesManager.InitServicesForAccountAndRegion(a.AccountID, r, client.Services{})
				}
			}

			multiClients := table.Multiplex(schema.ClientMeta(&c))

			regions, accounts := make(map[string]struct{}), make(map[string]struct{})
			for _, c := range multiClients {
				cc := c.(*client.Client)
				accounts[cc.AccountID], regions[cc.Region] = struct{}{}, struct{}{}
			}
			requiredPKs := make(map[string]struct{})
			if len(accounts) > 1 {
				requiredPKs["account_id"] = struct{}{}
			}
			if len(regions) > 1 {
				requiredPKs["region"] = struct{}{}
			}
			if len(requiredPKs) == 0 {
				return
			}
			// Multiple accounts OR regions confirmed

			for _, c := range table.PrimaryKeys() {
				if c == "arn" { // ARN satisfies both account_id and region
					delete(requiredPKs, "account_id")
					delete(requiredPKs, "region")
				}
				delete(requiredPKs, c)
			}

			assert.Zero(t, len(requiredPKs), "Missing PKs: %+v", requiredPKs)
		})
	}
}
