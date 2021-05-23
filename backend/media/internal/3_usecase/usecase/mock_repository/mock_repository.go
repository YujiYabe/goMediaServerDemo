// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	shared "media/pkg/shared"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockToService is a mock of ToService interface.
type MockToService struct {
	ctrl     *gomock.Controller
	recorder *MockToServiceMockRecorder
}

// MockToServiceMockRecorder is the mock recorder for MockToService.
type MockToServiceMockRecorder struct {
	mock *MockToService
}

// NewMockToService creates a new mock instance.
func NewMockToService(ctrl *gomock.Controller) *MockToService {
	mock := &MockToService{ctrl: ctrl}
	mock.recorder = &MockToServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToService) EXPECT() *MockToServiceMockRecorder {
	return m.recorder
}

// SvChangeStatus mocks base method.
func (m *MockToService) SvChangeStatus(status string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SvChangeStatus", status)
}

// SvChangeStatus indicates an expected call of SvChangeStatus.
func (mr *MockToServiceMockRecorder) SvChangeStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvChangeStatus", reflect.TypeOf((*MockToService)(nil).SvChangeStatus), status)
}

// SvChangeVlcProgress mocks base method.
func (m *MockToService) SvChangeVlcProgress(value float32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SvChangeVlcProgress", value)
}

// SvChangeVlcProgress indicates an expected call of SvChangeVlcProgress.
func (mr *MockToServiceMockRecorder) SvChangeVlcProgress(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvChangeVlcProgress", reflect.TypeOf((*MockToService)(nil).SvChangeVlcProgress), value)
}

// SvReceiveContent mocks base method.
func (m *MockToService) SvReceiveContent(address, funcName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SvReceiveContent", address, funcName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SvReceiveContent indicates an expected call of SvReceiveContent.
func (mr *MockToServiceMockRecorder) SvReceiveContent(address, funcName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvReceiveContent", reflect.TypeOf((*MockToService)(nil).SvReceiveContent), address, funcName)
}

// SvSendContent mocks base method.
func (m *MockToService) SvSendContent(address string, cc *shared.CommonContent) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SvSendContent", address, cc)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SvSendContent indicates an expected call of SvSendContent.
func (mr *MockToServiceMockRecorder) SvSendContent(address, cc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvSendContent", reflect.TypeOf((*MockToService)(nil).SvSendContent), address, cc)
}

// SvSwapPlayList mocks base method.
func (m *MockToService) SvSwapPlayList(playListString string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SvSwapPlayList", playListString)
}

// SvSwapPlayList indicates an expected call of SvSwapPlayList.
func (mr *MockToServiceMockRecorder) SvSwapPlayList(playListString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvSwapPlayList", reflect.TypeOf((*MockToService)(nil).SvSwapPlayList), playListString)
}

// MockToDomain is a mock of ToDomain interface.
type MockToDomain struct {
	ctrl     *gomock.Controller
	recorder *MockToDomainMockRecorder
}

// MockToDomainMockRecorder is the mock recorder for MockToDomain.
type MockToDomainMockRecorder struct {
	mock *MockToDomain
}

// NewMockToDomain creates a new mock instance.
func NewMockToDomain(ctrl *gomock.Controller) *MockToDomain {
	mock := &MockToDomain{ctrl: ctrl}
	mock.recorder = &MockToDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToDomain) EXPECT() *MockToDomainMockRecorder {
	return m.recorder
}

// ParseKey mocks base method.
func (m *MockToDomain) ParseKey(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseKey", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseKey indicates an expected call of ParseKey.
func (mr *MockToDomainMockRecorder) ParseKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseKey", reflect.TypeOf((*MockToDomain)(nil).ParseKey), key)
}

// StringToFloat32 mocks base method.
func (m *MockToDomain) StringToFloat32(valueString string) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringToFloat32", valueString)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StringToFloat32 indicates an expected call of StringToFloat32.
func (mr *MockToDomainMockRecorder) StringToFloat32(valueString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringToFloat32", reflect.TypeOf((*MockToDomain)(nil).StringToFloat32), valueString)
}