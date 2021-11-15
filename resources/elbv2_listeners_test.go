package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElbv2Listeners(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.Listener{}
	if err := faker.FakeData(&l); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeListeners(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeListenersOutput{
			Listeners: []elbv2Types.Listener{l},
		}, nil)

	c := elbv2Types.Certificate{}
	if err := faker.FakeData(&c); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeListenerCertificates(
		gomock.Any(),
		&elasticloadbalancingv2.DescribeListenerCertificatesInput{ListenerArn: l.ListenerArn},
		gomock.Any(),
	).Return(&elasticloadbalancingv2.DescribeListenerCertificatesOutput{
		Certificates: []elbv2Types.Certificate{c},
	}, nil)

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)
	return client.Services{
		ELBv2: m,
	}
}

func TestElbv2Listeners(t *testing.T) {
	awsTestHelper(t, Elbv2Listeners(), buildElbv2Listeners, TestOptions{})
}
