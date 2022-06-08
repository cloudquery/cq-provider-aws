package cloudtrail

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailTypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudtrailTrailsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	trailList, err := faker.FakeDataNullablePermutations(cloudtrailTypes.Trail{})
	if err != nil {
		t.Fatal(err)
	}

	describeTrailOutput := cloudtrail.DescribeTrailsOutput{
		TrailList: trailList.([]cloudtrailTypes.Trail),
	}

	for i := range describeTrailOutput.TrailList {
		describeTrailOutput.TrailList[i].TrailARN = aws.String(fmt.Sprintf("arn:aws:cloudtrail:eu-central-1:testAccount:trail/test-%d", i))
	}

	trailStatus := cloudtrail.GetTrailStatusOutput{}
	err = faker.FakeData(&trailStatus)
	if err != nil {
		t.Fatal(err)
	}
	eventSelector := cloudtrailTypes.EventSelector{}
	err = faker.FakeData(&eventSelector)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&describeTrailOutput,
		nil,
	)
	m.EXPECT().GetTrailStatus(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&trailStatus,
		nil,
	)
	m.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []cloudtrailTypes.EventSelector{eventSelector},
		},
		nil,
	)
	tags := cloudtrail.ListTagsOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.ResourceTagList[0].ResourceId = aws.String(fmt.Sprintf("arn:aws:cloudtrail:eu-central-1:testAccount:trail/test-%d", 1))
	tags.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&tags, nil)

	return services
}

func TestCloudtrailTrails(t *testing.T) {
	client.AwsMockTestHelper(t, CloudtrailTrails(), buildCloudtrailTrailsMock, client.TestOptions{})
}
