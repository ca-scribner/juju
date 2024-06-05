// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/secret/service (interfaces: WatcherFactory)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination watcherfactory_mock_test.go github.com/juju/juju/domain/secret/service WatcherFactory
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	watcher "github.com/juju/juju/core/watcher"
	eventsource "github.com/juju/juju/core/watcher/eventsource"
	gomock "go.uber.org/mock/gomock"
)

// MockWatcherFactory is a mock of WatcherFactory interface.
type MockWatcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherFactoryMockRecorder
}

// MockWatcherFactoryMockRecorder is the mock recorder for MockWatcherFactory.
type MockWatcherFactoryMockRecorder struct {
	mock *MockWatcherFactory
}

// NewMockWatcherFactory creates a new mock instance.
func NewMockWatcherFactory(ctrl *gomock.Controller) *MockWatcherFactory {
	mock := &MockWatcherFactory{ctrl: ctrl}
	mock.recorder = &MockWatcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherFactory) EXPECT() *MockWatcherFactoryMockRecorder {
	return m.recorder
}

// NewNamespaceMapperWatcher mocks base method.
func (m *MockWatcherFactory) NewNamespaceMapperWatcher(arg0 string, arg1 changestream.ChangeType, arg2 eventsource.NamespaceQuery, arg3 eventsource.Mapper) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNamespaceMapperWatcher", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNamespaceMapperWatcher indicates an expected call of NewNamespaceMapperWatcher.
func (mr *MockWatcherFactoryMockRecorder) NewNamespaceMapperWatcher(arg0, arg1, arg2, arg3 any) *MockWatcherFactoryNewNamespaceMapperWatcherCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNamespaceMapperWatcher", reflect.TypeOf((*MockWatcherFactory)(nil).NewNamespaceMapperWatcher), arg0, arg1, arg2, arg3)
	return &MockWatcherFactoryNewNamespaceMapperWatcherCall{Call: call}
}

// MockWatcherFactoryNewNamespaceMapperWatcherCall wrap *gomock.Call
type MockWatcherFactoryNewNamespaceMapperWatcherCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherFactoryNewNamespaceMapperWatcherCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockWatcherFactoryNewNamespaceMapperWatcherCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherFactoryNewNamespaceMapperWatcherCall) Do(f func(string, changestream.ChangeType, eventsource.NamespaceQuery, eventsource.Mapper) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewNamespaceMapperWatcherCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherFactoryNewNamespaceMapperWatcherCall) DoAndReturn(f func(string, changestream.ChangeType, eventsource.NamespaceQuery, eventsource.Mapper) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewNamespaceMapperWatcherCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewNamespaceWatcher mocks base method.
func (m *MockWatcherFactory) NewNamespaceWatcher(arg0 string, arg1 changestream.ChangeType, arg2 eventsource.NamespaceQuery) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNamespaceWatcher", arg0, arg1, arg2)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNamespaceWatcher indicates an expected call of NewNamespaceWatcher.
func (mr *MockWatcherFactoryMockRecorder) NewNamespaceWatcher(arg0, arg1, arg2 any) *MockWatcherFactoryNewNamespaceWatcherCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNamespaceWatcher", reflect.TypeOf((*MockWatcherFactory)(nil).NewNamespaceWatcher), arg0, arg1, arg2)
	return &MockWatcherFactoryNewNamespaceWatcherCall{Call: call}
}

// MockWatcherFactoryNewNamespaceWatcherCall wrap *gomock.Call
type MockWatcherFactoryNewNamespaceWatcherCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherFactoryNewNamespaceWatcherCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockWatcherFactoryNewNamespaceWatcherCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherFactoryNewNamespaceWatcherCall) Do(f func(string, changestream.ChangeType, eventsource.NamespaceQuery) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewNamespaceWatcherCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherFactoryNewNamespaceWatcherCall) DoAndReturn(f func(string, changestream.ChangeType, eventsource.NamespaceQuery) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewNamespaceWatcherCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
