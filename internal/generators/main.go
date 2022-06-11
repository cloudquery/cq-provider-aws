package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"gopkg.in/yaml.v3"
)

func ptrToStr(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// ExtractServiceFromCfType extracts the service name from a CloudFormation type name.
// For example, "AWS::EC2::Instance" returns "ec2".
func ExtractServiceFromCfType(cfType string) string {
	return strings.ToLower(strings.Split(cfType, "::")[1])
}

// ExtractServiceFromCfType extracts the service name from a CloudFormation type name.
// For example, "AWS::EC2::Instance" returns "ec2".
func ExtractResourceFromCfType(cfType string) string {
	return strings.ToLower(strings.Split(cfType, "::")[2])
}

// typeNameToTableName turns a CloudFormation resource type name into a CloudQuery table name.
// For example, "AWS::APIGateway::ApiKey" becomes "aws_apigateway_apikey".
func typeNameToTableName(typeName string) string {
	s := strings.Replace(typeName, "::", "_", -1)
	return strings.ToLower(s)
}

type ResourceDefinition struct {
	PackageName string
	Name        string
	RootTable   *TableDefinition
}

type TableDefinition struct {
	Name          string
	TypeName      string
	FileName      string
	TableFuncName string
	TableName     string
	Description   string
	Columns       []ColumnDefinition
	Relations     []*TableDefinition
	// schema.TableResolver definition
	Resolver string
	// Table extra functions
	IgnoreErrorFunc      string
	MultiplexFunc        string
	DeleteFilterFunc     string
	PostResourceResolver string

	// Table Creation Options
	// Options *config.TableOptionsConfig
}

type ColumnDefinition struct {
	Name        string
	Description string
	Type        string
	Resolver    string
}

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

//go:embed resource.go.tpl
var tableTemplate string

//go:embed provider.go.tpl
var providerTemplate string

//go:embed resources.yaml
var resources []byte

type CloudQueryResource struct {
	// Name of the resource is the name of the cloudformation file name
	// e.g. aws_s3_buckets as seed under schemas/*
	Name         string
	ServiceName  string
	ResourceName string
	//Cftype is the CloudFormation Type
	// Cftype string
}

// AWSResourceToGenerate contains which resources to generate
type AWSResourceToGenerate struct {
	Resources []CloudQueryResource
	Services  []string
}

func main() {
	var resourcesToGenerate AWSResourceToGenerate
	if err := yaml.Unmarshal(resources, &resourcesToGenerate); err != nil {
		log.Fatalf("error unmarshalling resources: %v", err)
	}
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}

	for i := range resourcesToGenerate.Resources {
		n := resourcesToGenerate.Resources[i].Name
		resourcesToGenerate.Resources[i].ServiceName = strings.Split(n, "_")[1]
		resourcesToGenerate.Resources[i].ResourceName = strings.Split(n, "_")[2]
	}

	services := make(map[string]bool)
	for _, resource := range resourcesToGenerate.Resources {
		services[resource.ServiceName] = true
	}

	for service := range services {
		resourcesToGenerate.Services = append(resourcesToGenerate.Services, service)
	}

	for _, resource := range resourcesToGenerate.Resources {
		cfResource, err := GenerateResource(resource.Name)
		if err != nil {
			log.Fatal(err)
		}
		serviceName := ExtractServiceFromCfType(*cfResource.TypeName)
		resourceName := ExtractResourceFromCfType(*cfResource.TypeName)
		// fmt.Printf("%+v\n", cfResource)
		// fmt.Println(cfResource)
		tableDefition, err := ConvertCFtoCQ(cfResource)
		if err != nil {
			log.Fatal(err)
		}
		resourceDefinition := ResourceDefinition{
			PackageName: serviceName,
			Name:        resource.Name,
			RootTable:   tableDefition,
		}
		t, err := template.New("").Funcs(funcMap).Parse(tableTemplate)
		if err != nil {
			log.Fatal(err)
		}
		goPath := fmt.Sprintf("./resources/services/%s/%s.go", serviceName, resourceName)
		os.MkdirAll(path.Dir(goPath), os.ModePerm)
		f, err := os.Create(goPath)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = t.Execute(f, resourceDefinition)
		if err != nil {
			f.Close()
			log.Fatal(err)
		}
		f.Close()

	}

	t, err := template.New("").Parse(providerTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f1, err := os.Create("./resources/provider/provider.go")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f1, resourcesToGenerate)
	if err != nil {
		log.Fatal(err)
	}
	f1.Close()
}

// GenerateResource gets the CloudFormation resource name
// reads the schema from the AWS SDK and returns unmarshalled Resource struct
func GenerateResource(name string) (*cfschema.Resource, error) {
	name = strings.Replace(name, "_", "-", -1)
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(fmt.Sprintf("./internal/generators/schemas/%s.json", name))
	if err != nil {
		return nil, fmt.Errorf("reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()
	if err != nil {
		return nil, fmt.Errorf("parsing CloudFormation Resource Type Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("expanding JSON Pointer references: %w", err)
	}

	return resource, nil
}

// See ful cloudformation resource schema defintion https://github.com/aws-cloudformation/cloudformation-cli/blob/master/src/rpdk/core/data/schema/provider.definition.schema.v1.json
var CfTypeToCQType = map[string]string{
	"string":  "schema.TypeString",
	"boolean": "schema.TypeBool",
	"number":  "schema.TypeFloat",
	"integer": "schema.TypeFloat",
	"array":   "schema.TypeJSON",
	"object":  "schema.TypeJSON",
}

func ConvertCFtoCQ(cfResource *cfschema.Resource) (*TableDefinition, error) {
	tableDefinition := TableDefinition{
		Name:            typeNameToTableName(*cfResource.TypeName),
		TypeName:        *cfResource.TypeName,
		Description:     ptrToStr(cfResource.Description),
		Resolver:        "fetch" + typeNameToTableName(*cfResource.TypeName),
		MultiplexFunc:   "client.AccountMultiplexer",
		IgnoreErrorFunc: "client.IgnoreAccessDeniedServiceDisabled",
	}
	for k, v := range cfResource.Properties {
		CQType := CfTypeToCQType[v.Type.String()]
		if CQType == "" {
			log.Fatal("unsupported type: ", v.Type.String())
		}
		// v.Type
		tableDefinition.Columns = append(tableDefinition.Columns,
			ColumnDefinition{
				Name:        k,
				Description: ptrToStr(v.Description),
				Type:        CfTypeToCQType[v.Type.String()],
			})

	}
	return &tableDefinition, nil
}
