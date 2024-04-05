// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_director.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockDirectorRepositoryInterface is a mock of DirectorRepositoryInterface interface.
type MockDirectorRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDirectorRepositoryInterfaceMockRecorder
}

// MockDirectorRepositoryInterfaceMockRecorder is the mock recorder for MockDirectorRepositoryInterface.
type MockDirectorRepositoryInterfaceMockRecorder struct {
	mock *MockDirectorRepositoryInterface
}

// NewMockDirectorRepositoryInterface creates a new mock instance.
func NewMockDirectorRepositoryInterface(ctrl *gomock.Controller) *MockDirectorRepositoryInterface {
	mock := &MockDirectorRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockDirectorRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectorRepositoryInterface) EXPECT() *MockDirectorRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDirectorRepositoryInterface) Create(director *entity.Director) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", director)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) Create(director interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).Create), director)
}

// Deactivate mocks base method.
func (m *MockDirectorRepositoryInterface) Deactivate(director *entity.Director) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", director)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) Deactivate(director interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).Deactivate), director)
}

// DoTheseDirectorsExist mocks base method.
func (m *MockDirectorRepositoryInterface) DoTheseDirectorsExist(directorIDs []string) (bool, []entity.Director, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTheseDirectorsExist", directorIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]entity.Director)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoTheseDirectorsExist indicates an expected call of DoTheseDirectorsExist.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) DoTheseDirectorsExist(directorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTheseDirectorsExist", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).DoTheseDirectorsExist), directorIDs)
}

// GetAll mocks base method.
func (m *MockDirectorRepositoryInterface) GetAll() ([]entity.Director, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Director)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockDirectorRepositoryInterface) GetByID(directorID string) (entity.Director, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", directorID)
	ret0, _ := ret[0].(entity.Director)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) GetByID(directorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).GetByID), directorID)
}

// Update mocks base method.
func (m *MockDirectorRepositoryInterface) Update(director *entity.Director) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", director)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDirectorRepositoryInterfaceMockRecorder) Update(director interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDirectorRepositoryInterface)(nil).Update), director)
}
