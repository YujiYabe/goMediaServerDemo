// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	domain "device/internal/4_domain/domain"
	shared "device/pkg/shared"
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

// SvSendIRData mocks base method.
func (m *MockToService) SvSendIRData(room, irKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SvSendIRData", room, irKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// SvSendIRData indicates an expected call of SvSendIRData.
func (mr *MockToServiceMockRecorder) SvSendIRData(room, irKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SvSendIRData", reflect.TypeOf((*MockToService)(nil).SvSendIRData), room, irKey)
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

// ChangeDeviceStatus mocks base method.
func (m *MockToDomain) ChangeDeviceStatus(cc *shared.CommonContent, rooms map[string]*domain.Room) map[string]*domain.Room {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDeviceStatus", cc, rooms)
	ret0, _ := ret[0].(map[string]*domain.Room)
	return ret0
}

// ChangeDeviceStatus indicates an expected call of ChangeDeviceStatus.
func (mr *MockToDomainMockRecorder) ChangeDeviceStatus(cc, rooms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDeviceStatus", reflect.TypeOf((*MockToDomain)(nil).ChangeDeviceStatus), cc, rooms)
}

// GetBroadcastStatus mocks base method.
func (m *MockToDomain) GetBroadcastStatus(object string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBroadcastStatus", object)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBroadcastStatus indicates an expected call of GetBroadcastStatus.
func (mr *MockToDomainMockRecorder) GetBroadcastStatus(object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBroadcastStatus", reflect.TypeOf((*MockToDomain)(nil).GetBroadcastStatus), object)
}

// GetDefaltValue mocks base method.
func (m *MockToDomain) GetDefaltValue() *domain.Room {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaltValue")
	ret0, _ := ret[0].(*domain.Room)
	return ret0
}

// GetDefaltValue indicates an expected call of GetDefaltValue.
func (mr *MockToDomainMockRecorder) GetDefaltValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaltValue", reflect.TypeOf((*MockToDomain)(nil).GetDefaltValue))
}

// JSONMarshal mocks base method.
func (m *MockToDomain) JSONMarshal(rooms map[string]*domain.Room) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONMarshal", rooms)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSONMarshal indicates an expected call of JSONMarshal.
func (mr *MockToDomainMockRecorder) JSONMarshal(rooms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONMarshal", reflect.TypeOf((*MockToDomain)(nil).JSONMarshal), rooms)
}

// StrConvAtoi mocks base method.
func (m *MockToDomain) StrConvAtoi(valueString string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StrConvAtoi", valueString)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StrConvAtoi indicates an expected call of StrConvAtoi.
func (mr *MockToDomainMockRecorder) StrConvAtoi(valueString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrConvAtoi", reflect.TypeOf((*MockToDomain)(nil).StrConvAtoi), valueString)
}
