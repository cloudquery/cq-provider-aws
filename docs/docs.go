package main

import (
	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	docs.GenerateDocs(resources.Provider(), "./docs")
}
