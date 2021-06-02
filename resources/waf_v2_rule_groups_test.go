package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWAFV2RuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafV2Client(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	if err := faker.FakeData(&tempRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListRuleGroupsOutput{
		RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
	}, nil)

	return client.Services{WafV2: m}
}

func TestWafV2RuleGroups(t *testing.T) {
	awsTestHelper(t, WafV2RuleGroups(), buildWAFV2RuleGroupsMock, TestOptions{})
}
