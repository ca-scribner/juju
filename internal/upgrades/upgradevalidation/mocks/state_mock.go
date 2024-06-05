// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/upgrades/upgradevalidation (interfaces: StatePool,State,Model)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/state_mock.go github.com/juju/juju/internal/upgrades/upgradevalidation StatePool,State,Model
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	replicaset "github.com/juju/replicaset/v3"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockStatePool is a mock of StatePool interface.
type MockStatePool struct {
	ctrl     *gomock.Controller
	recorder *MockStatePoolMockRecorder
}

// MockStatePoolMockRecorder is the mock recorder for MockStatePool.
type MockStatePoolMockRecorder struct {
	mock *MockStatePool
}

// NewMockStatePool creates a new mock instance.
func NewMockStatePool(ctrl *gomock.Controller) *MockStatePool {
	mock := &MockStatePool{ctrl: ctrl}
	mock.recorder = &MockStatePoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatePool) EXPECT() *MockStatePoolMockRecorder {
	return m.recorder
}

// MongoVersion mocks base method.
func (m *MockStatePool) MongoVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoVersion indicates an expected call of MongoVersion.
func (mr *MockStatePoolMockRecorder) MongoVersion() *MockStatePoolMongoVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockStatePool)(nil).MongoVersion))
	return &MockStatePoolMongoVersionCall{Call: call}
}

// MockStatePoolMongoVersionCall wrap *gomock.Call
type MockStatePoolMongoVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStatePoolMongoVersionCall) Return(arg0 string, arg1 error) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStatePoolMongoVersionCall) Do(f func() (string, error)) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStatePoolMongoVersionCall) DoAndReturn(f func() (string, error)) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// AllMachinesCount mocks base method.
func (m *MockState) AllMachinesCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachinesCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachinesCount indicates an expected call of AllMachinesCount.
func (mr *MockStateMockRecorder) AllMachinesCount() *MockStateAllMachinesCountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachinesCount", reflect.TypeOf((*MockState)(nil).AllMachinesCount))
	return &MockStateAllMachinesCountCall{Call: call}
}

// MockStateAllMachinesCountCall wrap *gomock.Call
type MockStateAllMachinesCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateAllMachinesCountCall) Return(arg0 int, arg1 error) *MockStateAllMachinesCountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateAllMachinesCountCall) Do(f func() (int, error)) *MockStateAllMachinesCountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateAllMachinesCountCall) DoAndReturn(f func() (int, error)) *MockStateAllMachinesCountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasUpgradeSeriesLocks mocks base method.
func (m *MockState) HasUpgradeSeriesLocks() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpgradeSeriesLocks")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUpgradeSeriesLocks indicates an expected call of HasUpgradeSeriesLocks.
func (mr *MockStateMockRecorder) HasUpgradeSeriesLocks() *MockStateHasUpgradeSeriesLocksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpgradeSeriesLocks", reflect.TypeOf((*MockState)(nil).HasUpgradeSeriesLocks))
	return &MockStateHasUpgradeSeriesLocksCall{Call: call}
}

// MockStateHasUpgradeSeriesLocksCall wrap *gomock.Call
type MockStateHasUpgradeSeriesLocksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateHasUpgradeSeriesLocksCall) Return(arg0 bool, arg1 error) *MockStateHasUpgradeSeriesLocksCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateHasUpgradeSeriesLocksCall) Do(f func() (bool, error)) *MockStateHasUpgradeSeriesLocksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateHasUpgradeSeriesLocksCall) DoAndReturn(f func() (bool, error)) *MockStateHasUpgradeSeriesLocksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineCountForBase mocks base method.
func (m *MockState) MachineCountForBase(arg0 ...state.Base) (map[string]int, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MachineCountForBase", varargs...)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineCountForBase indicates an expected call of MachineCountForBase.
func (mr *MockStateMockRecorder) MachineCountForBase(arg0 ...any) *MockStateMachineCountForBaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineCountForBase", reflect.TypeOf((*MockState)(nil).MachineCountForBase), arg0...)
	return &MockStateMachineCountForBaseCall{Call: call}
}

// MockStateMachineCountForBaseCall wrap *gomock.Call
type MockStateMachineCountForBaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateMachineCountForBaseCall) Return(arg0 map[string]int, arg1 error) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateMachineCountForBaseCall) Do(f func(...state.Base) (map[string]int, error)) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateMachineCountForBaseCall) DoAndReturn(f func(...state.Base) (map[string]int, error)) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MongoCurrentStatus mocks base method.
func (m *MockState) MongoCurrentStatus() (*replicaset.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoCurrentStatus")
	ret0, _ := ret[0].(*replicaset.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoCurrentStatus indicates an expected call of MongoCurrentStatus.
func (mr *MockStateMockRecorder) MongoCurrentStatus() *MockStateMongoCurrentStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoCurrentStatus", reflect.TypeOf((*MockState)(nil).MongoCurrentStatus))
	return &MockStateMongoCurrentStatusCall{Call: call}
}

// MockStateMongoCurrentStatusCall wrap *gomock.Call
type MockStateMongoCurrentStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateMongoCurrentStatusCall) Return(arg0 *replicaset.Status, arg1 error) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateMongoCurrentStatusCall) Do(f func() (*replicaset.Status, error)) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateMongoCurrentStatusCall) DoAndReturn(f func() (*replicaset.Status, error)) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
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

// AgentVersion mocks base method.
func (m *MockModel) AgentVersion() (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion.
func (mr *MockModelMockRecorder) AgentVersion() *MockModelAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockModel)(nil).AgentVersion))
	return &MockModelAgentVersionCall{Call: call}
}

// MockModelAgentVersionCall wrap *gomock.Call
type MockModelAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelAgentVersionCall) Return(arg0 version.Number, arg1 error) *MockModelAgentVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelAgentVersionCall) Do(f func() (version.Number, error)) *MockModelAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelAgentVersionCall) DoAndReturn(f func() (version.Number, error)) *MockModelAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MigrationMode mocks base method.
func (m *MockModel) MigrationMode() state.MigrationMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationMode")
	ret0, _ := ret[0].(state.MigrationMode)
	return ret0
}

// MigrationMode indicates an expected call of MigrationMode.
func (mr *MockModelMockRecorder) MigrationMode() *MockModelMigrationModeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationMode", reflect.TypeOf((*MockModel)(nil).MigrationMode))
	return &MockModelMigrationModeCall{Call: call}
}

// MockModelMigrationModeCall wrap *gomock.Call
type MockModelMigrationModeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationModeCall) Return(arg0 state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationModeCall) Do(f func() state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationModeCall) DoAndReturn(f func() state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *MockModelNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
	return &MockModelNameCall{Call: call}
}

// MockModelNameCall wrap *gomock.Call
type MockModelNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelNameCall) Return(arg0 string) *MockModelNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelNameCall) Do(f func() string) *MockModelNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelNameCall) DoAndReturn(f func() string) *MockModelNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Owner mocks base method.
func (m *MockModel) Owner() names.UserTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.UserTag)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockModelMockRecorder) Owner() *MockModelOwnerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockModel)(nil).Owner))
	return &MockModelOwnerCall{Call: call}
}

// MockModelOwnerCall wrap *gomock.Call
type MockModelOwnerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelOwnerCall) Return(arg0 names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelOwnerCall) Do(f func() names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelOwnerCall) DoAndReturn(f func() names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
