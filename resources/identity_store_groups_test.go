package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIdentityStoreGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIdentityStoreClient(ctrl)

	u := identitystore.ListGroupsOutput{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&u, nil)

	u.NextToken = nil

	return client.Services{
		IdentityStore: m,
	}
}

func TestIdentityStoreGroups(t *testing.T) {
	awsTestHelper(t, IdentityStoreGroups(), buildIdentityStoreGroups, TestOptions{})
}
