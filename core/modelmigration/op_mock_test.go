// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/modelmigration (interfaces: Operation)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination op_mock_test.go github.com/juju/juju/core/modelmigration Operation
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	description "github.com/juju/description/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockOperation is a mock of Operation interface.
type MockOperation struct {
	ctrl     *gomock.Controller
	recorder *MockOperationMockRecorder
}

// MockOperationMockRecorder is the mock recorder for MockOperation.
type MockOperationMockRecorder struct {
	mock *MockOperation
}

// NewMockOperation creates a new mock instance.
func NewMockOperation(ctrl *gomock.Controller) *MockOperation {
	mock := &MockOperation{ctrl: ctrl}
	mock.recorder = &MockOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperation) EXPECT() *MockOperationMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockOperation) Execute(arg0 context.Context, arg1 description.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockOperationMockRecorder) Execute(arg0, arg1 any) *MockOperationExecuteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockOperation)(nil).Execute), arg0, arg1)
	return &MockOperationExecuteCall{Call: call}
}

// MockOperationExecuteCall wrap *gomock.Call
type MockOperationExecuteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOperationExecuteCall) Return(arg0 error) *MockOperationExecuteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOperationExecuteCall) Do(f func(context.Context, description.Model) error) *MockOperationExecuteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOperationExecuteCall) DoAndReturn(f func(context.Context, description.Model) error) *MockOperationExecuteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockOperation) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockOperationMockRecorder) Name() *MockOperationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOperation)(nil).Name))
	return &MockOperationNameCall{Call: call}
}

// MockOperationNameCall wrap *gomock.Call
type MockOperationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOperationNameCall) Return(arg0 string) *MockOperationNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOperationNameCall) Do(f func() string) *MockOperationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOperationNameCall) DoAndReturn(f func() string) *MockOperationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Rollback mocks base method.
func (m *MockOperation) Rollback(arg0 context.Context, arg1 description.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockOperationMockRecorder) Rollback(arg0, arg1 any) *MockOperationRollbackCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockOperation)(nil).Rollback), arg0, arg1)
	return &MockOperationRollbackCall{Call: call}
}

// MockOperationRollbackCall wrap *gomock.Call
type MockOperationRollbackCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOperationRollbackCall) Return(arg0 error) *MockOperationRollbackCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOperationRollbackCall) Do(f func(context.Context, description.Model) error) *MockOperationRollbackCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOperationRollbackCall) DoAndReturn(f func(context.Context, description.Model) error) *MockOperationRollbackCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Setup mocks base method.
func (m *MockOperation) Setup(arg0 Scope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockOperationMockRecorder) Setup(arg0 any) *MockOperationSetupCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockOperation)(nil).Setup), arg0)
	return &MockOperationSetupCall{Call: call}
}

// MockOperationSetupCall wrap *gomock.Call
type MockOperationSetupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockOperationSetupCall) Return(arg0 error) *MockOperationSetupCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockOperationSetupCall) Do(f func(Scope) error) *MockOperationSetupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockOperationSetupCall) DoAndReturn(f func(Scope) error) *MockOperationSetupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
