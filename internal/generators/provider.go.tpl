package provider

import (

	"github.com/cloudquery/cq-provider-aws/client"
  {{range $service := .Services}}
  "github.com/cloudquery/cq-provider-aws/resources/services/{{$service}}"
  {{- end }}
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (

	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:             "aws",
		Version:          Version,
		Configure:        client.Configure,
//		ErrorClassifier:  client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
      {{range $resource := .Resources}}
      "{{$resource.ServiceName}}.{{$resource.ResourceName}}": {{$resource.ServiceName}}.Aws_{{$resource.ServiceName}}_{{$resource.ResourceName}}(),
      {{- end }}
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
