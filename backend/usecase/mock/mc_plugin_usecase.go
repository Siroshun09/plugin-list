// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/mc_plugin_usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase/mc_plugin_usecase.go -destination=usecase/mock/mc_plugin_usecase.go
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/Siroshun09/plugin-list/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockMCPluginUseCase is a mock of MCPluginUseCase interface.
type MockMCPluginUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMCPluginUseCaseMockRecorder
}

// MockMCPluginUseCaseMockRecorder is the mock recorder for MockMCPluginUseCase.
type MockMCPluginUseCaseMockRecorder struct {
	mock *MockMCPluginUseCase
}

// NewMockMCPluginUseCase creates a new mock instance.
func NewMockMCPluginUseCase(ctrl *gomock.Controller) *MockMCPluginUseCase {
	mock := &MockMCPluginUseCase{ctrl: ctrl}
	mock.recorder = &MockMCPluginUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMCPluginUseCase) EXPECT() *MockMCPluginUseCaseMockRecorder {
	return m.recorder
}

// DeleteMCPlugin mocks base method.
func (m *MockMCPluginUseCase) DeleteMCPlugin(ctx context.Context, serverName, pluginName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMCPlugin", ctx, serverName, pluginName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMCPlugin indicates an expected call of DeleteMCPlugin.
func (mr *MockMCPluginUseCaseMockRecorder) DeleteMCPlugin(ctx, serverName, pluginName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMCPlugin", reflect.TypeOf((*MockMCPluginUseCase)(nil).DeleteMCPlugin), ctx, serverName, pluginName)
}

// GetInstalledPluginInfo mocks base method.
func (m *MockMCPluginUseCase) GetInstalledPluginInfo(ctx context.Context, pluginName string) ([]domain.MCPlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstalledPluginInfo", ctx, pluginName)
	ret0, _ := ret[0].([]domain.MCPlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstalledPluginInfo indicates an expected call of GetInstalledPluginInfo.
func (mr *MockMCPluginUseCaseMockRecorder) GetInstalledPluginInfo(ctx, pluginName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstalledPluginInfo", reflect.TypeOf((*MockMCPluginUseCase)(nil).GetInstalledPluginInfo), ctx, pluginName)
}

// GetMCPluginsByServerName mocks base method.
func (m *MockMCPluginUseCase) GetMCPluginsByServerName(ctx context.Context, serverName string) ([]domain.MCPlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMCPluginsByServerName", ctx, serverName)
	ret0, _ := ret[0].([]domain.MCPlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMCPluginsByServerName indicates an expected call of GetMCPluginsByServerName.
func (mr *MockMCPluginUseCaseMockRecorder) GetMCPluginsByServerName(ctx, serverName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMCPluginsByServerName", reflect.TypeOf((*MockMCPluginUseCase)(nil).GetMCPluginsByServerName), ctx, serverName)
}

// GetPluginNames mocks base method.
func (m *MockMCPluginUseCase) GetPluginNames(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginNames", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginNames indicates an expected call of GetPluginNames.
func (mr *MockMCPluginUseCaseMockRecorder) GetPluginNames(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginNames", reflect.TypeOf((*MockMCPluginUseCase)(nil).GetPluginNames), ctx)
}

// GetServerNames mocks base method.
func (m *MockMCPluginUseCase) GetServerNames(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerNames", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerNames indicates an expected call of GetServerNames.
func (mr *MockMCPluginUseCaseMockRecorder) GetServerNames(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerNames", reflect.TypeOf((*MockMCPluginUseCase)(nil).GetServerNames), ctx)
}

// SubmitMCPlugin mocks base method.
func (m *MockMCPluginUseCase) SubmitMCPlugin(ctx context.Context, plugin domain.MCPlugin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitMCPlugin", ctx, plugin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitMCPlugin indicates an expected call of SubmitMCPlugin.
func (mr *MockMCPluginUseCaseMockRecorder) SubmitMCPlugin(ctx, plugin any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitMCPlugin", reflect.TypeOf((*MockMCPluginUseCase)(nil).SubmitMCPlugin), ctx, plugin)
}
