package cloudfront

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudfrontDistributionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	ds := cloudfrontTypes.DistributionSummary{}
	dsList, err := faker.FakeDataNullablePermutations(ds)
	if err != nil {
		t.Fatal(err)
	}
	cloudfrontOutput := &cloudfront.ListDistributionsOutput{
		DistributionList: &cloudfrontTypes.DistributionList{
			Items: dsList.([]cloudfrontTypes.DistributionSummary),
		},
	}
	for i := range cloudfrontOutput.DistributionList.Items {
		s := "somearn" + fmt.Sprintf("%d", i)
		cloudfrontOutput.DistributionList.Items[i].ARN = &s
	}
	m.EXPECT().ListDistributions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)

	distribution := &cloudfront.GetDistributionOutput{}
	if err := faker.FakeData(&distribution); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetDistribution(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		distribution,
		nil,
	)

	tags := &cloudfront.ListTagsForResourceOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)
	return services
}

func TestCloudfrontDistributions(t *testing.T) {
	client.AwsMockTestHelper(t, CloudfrontDistributions(), buildCloudfrontDistributionsMock, client.TestOptions{})
}
