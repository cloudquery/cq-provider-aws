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

func buildWAFRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	if err := faker.FakeData(&tempRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&waf.ListRuleGroupsOutput{
		RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
	}, nil)

	return client.Services{Waf: m}
}

func TestWafRuleGroups(t *testing.T) {
	awsTestHelper(t, WafRulegroups(), buildWAFRuleGroupsMock, TestOptions{})
}
