package cloudformation

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildStacks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudFormationClient(ctrl)

	if err := faker.SetRandomMapAndSliceMinSize(1); err != nil {
		t.Fatal(err)
	}
	if err := faker.SetRandomMapAndSliceMaxSize(1); err != nil {
		t.Fatal(err)
	}

	var stack types.Stack
	stackList, err := faker.FakeDataNullablePermutations(stack)

	if err != nil {
		t.Fatal(err)
	}
	o := cloudformation.DescribeStacksOutput{Stacks: stackList.([]types.Stack)}
	for i := range o.Stacks {
		s := "somearn" + fmt.Sprintf("%d", i)
		o.Stacks[i].StackId = &s
	}
	mock.EXPECT().DescribeStacks(
		gomock.Any(),
		&cloudformation.DescribeStacksInput{},
		gomock.Any(),
	).Return(
		&o,
		nil,
	)

	var resource types.StackResourceSummary
	if err := faker.FakeData(&resource); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListStackResources(
		gomock.Any(),
		&cloudformation.ListStackResourcesInput{StackName: stack.StackName},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackResourcesOutput{StackResourceSummaries: []types.StackResourceSummary{resource}},
		nil,
	)

	return client.Services{Cloudformation: mock}
}

func TestCloudformationStacks(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildStacks, client.TestOptions{})
}
