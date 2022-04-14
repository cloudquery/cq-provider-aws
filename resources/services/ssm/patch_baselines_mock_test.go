package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildPatchBaselinesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSSMClient(ctrl)

	var baseline ssm.GetPatchBaselineOutput
	if err := faker.FakeData(&baseline); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribePatchBaselines(
		gomock.Any(),
		&ssm.DescribePatchBaselinesInput{},
		gomock.Any(),
	).Return(
		&ssm.DescribePatchBaselinesOutput{
			BaselineIdentities: []types.PatchBaselineIdentity{{BaselineId: baseline.BaselineId}},
		},
		nil,
	)

	mock.EXPECT().GetPatchBaseline(
		gomock.Any(),
		&ssm.GetPatchBaselineInput{BaselineId: baseline.BaselineId},
		gomock.Any(),
	).Return(&baseline, nil)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&ssm.ListTagsForResourceInput{ResourceId: baseline.BaselineId, ResourceType: types.ResourceTypeForTaggingPatchBaseline},
		gomock.Any(),
	).Return(
		&ssm.ListTagsForResourceOutput{TagList: []types.Tag{
			{Key: aws.String("key"), Value: aws.String("value")},
		}},
		nil,
	)

	return client.Services{SSM: mock}
}

func TestPatchBaselines(t *testing.T) {
	client.AwsMockTestHelper(t, PatchBaselines(), buildPatchBaselinesMock, client.TestOptions{})
}
