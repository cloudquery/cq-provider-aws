// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: DirectconnectClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	directconnect "github.com/aws/aws-sdk-go-v2/service/directconnect"
	gomock "github.com/golang/mock/gomock"
)

// MockDirectconnectClient is a mock of DirectconnectClient interface.
type MockDirectconnectClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectconnectClientMockRecorder
}

// MockDirectconnectClientMockRecorder is the mock recorder for MockDirectconnectClient.
type MockDirectconnectClientMockRecorder struct {
	mock *MockDirectconnectClient
}

// NewMockDirectconnectClient creates a new mock instance.
func NewMockDirectconnectClient(ctrl *gomock.Controller) *MockDirectconnectClient {
	mock := &MockDirectconnectClient{ctrl: ctrl}
	mock.recorder = &MockDirectconnectClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectconnectClient) EXPECT() *MockDirectconnectClientMockRecorder {
	return m.recorder
}

// DescribeDirectConnectGateways mocks base method.
func (m *MockDirectconnectClient) DescribeDirectConnectGateways(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectConnectGateways indicates an expected call of DescribeDirectConnectGateways.
func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGateways), varargs...)
}

// DescribeVirtualGateways mocks base method.
func (m *MockDirectconnectClient) DescribeVirtualGateways(arg0 context.Context, arg1 *directconnect.DescribeVirtualGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualGateways indicates an expected call of DescribeVirtualGateways.
func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualGateways), varargs...)
}

// DescribeVirtualInterfaces mocks base method.
func (m *MockDirectconnectClient) DescribeVirtualInterfaces(arg0 context.Context, arg1 *directconnect.DescribeVirtualInterfacesInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualInterfaces", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualInterfaces indicates an expected call of DescribeVirtualInterfaces.
func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualInterfaces", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualInterfaces), varargs...)
}
