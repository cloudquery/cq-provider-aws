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

func buildWAFRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleSum := types.RuleSummary{}
	if err := faker.FakeData(&tempRuleSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListRulesOutput{
		Rules: []types.RuleSummary{tempRuleSum},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafRules(t *testing.T) {
	awsTestHelper(t, WafRules(), buildWAFRulesMock, TestOptions{})
}
