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

func buildMqConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMQClient(ctrl)

	c := types.Configuration{}
	err := faker.FakeData(&c)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mq.ListConfigurationsOutput{
			Configurations: []types.Configuration{c},
		}, nil)

	return client.Services{
		MQ: m,
	}
}

func TestMqConfigurations(t *testing.T) {
	awsTestHelper(t, MqConfigurations(), buildMqConfigurations, TestOptions{})
}
