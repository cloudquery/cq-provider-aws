package resources

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/gocarina/gocsv"
	"github.com/golang/mock/gomock"
)

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.User{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}
	g := iamTypes.Group{}
	err = faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}
	km := iamTypes.AccessKeyMetadata{}
	err = faker.FakeData(&km)
	if err != nil {
		t.Fatal(err)
	}
	aup := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&aup)
	if err != nil {
		t.Fatal(err)
	}
	akl := iam.GetAccessKeyLastUsedOutput{}
	err = faker.FakeData(&akl)
	if err != nil {
		t.Fatal(err)
	}

	ru := reportUser{}
	err = faker.FakeData(&ru)
	if err != nil {
		t.Fatal(err)
	}
	ru.ARN = aws.ToString(u.Arn)
	ru.PasswordNextRotation = time.Now().Format(time.RFC3339)
	ru.PasswordLastChanged = time.Now().Format(time.RFC3339)
	content, err := gocsv.MarshalBytes([]reportUser{ru})
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return(
		&iam.ListUsersOutput{
			Users: []iamTypes.User{u},
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsForUserOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any()).Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []iamTypes.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any()).Return(
		&akl, nil)

	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any()).Return(
		&iam.GetCredentialReportOutput{
			Content: content,
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func TestIamUsers(t *testing.T) {
	awsTestHelper(t, IamUsers(), buildIamUsers)
}
