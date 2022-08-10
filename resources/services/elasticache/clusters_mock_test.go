package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticacheClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElastiCache(ctrl)
	output := elasticache.DescribeCacheClustersOutput{}
	err := faker.FakeData(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		ElastiCache: mockElasticache,
	}
}

func TestEksClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildElasticacheClusters, client.TestOptions{})
}
