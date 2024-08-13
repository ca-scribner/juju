// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/machine (interfaces: ModelConfigAPI)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/modelconfig_api_mock.go github.com/juju/juju/cmd/juju/machine ModelConfigAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigAPI is a mock of ModelConfigAPI interface.
type MockModelConfigAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigAPIMockRecorder
}

// MockModelConfigAPIMockRecorder is the mock recorder for MockModelConfigAPI.
type MockModelConfigAPIMockRecorder struct {
	mock *MockModelConfigAPI
}

// NewMockModelConfigAPI creates a new mock instance.
func NewMockModelConfigAPI(ctrl *gomock.Controller) *MockModelConfigAPI {
	mock := &MockModelConfigAPI{ctrl: ctrl}
	mock.recorder = &MockModelConfigAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigAPI) EXPECT() *MockModelConfigAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockModelConfigAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockModelConfigAPIMockRecorder) Close() *MockModelConfigAPICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelConfigAPI)(nil).Close))
	return &MockModelConfigAPICloseCall{Call: call}
}

// MockModelConfigAPICloseCall wrap *gomock.Call
type MockModelConfigAPICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigAPICloseCall) Return(arg0 error) *MockModelConfigAPICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigAPICloseCall) Do(f func() error) *MockModelConfigAPICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigAPICloseCall) DoAndReturn(f func() error) *MockModelConfigAPICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelGet mocks base method.
func (m *MockModelConfigAPI) ModelGet(arg0 context.Context) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet", arg0)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockModelConfigAPIMockRecorder) ModelGet(arg0 any) *MockModelConfigAPIModelGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigAPI)(nil).ModelGet), arg0)
	return &MockModelConfigAPIModelGetCall{Call: call}
}

// MockModelConfigAPIModelGetCall wrap *gomock.Call
type MockModelConfigAPIModelGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigAPIModelGetCall) Return(arg0 map[string]any, arg1 error) *MockModelConfigAPIModelGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigAPIModelGetCall) Do(f func(context.Context) (map[string]any, error)) *MockModelConfigAPIModelGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigAPIModelGetCall) DoAndReturn(f func(context.Context) (map[string]any, error)) *MockModelConfigAPIModelGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
