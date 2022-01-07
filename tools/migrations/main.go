package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-aws/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/migrations"
)

func main() {
	if err := migrations.Run(context.Background(), provider.Provider(), ""); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
