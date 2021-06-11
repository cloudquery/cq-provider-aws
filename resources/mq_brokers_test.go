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

func buildMqBrokers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMQClient(ctrl)

	bs := types.BrokerSummary{}
	err := faker.FakeData(&bs)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBrokers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mq.ListBrokersOutput{
			BrokerSummaries: []types.BrokerSummary{bs},
		}, nil)

	return client.Services{
		MQ: m,
	}
}

func TestMqBrokers(t *testing.T) {
	awsTestHelper(t, MqBrokers(), buildMqBrokers, TestOptions{})
}
