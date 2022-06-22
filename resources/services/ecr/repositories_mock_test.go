package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	r := types.Repository{}
	err := faker.FakeData(&r)
	if err != nil {
		t.Fatal(err)
	}
	i := types.ImageDetail{}
	err = faker.FakeData(&i)
	if err != nil {
		t.Fatal(err)
	}
	i.RepositoryName = r.RepositoryName
	maxResults := int32(1000)

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRepositoriesOutput{
			Repositories: []types.Repository{r},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	var f types.ImageScanFindings
	if err := faker.FakeData(&f); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeImageScanFindings(
		gomock.Any(),
		&ecr.DescribeImageScanFindingsInput{
			ImageId:        &types.ImageIdentifier{ImageDigest: i.ImageDigest},
			RepositoryName: r.RepositoryName,
			MaxResults:     &maxResults,
		},
		gomock.Any(),
	).Return(
		&ecr.DescribeImageScanFindingsOutput{ImageScanFindings: &f},
		nil,
	)

	return client.Services{
		ECR: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrRepositoriesMock, client.TestOptions{})
}
