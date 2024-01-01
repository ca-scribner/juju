// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/action (interfaces: State,Model)

// Package action is a generated GoMock package.
package action

import (
	reflect "reflect"

	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AllApplications mocks base method.
func (m *MockState) AllApplications() ([]*state.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplications")
	ret0, _ := ret[0].([]*state.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplications indicates an expected call of AllApplications.
func (mr *MockStateMockRecorder) AllApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplications", reflect.TypeOf((*MockState)(nil).AllApplications))
}

// AllMachines mocks base method.
func (m *MockState) AllMachines() ([]*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines.
func (mr *MockStateMockRecorder) AllMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockState)(nil).AllMachines))
}

// Application mocks base method.
func (m *MockState) Application(arg0 string) (*state.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(*state.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockState)(nil).Application), arg0)
}

// FindEntity mocks base method.
func (m *MockState) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockStateMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockState)(nil).FindEntity), arg0)
}

// GetBlockForType mocks base method.
func (m *MockState) GetBlockForType(arg0 state.BlockType) (state.Block, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockForType", arg0)
	ret0, _ := ret[0].(state.Block)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBlockForType indicates an expected call of GetBlockForType.
func (mr *MockStateMockRecorder) GetBlockForType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockForType", reflect.TypeOf((*MockState)(nil).GetBlockForType), arg0)
}

// Model mocks base method.
func (m *MockState) Model() (Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
}

// WatchActionLogs mocks base method.
func (m *MockState) WatchActionLogs(arg0 string) state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchActionLogs", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchActionLogs indicates an expected call of WatchActionLogs.
func (mr *MockStateMockRecorder) WatchActionLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchActionLogs", reflect.TypeOf((*MockState)(nil).WatchActionLogs), arg0)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// ActionByTag mocks base method.
func (m *MockModel) ActionByTag(arg0 names.ActionTag) (state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionByTag", arg0)
	ret0, _ := ret[0].(state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionByTag indicates an expected call of ActionByTag.
func (mr *MockModelMockRecorder) ActionByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionByTag", reflect.TypeOf((*MockModel)(nil).ActionByTag), arg0)
}

// AddAction mocks base method.
func (m *MockModel) AddAction(arg0 state.ActionReceiver, arg1, arg2 string, arg3 map[string]interface{}, arg4 *bool, arg5 *string) (state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAction", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAction indicates an expected call of AddAction.
func (mr *MockModelMockRecorder) AddAction(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAction", reflect.TypeOf((*MockModel)(nil).AddAction), arg0, arg1, arg2, arg3, arg4, arg5)
}

// EnqueueOperation mocks base method.
func (m *MockModel) EnqueueOperation(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueOperation", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueueOperation indicates an expected call of EnqueueOperation.
func (mr *MockModelMockRecorder) EnqueueOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueOperation", reflect.TypeOf((*MockModel)(nil).EnqueueOperation), arg0, arg1)
}

// FailOperationEnqueuing mocks base method.
func (m *MockModel) FailOperationEnqueuing(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailOperationEnqueuing", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FailOperationEnqueuing indicates an expected call of FailOperationEnqueuing.
func (mr *MockModelMockRecorder) FailOperationEnqueuing(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailOperationEnqueuing", reflect.TypeOf((*MockModel)(nil).FailOperationEnqueuing), arg0, arg1, arg2)
}

// FindActionsByName mocks base method.
func (m *MockModel) FindActionsByName(arg0 string) ([]state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActionsByName", arg0)
	ret0, _ := ret[0].([]state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActionsByName indicates an expected call of FindActionsByName.
func (mr *MockModelMockRecorder) FindActionsByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActionsByName", reflect.TypeOf((*MockModel)(nil).FindActionsByName), arg0)
}

// ListOperations mocks base method.
func (m *MockModel) ListOperations(arg0 []string, arg1 []names.Tag, arg2 []state.ActionStatus, arg3, arg4 int) ([]state.OperationInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]state.OperationInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockModelMockRecorder) ListOperations(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockModel)(nil).ListOperations), arg0, arg1, arg2, arg3, arg4)
}

// ModelTag mocks base method.
func (m *MockModel) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockModelMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockModel)(nil).ModelTag))
}

// OperationWithActions mocks base method.
func (m *MockModel) OperationWithActions(arg0 string) (*state.OperationInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationWithActions", arg0)
	ret0, _ := ret[0].(*state.OperationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperationWithActions indicates an expected call of OperationWithActions.
func (mr *MockModelMockRecorder) OperationWithActions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationWithActions", reflect.TypeOf((*MockModel)(nil).OperationWithActions), arg0)
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
}
