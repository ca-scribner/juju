// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: SecretsStore,SecretBackendsStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	set "github.com/juju/collections/set"
	secrets "github.com/juju/juju/core/secrets"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretsStore is a mock of SecretsStore interface.
type MockSecretsStore struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsStoreMockRecorder
}

// MockSecretsStoreMockRecorder is the mock recorder for MockSecretsStore.
type MockSecretsStoreMockRecorder struct {
	mock *MockSecretsStore
}

// NewMockSecretsStore creates a new mock instance.
func NewMockSecretsStore(ctrl *gomock.Controller) *MockSecretsStore {
	mock := &MockSecretsStore{ctrl: ctrl}
	mock.recorder = &MockSecretsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsStore) EXPECT() *MockSecretsStoreMockRecorder {
	return m.recorder
}

// ChangeSecretBackend mocks base method.
func (m *MockSecretsStore) ChangeSecretBackend(arg0 state.ChangeSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSecretBackend", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSecretBackend indicates an expected call of ChangeSecretBackend.
func (mr *MockSecretsStoreMockRecorder) ChangeSecretBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSecretBackend", reflect.TypeOf((*MockSecretsStore)(nil).ChangeSecretBackend), arg0)
}

// CreateSecret mocks base method.
func (m *MockSecretsStore) CreateSecret(arg0 *secrets.URI, arg1 state.CreateSecretParams) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretsStoreMockRecorder) CreateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretsStore)(nil).CreateSecret), arg0, arg1)
}

// DeleteSecret mocks base method.
func (m *MockSecretsStore) DeleteSecret(arg0 *secrets.URI, arg1 ...int) ([]secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSecret", varargs...)
	ret0, _ := ret[0].([]secrets.ValueRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretsStoreMockRecorder) DeleteSecret(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretsStore)(nil).DeleteSecret), varargs...)
}

// GetSecret mocks base method.
func (m *MockSecretsStore) GetSecret(arg0 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretsStoreMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsStore)(nil).GetSecret), arg0)
}

// GetSecretRevision mocks base method.
func (m *MockSecretsStore) GetSecretRevision(arg0 *secrets.URI, arg1 int) (*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretRevision", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretRevisionMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretRevision indicates an expected call of GetSecretRevision.
func (mr *MockSecretsStoreMockRecorder) GetSecretRevision(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretRevision", reflect.TypeOf((*MockSecretsStore)(nil).GetSecretRevision), arg0, arg1)
}

// GetSecretValue mocks base method.
func (m *MockSecretsStore) GetSecretValue(arg0 *secrets.URI, arg1 int) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsStoreMockRecorder) GetSecretValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsStore)(nil).GetSecretValue), arg0, arg1)
}

// ListModelSecrets mocks base method.
func (m *MockSecretsStore) ListModelSecrets(arg0 bool) (map[string]set.Strings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelSecrets", arg0)
	ret0, _ := ret[0].(map[string]set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelSecrets indicates an expected call of ListModelSecrets.
func (mr *MockSecretsStoreMockRecorder) ListModelSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelSecrets", reflect.TypeOf((*MockSecretsStore)(nil).ListModelSecrets), arg0)
}

// ListSecretRevisions mocks base method.
func (m *MockSecretsStore) ListSecretRevisions(arg0 *secrets.URI) ([]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretRevisions", arg0)
	ret0, _ := ret[0].([]*secrets.SecretRevisionMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretRevisions indicates an expected call of ListSecretRevisions.
func (mr *MockSecretsStoreMockRecorder) ListSecretRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretRevisions", reflect.TypeOf((*MockSecretsStore)(nil).ListSecretRevisions), arg0)
}

// ListSecrets mocks base method.
func (m *MockSecretsStore) ListSecrets(arg0 state.SecretsFilter) ([]*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsStoreMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsStore)(nil).ListSecrets), arg0)
}

// ListUnusedSecretRevisions mocks base method.
func (m *MockSecretsStore) ListUnusedSecretRevisions(arg0 *secrets.URI) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnusedSecretRevisions", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnusedSecretRevisions indicates an expected call of ListUnusedSecretRevisions.
func (mr *MockSecretsStoreMockRecorder) ListUnusedSecretRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnusedSecretRevisions", reflect.TypeOf((*MockSecretsStore)(nil).ListUnusedSecretRevisions), arg0)
}

// SecretGrants mocks base method.
func (m *MockSecretsStore) SecretGrants(arg0 *secrets.URI, arg1 secrets.SecretRole) ([]secrets.AccessInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretGrants", arg0, arg1)
	ret0, _ := ret[0].([]secrets.AccessInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGrants indicates an expected call of SecretGrants.
func (mr *MockSecretsStoreMockRecorder) SecretGrants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGrants", reflect.TypeOf((*MockSecretsStore)(nil).SecretGrants), arg0, arg1)
}

// UpdateSecret mocks base method.
func (m *MockSecretsStore) UpdateSecret(arg0 *secrets.URI, arg1 state.UpdateSecretParams) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretsStoreMockRecorder) UpdateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretsStore)(nil).UpdateSecret), arg0, arg1)
}

