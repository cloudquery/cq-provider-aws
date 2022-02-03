package provider_test

import (
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrationPostgreSQL(t *testing.T) {
	dsn := os.Getenv("CQ_MIGRATION_TEST_PG_DSN")
	if dsn == "" {
		os.Setenv("CQ_MIGRATION_TEST_DSN", "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	} else {
		os.Setenv("CQ_MIGRATION_TEST_DSN", dsn)
	}
	migration.RunMigrationsTest(t, provider.Provider(), []string{"latest"})
}

func TestMigrationTimescaleDB(t *testing.T) {
	dsn := os.Getenv("CQ_MIGRATION_TEST_TSDB_DSN")
	if dsn == "" {
		os.Setenv("CQ_MIGRATION_TEST_DSN", "tsdb://postgres:pass@localhost:5432/postgres?sslmode=disable")
	} else {
		os.Setenv("CQ_MIGRATION_TEST_DSN", dsn)
	}
	migration.RunMigrationsTest(t, provider.Provider(), []string{"latest"})
}
