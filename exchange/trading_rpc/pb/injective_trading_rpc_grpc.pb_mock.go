// Code generated by MockGen. DO NOT EDIT.
// Source: api/gen/grpc/_trading_rpc/pb/_trading_rpc_grpc.pb.go

// Package _trading_rpcpb is a generated GoMock package.
package _trading_rpcpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockHeliosTradingRPCClient is a mock of HeliosTradingRPCClient interface.
type MockHeliosTradingRPCClient struct {
	ctrl     *gomock.Controller
	recorder *MockHeliosTradingRPCClientMockRecorder
}

// MockHeliosTradingRPCClientMockRecorder is the mock recorder for MockHeliosTradingRPCClient.
type MockHeliosTradingRPCClientMockRecorder struct {
	mock *MockHeliosTradingRPCClient
}

// NewMockHeliosTradingRPCClient creates a new mock instance.
func NewMockHeliosTradingRPCClient(ctrl *gomock.Controller) *MockHeliosTradingRPCClient {
	mock := &MockHeliosTradingRPCClient{ctrl: ctrl}
	mock.recorder = &MockHeliosTradingRPCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeliosTradingRPCClient) EXPECT() *MockHeliosTradingRPCClientMockRecorder {
	return m.recorder
}

// ListTradingStrategies mocks base method.
func (m *MockHeliosTradingRPCClient) ListTradingStrategies(ctx context.Context, in *ListTradingStrategiesRequest, opts ...grpc.CallOption) (*ListTradingStrategiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTradingStrategies", varargs...)
	ret0, _ := ret[0].(*ListTradingStrategiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTradingStrategies indicates an expected call of ListTradingStrategies.
func (mr *MockHeliosTradingRPCClientMockRecorder) ListTradingStrategies(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTradingStrategies", reflect.TypeOf((*MockHeliosTradingRPCClient)(nil).ListTradingStrategies), varargs...)
}

// MockHeliosTradingRPCServer is a mock of HeliosTradingRPCServer interface.
type MockHeliosTradingRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockHeliosTradingRPCServerMockRecorder
}

// MockHeliosTradingRPCServerMockRecorder is the mock recorder for MockHeliosTradingRPCServer.
type MockHeliosTradingRPCServerMockRecorder struct {
	mock *MockHeliosTradingRPCServer
}

// NewMockHeliosTradingRPCServer creates a new mock instance.
func NewMockHeliosTradingRPCServer(ctrl *gomock.Controller) *MockHeliosTradingRPCServer {
	mock := &MockHeliosTradingRPCServer{ctrl: ctrl}
	mock.recorder = &MockHeliosTradingRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeliosTradingRPCServer) EXPECT() *MockHeliosTradingRPCServerMockRecorder {
	return m.recorder
}

// ListTradingStrategies mocks base method.
func (m *MockHeliosTradingRPCServer) ListTradingStrategies(arg0 context.Context, arg1 *ListTradingStrategiesRequest) (*ListTradingStrategiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTradingStrategies", arg0, arg1)
	ret0, _ := ret[0].(*ListTradingStrategiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTradingStrategies indicates an expected call of ListTradingStrategies.
func (mr *MockHeliosTradingRPCServerMockRecorder) ListTradingStrategies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTradingStrategies", reflect.TypeOf((*MockHeliosTradingRPCServer)(nil).ListTradingStrategies), arg0, arg1)
}

// mustEmbedUnimplementedHeliosTradingRPCServer mocks base method.
func (m *MockHeliosTradingRPCServer) mustEmbedUnimplementedHeliosTradingRPCServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHeliosTradingRPCServer")
}

// mustEmbedUnimplementedHeliosTradingRPCServer indicates an expected call of mustEmbedUnimplementedHeliosTradingRPCServer.
func (mr *MockHeliosTradingRPCServerMockRecorder) mustEmbedUnimplementedHeliosTradingRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHeliosTradingRPCServer", reflect.TypeOf((*MockHeliosTradingRPCServer)(nil).mustEmbedUnimplementedHeliosTradingRPCServer))
}

// MockUnsafeHeliosTradingRPCServer is a mock of UnsafeHeliosTradingRPCServer interface.
type MockUnsafeHeliosTradingRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeHeliosTradingRPCServerMockRecorder
}

// MockUnsafeHeliosTradingRPCServerMockRecorder is the mock recorder for MockUnsafeHeliosTradingRPCServer.
type MockUnsafeHeliosTradingRPCServerMockRecorder struct {
	mock *MockUnsafeHeliosTradingRPCServer
}

// NewMockUnsafeHeliosTradingRPCServer creates a new mock instance.
func NewMockUnsafeHeliosTradingRPCServer(ctrl *gomock.Controller) *MockUnsafeHeliosTradingRPCServer {
	mock := &MockUnsafeHeliosTradingRPCServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeHeliosTradingRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeHeliosTradingRPCServer) EXPECT() *MockUnsafeHeliosTradingRPCServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedHeliosTradingRPCServer mocks base method.
func (m *MockUnsafeHeliosTradingRPCServer) mustEmbedUnimplementedHeliosTradingRPCServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHeliosTradingRPCServer")
}

// mustEmbedUnimplementedHeliosTradingRPCServer indicates an expected call of mustEmbedUnimplementedHeliosTradingRPCServer.
func (mr *MockUnsafeHeliosTradingRPCServerMockRecorder) mustEmbedUnimplementedHeliosTradingRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHeliosTradingRPCServer", reflect.TypeOf((*MockUnsafeHeliosTradingRPCServer)(nil).mustEmbedUnimplementedHeliosTradingRPCServer))
}
