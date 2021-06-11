package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMqUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMQClient(ctrl)

	us := types.UserSummary{}
	err := faker.FakeData(&us)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mq.ListUsersOutput{
			Users: []types.UserSummary{us},
		}, nil)

	return client.Services{
		MQ: m,
	}
}

func TestMqUsers(t *testing.T) {
	awsTestHelper(t, MqUsers(), buildMqUsers, TestOptions{})
}
