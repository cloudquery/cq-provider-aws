package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildLambdaFunctionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	f := lambda.GetFunctionOutput{}
	err := faker.FakeData(&f)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&f, nil)

	fc := types.FunctionConfiguration{}
	err = faker.FakeData(&fc)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListFunctions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListFunctionsOutput{
			Functions: []types.FunctionConfiguration{fc},
		}, nil)

	a := types.AliasConfiguration{}
	err = faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListAliasesOutput{
			Aliases: []types.AliasConfiguration{a},
		}, nil)

	i := types.FunctionEventInvokeConfig{}
	err = faker.FakeData(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListFunctionEventInvokeConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListFunctionEventInvokeConfigsOutput{
			FunctionEventInvokeConfigs: []types.FunctionEventInvokeConfig{i},
		}, nil)

	cc := types.ProvisionedConcurrencyConfigListItem{}
	err = faker.FakeData(&cc)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListProvisionedConcurrencyConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListProvisionedConcurrencyConfigsOutput{
			ProvisionedConcurrencyConfigs: []types.ProvisionedConcurrencyConfigListItem{cc},
		}, nil)

	esm := types.EventSourceMappingConfiguration{}
	err = faker.FakeData(&esm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListEventSourceMappings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListEventSourceMappingsOutput{
			EventSourceMappings: []types.EventSourceMappingConfiguration{esm},
		}, nil)

	fp := lambda.GetPolicyOutput{}
	err = faker.FakeData(&fp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\":1}"
	fp.Policy = &document
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&fp, nil)

	csco := lambda.GetFunctionCodeSigningConfigOutput{}
	err = faker.FakeData(&csco)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetFunctionCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&csco, nil)

	csc := types.CodeSigningConfig{}
	err = faker.FakeData(&csc)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.GetCodeSigningConfigOutput{
			CodeSigningConfig: &csc,
		}, nil)

	m.EXPECT().ListVersionsByFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListVersionsByFunctionOutput{
			Versions: []types.FunctionConfiguration{fc},
		}, nil)

	return client.Services{
		Lambda: m,
	}
}

func TestLambdaFunctions(t *testing.T) {
	awsTestHelper(t, LambdaFunctions(), buildLambdaFunctionsMock)
}
