package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWAFWebACLMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempWebACLSum := types.WebACLSummary{}
	if err := faker.FakeData(&tempWebACLSum); err != nil {
		t.Fatal(err)
	}
	tempWebACL := types.WebACL{}
	if err := faker.FakeData(&tempWebACL); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListWebACLs(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListWebACLsOutput{
		WebACLs: []types.WebACLSummary{tempWebACLSum},
	}, nil)
	m.EXPECT().GetWebACL(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.GetWebACLOutput{
		WebACL: &tempWebACL,
	}, nil)

	return client.Services{Waf: m}
}

func TestWafWebACL(t *testing.T) {
	awsTestHelper(t, WafWebacls(), buildWAFWebACLMock)
}
