package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

func awsTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) client.Services) {
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
			}), []string{"us-east-1"})
			c.ServicesManager.InitServicesForAccountAndRegion("testAccount", "us-east-1", builder(t, ctrl))
			return &c, nil
		},
	})
}

func awsTestIntegrationHelper(t *testing.T, table *schema.Table, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification) {
	cfg := client.Config{
		Regions:  []string{"us-east-1"},
		AWSDebug: false,
	}

	providertest.IntegrationTest(t, Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		VerificationBuilder: verificationBuilder,
	})
}
