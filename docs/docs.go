package main

import (
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-aws/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

/*
	{
		type: <resource | table >,
		name: name,
		parent: <name>
		relation: <name>
	}
*/

type relationObj struct {
	Obj_type string `json:"type"`
	Name     string `json:"name"`
	Parent   string `json:"parent"`
}

func findRelations(resource string, tables []*schema.Table, foundRelations []relationObj) []relationObj {
	for _, relation := range tables {
		foundRelations = findRelations(relation.Name, relation.Relations, foundRelations)
		foundRelations = append(foundRelations, relationObj{
			Obj_type: "table",
			Name:     relation.Name,
			Parent:   resource,
		})
	}

	return foundRelations

}

func main() {
	// outputPath := "./docs"
	// if err := docs.GenerateDocs(provider.Provider(), outputPath, true); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	// }

	allRelations := make(map[string][]relationObj)

	for resourceName, resource := range provider.Provider().ResourceMap {

		// resource.
		// Table Name
		foundRelations := make([]relationObj, 0)
		foundRelations = append(foundRelations, relationObj{
			Name:   resource.Name,
			Parent: "",
		})
		allRelations[resourceName] = findRelations(resource.Name, resource.Relations, foundRelations)

	}
	// fmt.Println(allRelations)
	relationsJson, _ := json.Marshal(allRelations)
	fmt.Println(string(relationsJson))

}
