package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2RegionalConfig(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	config := ec2RegionalConfig{}
	if err := faker.FakeData(&config); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetEbsDefaultKmsKeyId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsDefaultKmsKeyIdOutput{KmsKeyId: aws.String("some/key/id")}, nil).AnyTimes()
	m.EXPECT().GetEbsEncryptionByDefault(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsEncryptionByDefaultOutput{EbsEncryptionByDefault: true}, nil).AnyTimes()

	return client.Services{
		EC2: m,
	}
}

func TestEc2RegionalConfig(t *testing.T) {
	awsTestHelper(t, Ec2RegionalConfig(), buildEc2RegionalConfig)
}
