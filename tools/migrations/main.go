package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-aws/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/migrations"
	"github.com/hashicorp/go-hclog"
)

func main() {
	const outputPath = "./resources/provider/migrations"

	if err := migrations.Generate(context.Background(), hclog.L(), provider.Provider(), outputPath, ""); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate migrations: %s\n", err)
	}
}
