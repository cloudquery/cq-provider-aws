package provider

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
)

func TestLoad(t *testing.T) {
	cfg := `
		max_retries = 20
  		max_backoff = 60
  	`
	// we want to give a way for the user ci to override with additional cfg for integration test.
	// For example:
	//
	// accounts "cq-provider-aws" {
	// 	role_arn = "arn:aws:iam::70xxxxxxxxxx:role/CqProviderxxxxxxxxxxxxxxxxx"
	// }
	// accounts "cq-dev" {
	//  role_arn = "arn:aws:iam::70xxxxxxxxxx:role/CqProviderxxxxxxxxxxxxxxxxx"
	// }

	additionalConfig := os.Getenv("CQ_TEST_CFG")
	if additionalConfig != "" {
		cfg += additionalConfig
	}

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider:              LoadTestProvider(),
		Config:                cfg,
		NotParallel:           true,
		ParallelFetchingLimit: 10000,
	})
}

func LoatTestConfigure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, error) {
	return &LoadClient{
		logger: logger,
	}, nil
}

func LoadTestProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "load",
		Version:   Version,
		Configure: LoatTestConfigure,
		// ErrorClassifier:  client.ErrorClassifier,
		// Migrations:       awsMigrations,
		// ModuleInfoReader: module.EmbeddedReader(moduleData, "moduledata"),
		ResourceMap: map[string]*schema.Table{
			"load.test": LoadTestTable(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}

type LoadClient struct {
	logger hclog.Logger
}

func (c *LoadClient) Logger() hclog.Logger {
	return c.logger
}

func LoadTestMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	clients := make([]schema.ClientMeta, 1000)
	for i := range clients {
		clients[i] = meta
	}
	return clients
}

func LoadTestTable() *schema.Table {
	return &schema.Table{
		Name:          "aws_load_test",
		Description:   "LoadTestTable",
		Resolver:      fetchLoadTestResource,
		Multiplex:     LoadTestMultiplex,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_authority_arn",
				Description: "The Amazon Resource Name (ARN) of the ACM PCA private certificate authority (CA) that issued the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The source of the certificate",
				Type:        schema.TypeString,
			},
		},
	}
}

type LoadTest struct {
	Arn                     string `json:"arn"`
	CertificateAuthorityArn string `json:"certificate_arn"`
	Type                    string `json:"type"`
}

func fetchLoadTestResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	generatedResult := make([]LoadTest, 10000)
	for i := range generatedResult {
		generatedResult[i].Arn = uuid.New().String()
		generatedResult[i].CertificateAuthorityArn = uuid.New().String()
		generatedResult[i].Type = "random type"
	}
	res <- generatedResult
	return nil
}
