package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/cloudquery/cq-provider-aws/client"
)

const (
	awsEndpointFile = "https://raw.githubusercontent.com/aws/aws-sdk-go-v2/main/codegen/smithy-aws-go-codegen/src/main/resources/software/amazon/smithy/aws/go/codegen/endpoints.json"
)

type supportedServicesData struct {
	Partitions []struct {
		PartitionId   string `json:"partition"`
		PartitionName string `json:"partitionName"`
		Services      map[string]struct {
			Endpoints map[string]struct {
				Deprecated bool `json:"endpoints" default:"false"`
			} `json:"endpoints"`
		} `json:"services"`
	} `json:"partitions"`
}

func getPartitionRegionServiceData() (*client.SupportedServiceRegionsData, error) {
	// fetch the aws endpoints json file
	req, err := http.NewRequest(http.MethodGet, awsEndpointFile, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get aws supported resources for region, status code: %d", resp.StatusCode)
	}

	var data supportedServicesData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	awsPartitions := make(map[string]client.AwsPartition)
	for _, p := range data.Partitions {
		services := make(map[string]*client.AwsService)
		for sk, s := range p.Services {
			endpoints := make(map[string]*map[string]interface{})
			for ek, e := range s.Endpoints {
				if !e.Deprecated {
					endpoints[ek] = &map[string]interface{}{}
				}
			}
			services[sk] = &client.AwsService{
				Regions: endpoints,
			}
		}

		awsPartitions[p.PartitionId] = client.AwsPartition{
			Id:       p.PartitionId,
			Name:     p.PartitionName,
			Services: services,
		}
	}

	return &client.SupportedServiceRegionsData{
		Partitions: awsPartitions,
	}, nil
}

func saveToJsonFile(data *client.SupportedServiceRegionsData, filePath string) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, b, 0644)
}

func partitionRegionServiceGenerator() error {
	data, err := getPartitionRegionServiceData()
	if err != nil {
		return err
	}

	if err := saveToJsonFile(data, filepath.Join("client", client.PartitionServiceRegionFile)); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := partitionRegionServiceGenerator(); err != nil {
		panic(err)
	}
}