package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2Eips(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	a := ec2Types.Address{}
	err := faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeAddresses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeAddressesOutput{
			Addresses: []ec2Types.Address{a},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2Eips(t *testing.T) {
	awsTestHelper(t, Ec2Eips(), buildEc2Eips, TestOptions{})
}
