// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: DynamoDBClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

// MockDynamoDBClient is a mock of DynamoDBClient interface.
type MockDynamoDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDynamoDBClientMockRecorder
}

// MockDynamoDBClientMockRecorder is the mock recorder for MockDynamoDBClient.
type MockDynamoDBClientMockRecorder struct {
	mock *MockDynamoDBClient
}

// NewMockDynamoDBClient creates a new mock instance.
func NewMockDynamoDBClient(ctrl *gomock.Controller) *MockDynamoDBClient {
	mock := &MockDynamoDBClient{ctrl: ctrl}
	mock.recorder = &MockDynamoDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDynamoDBClient) EXPECT() *MockDynamoDBClientMockRecorder {
	return m.recorder
}

// DescribeTable mocks base method.
func (m *MockDynamoDBClient) DescribeTable(arg0 context.Context, arg1 *dynamodb.DescribeTableInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTable", varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTable indicates an expected call of DescribeTable.
func (mr *MockDynamoDBClientMockRecorder) DescribeTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockDynamoDBClient)(nil).DescribeTable), varargs...)
}

// ListTables mocks base method.
func (m *MockDynamoDBClient) ListTables(arg0 context.Context, arg1 *dynamodb.ListTablesInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTables", varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables.
func (mr *MockDynamoDBClientMockRecorder) ListTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockDynamoDBClient)(nil).ListTables), varargs...)
}

// ListTagsOfResource mocks base method.
func (m *MockDynamoDBClient) ListTagsOfResource(arg0 context.Context, arg1 *dynamodb.ListTagsOfResourceInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsOfResource", varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTagsOfResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsOfResource indicates an expected call of ListTagsOfResource.
func (mr *MockDynamoDBClientMockRecorder) ListTagsOfResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsOfResource", reflect.TypeOf((*MockDynamoDBClient)(nil).ListTagsOfResource), varargs...)
}
