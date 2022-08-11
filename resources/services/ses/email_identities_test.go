package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildEmailIdentities(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSESClient(ctrl)

	li := sesv2.ListEmailIdentitiesOutput{}
	err := faker.FakeData(&li)
	if err != nil {
		t.Fatal(err)
	}
	li.NextToken = nil
	sesClient.EXPECT().ListEmailIdentities(gomock.Any(), gomock.Any()).Return(&li, nil)

	gi := sesv2.GetEmailIdentityOutput{}
	err = faker.FakeData(&gi)
	if err != nil {
		t.Fatal(err)
	}
	sesClient.EXPECT().GetEmailIdentity(gomock.Any(), gomock.Any()).Return(&gi, nil)

	return client.Services{
		SES: sesClient,
	}
}

func TestEmailIdentities(t *testing.T) {
	client.AwsMockTestHelper(t, EmailIdentities(), buildEmailIdentities, client.TestOptions{})
}
