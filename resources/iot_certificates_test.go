package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	certs := iot.ListCertificatesOutput{}
	err := faker.FakeData(&certs)
	if err != nil {
		t.Fatal(err)
	}
	certs.NextMarker = nil
	m.EXPECT().ListCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&certs, nil)

	cd := iot.DescribeCertificateOutput{}
	err = faker.FakeData(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	return client.Services{
		IOT: m,
	}
}

func TestIotCertificates(t *testing.T) {
	awsTestHelper(t, IotCertificates(), buildIotCertificatesMock, TestOptions{})
}
