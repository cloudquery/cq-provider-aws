package resourcegroups

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildResourceGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockResourceGroupsClient(ctrl)
	gId := types.GroupIdentifier{}
	err := faker.FakeData(&gId)
	if err != nil {
		t.Fatal(err)
	}

	tagsResponse := resourcegroups.GetTagsOutput{}
	err = faker.FakeData(&tagsResponse)
	if err != nil {
		t.Fatal(err)
	}

	query := types.GroupQuery{}
	err = faker.FakeData(&query)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.ListGroupsOutput{
			GroupIdentifiers: []types.GroupIdentifier{gId},
		}, nil)
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsResponse, nil)
	m.EXPECT().GetGroupQuery(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.GetGroupQueryOutput{
			GroupQuery: &query,
		}, nil)

	return client.Services{
		ResourceGroups: m,
	}
}

func TestResourceGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceGroups(), buildResourceGroupsMock, client.TestOptions{})
}