// WatchObsolete mocks base method.
func (m *MockSecretsStore) WatchObsolete(arg0 []names.Tag) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchObsolete", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchObsolete indicates an expected call of WatchObsolete.
func (mr *MockSecretsStoreMockRecorder) WatchObsolete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchObsolete", reflect.TypeOf((*MockSecretsStore)(nil).WatchObsolete), arg0)
}

// WatchRevisionsToPrune mocks base method.
func (m *MockSecretsStore) WatchRevisionsToPrune(arg0 []names.Tag) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRevisionsToPrune", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRevisionsToPrune indicates an expected call of WatchRevisionsToPrune.
func (mr *MockSecretsStoreMockRecorder) WatchRevisionsToPrune(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRevisionsToPrune", reflect.TypeOf((*MockSecretsStore)(nil).WatchRevisionsToPrune), arg0)
}

// MockSecretBackendsStorage is a mock of SecretBackendsStorage interface.
type MockSecretBackendsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendsStorageMockRecorder
}

// MockSecretBackendsStorageMockRecorder is the mock recorder for MockSecretBackendsStorage.
type MockSecretBackendsStorageMockRecorder struct {
	mock *MockSecretBackendsStorage
}

// NewMockSecretBackendsStorage creates a new mock instance.
func NewMockSecretBackendsStorage(ctrl *gomock.Controller) *MockSecretBackendsStorage {
	mock := &MockSecretBackendsStorage{ctrl: ctrl}
	mock.recorder = &MockSecretBackendsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendsStorage) EXPECT() *MockSecretBackendsStorageMockRecorder {
	return m.recorder
}

// CreateSecretBackend mocks base method.
func (m *MockSecretBackendsStorage) CreateSecretBackend(arg0 state.CreateSecretBackendParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecretBackend", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecretBackend indicates an expected call of CreateSecretBackend.
func (mr *MockSecretBackendsStorageMockRecorder) CreateSecretBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecretBackend", reflect.TypeOf((*MockSecretBackendsStorage)(nil).CreateSecretBackend), arg0)
}

// DeleteSecretBackend mocks base method.
func (m *MockSecretBackendsStorage) DeleteSecretBackend(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecretBackend indicates an expected call of DeleteSecretBackend.
func (mr *MockSecretBackendsStorageMockRecorder) DeleteSecretBackend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecretBackend", reflect.TypeOf((*MockSecretBackendsStorage)(nil).DeleteSecretBackend), arg0, arg1)
}

// GetSecretBackend mocks base method.
func (m *MockSecretBackendsStorage) GetSecretBackend(arg0 string) (*secrets.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackend", arg0)
	ret0, _ := ret[0].(*secrets.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackend indicates an expected call of GetSecretBackend.
func (mr *MockSecretBackendsStorageMockRecorder) GetSecretBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackend", reflect.TypeOf((*MockSecretBackendsStorage)(nil).GetSecretBackend), arg0)
}

// GetSecretBackendByID mocks base method.
func (m *MockSecretBackendsStorage) GetSecretBackendByID(arg0 string) (*secrets.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendByID", arg0)
	ret0, _ := ret[0].(*secrets.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendByID indicates an expected call of GetSecretBackendByID.
func (mr *MockSecretBackendsStorageMockRecorder) GetSecretBackendByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendByID", reflect.TypeOf((*MockSecretBackendsStorage)(nil).GetSecretBackendByID), arg0)
}

// ListSecretBackends mocks base method.
func (m *MockSecretBackendsStorage) ListSecretBackends() ([]*secrets.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretBackends")
	ret0, _ := ret[0].([]*secrets.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretBackends indicates an expected call of ListSecretBackends.
func (mr *MockSecretBackendsStorageMockRecorder) ListSecretBackends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretBackends", reflect.TypeOf((*MockSecretBackendsStorage)(nil).ListSecretBackends))
}

// SecretBackendRotated mocks base method.
func (m *MockSecretBackendsStorage) SecretBackendRotated(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBackendRotated", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretBackendRotated indicates an expected call of SecretBackendRotated.
func (mr *MockSecretBackendsStorageMockRecorder) SecretBackendRotated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBackendRotated", reflect.TypeOf((*MockSecretBackendsStorage)(nil).SecretBackendRotated), arg0, arg1)
}

// UpdateSecretBackend mocks base method.
func (m *MockSecretBackendsStorage) UpdateSecretBackend(arg0 state.UpdateSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretBackend", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretBackend indicates an expected call of UpdateSecretBackend.
func (mr *MockSecretBackendsStorageMockRecorder) UpdateSecretBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretBackend", reflect.TypeOf((*MockSecretBackendsStorage)(nil).UpdateSecretBackend), arg0)
}
