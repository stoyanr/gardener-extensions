// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extensions/controllers/provider-alicloud/pkg/alicloud/client (interfaces: VPC,Factory,ClientFactory,ECS,STS)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	vpc "github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	client "github.com/gardener/gardener-extensions/controllers/provider-alicloud/pkg/alicloud/client"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVPC is a mock of VPC interface
type MockVPC struct {
	ctrl     *gomock.Controller
	recorder *MockVPCMockRecorder
}

// MockVPCMockRecorder is the mock recorder for MockVPC
type MockVPCMockRecorder struct {
	mock *MockVPC
}

// NewMockVPC creates a new mock instance
func NewMockVPC(ctrl *gomock.Controller) *MockVPC {
	mock := &MockVPC{ctrl: ctrl}
	mock.recorder = &MockVPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVPC) EXPECT() *MockVPCMockRecorder {
	return m.recorder
}

// DescribeEipAddresses mocks base method
func (m *MockVPC) DescribeEipAddresses(arg0 *vpc.DescribeEipAddressesRequest) (*vpc.DescribeEipAddressesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEipAddresses", arg0)
	ret0, _ := ret[0].(*vpc.DescribeEipAddressesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEipAddresses indicates an expected call of DescribeEipAddresses
func (mr *MockVPCMockRecorder) DescribeEipAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEipAddresses", reflect.TypeOf((*MockVPC)(nil).DescribeEipAddresses), arg0)
}

// DescribeNatGateways mocks base method
func (m *MockVPC) DescribeNatGateways(arg0 *vpc.DescribeNatGatewaysRequest) (*vpc.DescribeNatGatewaysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNatGateways", arg0)
	ret0, _ := ret[0].(*vpc.DescribeNatGatewaysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNatGateways indicates an expected call of DescribeNatGateways
func (mr *MockVPCMockRecorder) DescribeNatGateways(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNatGateways", reflect.TypeOf((*MockVPC)(nil).DescribeNatGateways), arg0)
}

// DescribeVpcs mocks base method
func (m *MockVPC) DescribeVpcs(arg0 *vpc.DescribeVpcsRequest) (*vpc.DescribeVpcsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcs", arg0)
	ret0, _ := ret[0].(*vpc.DescribeVpcsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcs indicates an expected call of DescribeVpcs
func (mr *MockVPCMockRecorder) DescribeVpcs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockVPC)(nil).DescribeVpcs), arg0)
}

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewVPC mocks base method
func (m *MockFactory) NewVPC(arg0, arg1, arg2 string) (client.VPC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVPC", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.VPC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVPC indicates an expected call of NewVPC
func (mr *MockFactoryMockRecorder) NewVPC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVPC", reflect.TypeOf((*MockFactory)(nil).NewVPC), arg0, arg1, arg2)
}

// MockClientFactory is a mock of ClientFactory interface
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// NewECSClient mocks base method
func (m *MockClientFactory) NewECSClient(arg0 context.Context, arg1, arg2, arg3 string) (client.ECS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewECSClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(client.ECS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewECSClient indicates an expected call of NewECSClient
func (mr *MockClientFactoryMockRecorder) NewECSClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewECSClient", reflect.TypeOf((*MockClientFactory)(nil).NewECSClient), arg0, arg1, arg2, arg3)
}

// NewSTSClient mocks base method
func (m *MockClientFactory) NewSTSClient(arg0 context.Context, arg1, arg2, arg3 string) (client.STS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSTSClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(client.STS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSTSClient indicates an expected call of NewSTSClient
func (mr *MockClientFactoryMockRecorder) NewSTSClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSTSClient", reflect.TypeOf((*MockClientFactory)(nil).NewSTSClient), arg0, arg1, arg2, arg3)
}

// MockECS is a mock of ECS interface
type MockECS struct {
	ctrl     *gomock.Controller
	recorder *MockECSMockRecorder
}

// MockECSMockRecorder is the mock recorder for MockECS
type MockECSMockRecorder struct {
	mock *MockECS
}

// NewMockECS creates a new mock instance
func NewMockECS(ctrl *gomock.Controller) *MockECS {
	mock := &MockECS{ctrl: ctrl}
	mock.recorder = &MockECSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECS) EXPECT() *MockECSMockRecorder {
	return m.recorder
}

// CheckIfImageExists mocks base method
func (m *MockECS) CheckIfImageExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfImageExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfImageExists indicates an expected call of CheckIfImageExists
func (mr *MockECSMockRecorder) CheckIfImageExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfImageExists", reflect.TypeOf((*MockECS)(nil).CheckIfImageExists), arg0, arg1)
}

// ShareImageToAccount mocks base method
func (m *MockECS) ShareImageToAccount(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShareImageToAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShareImageToAccount indicates an expected call of ShareImageToAccount
func (mr *MockECSMockRecorder) ShareImageToAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShareImageToAccount", reflect.TypeOf((*MockECS)(nil).ShareImageToAccount), arg0, arg1, arg2, arg3)
}

// MockSTS is a mock of STS interface
type MockSTS struct {
	ctrl     *gomock.Controller
	recorder *MockSTSMockRecorder
}

// MockSTSMockRecorder is the mock recorder for MockSTS
type MockSTSMockRecorder struct {
	mock *MockSTS
}

// NewMockSTS creates a new mock instance
func NewMockSTS(ctrl *gomock.Controller) *MockSTS {
	mock := &MockSTS{ctrl: ctrl}
	mock.recorder = &MockSTSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSTS) EXPECT() *MockSTSMockRecorder {
	return m.recorder
}

// GetAccountIDFromCallerIdentity mocks base method
func (m *MockSTS) GetAccountIDFromCallerIdentity(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountIDFromCallerIdentity", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountIDFromCallerIdentity indicates an expected call of GetAccountIDFromCallerIdentity
func (mr *MockSTSMockRecorder) GetAccountIDFromCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountIDFromCallerIdentity", reflect.TypeOf((*MockSTS)(nil).GetAccountIDFromCallerIdentity), arg0)
}
