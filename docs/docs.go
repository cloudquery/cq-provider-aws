package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	outputPath := "./docs"
	dir, err := ioutil.ReadDir(outputPath + "/tables")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
		os.Exit(1)
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{outputPath + "/tables", d.Name()}...))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
			os.Exit(1)
		}

	}

	err = docs.GenerateDocs(resources.Provider(), outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
