// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facade (interfaces: WatcherRegistry)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/facade_mock.go github.com/juju/juju/apiserver/facade WatcherRegistry
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	worker "github.com/juju/worker/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockWatcherRegistry is a mock of WatcherRegistry interface.
type MockWatcherRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherRegistryMockRecorder
}

// MockWatcherRegistryMockRecorder is the mock recorder for MockWatcherRegistry.
type MockWatcherRegistryMockRecorder struct {
	mock *MockWatcherRegistry
}

// NewMockWatcherRegistry creates a new mock instance.
func NewMockWatcherRegistry(ctrl *gomock.Controller) *MockWatcherRegistry {
	mock := &MockWatcherRegistry{ctrl: ctrl}
	mock.recorder = &MockWatcherRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherRegistry) EXPECT() *MockWatcherRegistryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockWatcherRegistry) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockWatcherRegistryMockRecorder) Count() *MockWatcherRegistryCountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockWatcherRegistry)(nil).Count))
	return &MockWatcherRegistryCountCall{Call: call}
}

// MockWatcherRegistryCountCall wrap *gomock.Call
type MockWatcherRegistryCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryCountCall) Return(arg0 int) *MockWatcherRegistryCountCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryCountCall) Do(f func() int) *MockWatcherRegistryCountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryCountCall) DoAndReturn(f func() int) *MockWatcherRegistryCountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockWatcherRegistry) Get(arg0 string) (worker.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(worker.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockWatcherRegistryMockRecorder) Get(arg0 any) *MockWatcherRegistryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWatcherRegistry)(nil).Get), arg0)
	return &MockWatcherRegistryGetCall{Call: call}
}

// MockWatcherRegistryGetCall wrap *gomock.Call
type MockWatcherRegistryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryGetCall) Return(arg0 worker.Worker, arg1 error) *MockWatcherRegistryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryGetCall) Do(f func(string) (worker.Worker, error)) *MockWatcherRegistryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryGetCall) DoAndReturn(f func(string) (worker.Worker, error)) *MockWatcherRegistryGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Kill mocks base method.
func (m *MockWatcherRegistry) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockWatcherRegistryMockRecorder) Kill() *MockWatcherRegistryKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockWatcherRegistry)(nil).Kill))
	return &MockWatcherRegistryKillCall{Call: call}
}

// MockWatcherRegistryKillCall wrap *gomock.Call
type MockWatcherRegistryKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryKillCall) Return() *MockWatcherRegistryKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryKillCall) Do(f func()) *MockWatcherRegistryKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryKillCall) DoAndReturn(f func()) *MockWatcherRegistryKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Register mocks base method.
func (m *MockWatcherRegistry) Register(arg0 worker.Worker) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockWatcherRegistryMockRecorder) Register(arg0 any) *MockWatcherRegistryRegisterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockWatcherRegistry)(nil).Register), arg0)
	return &MockWatcherRegistryRegisterCall{Call: call}
}

// MockWatcherRegistryRegisterCall wrap *gomock.Call
type MockWatcherRegistryRegisterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryRegisterCall) Return(arg0 string, arg1 error) *MockWatcherRegistryRegisterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryRegisterCall) Do(f func(worker.Worker) (string, error)) *MockWatcherRegistryRegisterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryRegisterCall) DoAndReturn(f func(worker.Worker) (string, error)) *MockWatcherRegistryRegisterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RegisterNamed mocks base method.
func (m *MockWatcherRegistry) RegisterNamed(arg0 string, arg1 worker.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNamed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterNamed indicates an expected call of RegisterNamed.
func (mr *MockWatcherRegistryMockRecorder) RegisterNamed(arg0, arg1 any) *MockWatcherRegistryRegisterNamedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNamed", reflect.TypeOf((*MockWatcherRegistry)(nil).RegisterNamed), arg0, arg1)
	return &MockWatcherRegistryRegisterNamedCall{Call: call}
}

// MockWatcherRegistryRegisterNamedCall wrap *gomock.Call
type MockWatcherRegistryRegisterNamedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryRegisterNamedCall) Return(arg0 error) *MockWatcherRegistryRegisterNamedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryRegisterNamedCall) Do(f func(string, worker.Worker) error) *MockWatcherRegistryRegisterNamedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryRegisterNamedCall) DoAndReturn(f func(string, worker.Worker) error) *MockWatcherRegistryRegisterNamedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *MockWatcherRegistry) Stop(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockWatcherRegistryMockRecorder) Stop(arg0 any) *MockWatcherRegistryStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWatcherRegistry)(nil).Stop), arg0)
	return &MockWatcherRegistryStopCall{Call: call}
}

// MockWatcherRegistryStopCall wrap *gomock.Call
type MockWatcherRegistryStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryStopCall) Return(arg0 error) *MockWatcherRegistryStopCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryStopCall) Do(f func(string) error) *MockWatcherRegistryStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryStopCall) DoAndReturn(f func(string) error) *MockWatcherRegistryStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockWatcherRegistry) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockWatcherRegistryMockRecorder) Wait() *MockWatcherRegistryWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWatcherRegistry)(nil).Wait))
	return &MockWatcherRegistryWaitCall{Call: call}
}

// MockWatcherRegistryWaitCall wrap *gomock.Call
type MockWatcherRegistryWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherRegistryWaitCall) Return(arg0 error) *MockWatcherRegistryWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherRegistryWaitCall) Do(f func() error) *MockWatcherRegistryWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherRegistryWaitCall) DoAndReturn(f func() error) *MockWatcherRegistryWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
