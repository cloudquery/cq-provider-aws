package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildFSXClientForFileSystems(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockFsxClient(ctrl)

	var fs types.FileSystem
	if err := faker.FakeData(&fs); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeFileSystems(
		gomock.Any(),
		&fsx.DescribeFileSystemsInput{},
		gomock.Any(),
	).Return(
		&fsx.DescribeFileSystemsOutput{FileSystems: []types.FileSystem{fs}},
		nil,
	)

	var aliases []types.Alias
	if err := faker.FakeData(&aliases); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeFileSystemAliases(
		gomock.Any(),
		&fsx.DescribeFileSystemAliasesInput{FileSystemId: fs.FileSystemId},
		gomock.Any(),
	).Return(
		&fsx.DescribeFileSystemAliasesOutput{
			Aliases: aliases,
		},
		nil,
	)

	return client.Services{FSX: mock}
}

func TestFSXFileSystems(t *testing.T) {
	client.AwsMockTestHelper(t, FsxFileSystems(), buildFSXClientForFileSystems, client.TestOptions{})
}
