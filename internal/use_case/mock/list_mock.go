// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_list.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockListRepositoryInterface is a mock of ListRepositoryInterface interface.
type MockListRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockListRepositoryInterfaceMockRecorder
}

// MockListRepositoryInterfaceMockRecorder is the mock recorder for MockListRepositoryInterface.
type MockListRepositoryInterfaceMockRecorder struct {
	mock *MockListRepositoryInterface
}

// NewMockListRepositoryInterface creates a new mock instance.
func NewMockListRepositoryInterface(ctrl *gomock.Controller) *MockListRepositoryInterface {
	mock := &MockListRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockListRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListRepositoryInterface) EXPECT() *MockListRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockListRepositoryInterface) Create(list *entity.List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", list)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockListRepositoryInterfaceMockRecorder) Create(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockListRepositoryInterface)(nil).Create), list)
}

// Deactivate mocks base method.
func (m *MockListRepositoryInterface) Deactivate(listID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", listID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockListRepositoryInterfaceMockRecorder) Deactivate(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockListRepositoryInterface)(nil).Deactivate), listID)
}

// GetAll mocks base method.
func (m *MockListRepositoryInterface) GetAll() ([]entity.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockListRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockListRepositoryInterface)(nil).GetAll))
}

// GetAllMoviesByListID mocks base method.
func (m *MockListRepositoryInterface) GetAllMoviesByListID(listID string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMoviesByListID", listID)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMoviesByListID indicates an expected call of GetAllMoviesByListID.
func (mr *MockListRepositoryInterfaceMockRecorder) GetAllMoviesByListID(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMoviesByListID", reflect.TypeOf((*MockListRepositoryInterface)(nil).GetAllMoviesByListID), listID)
}

// GetAllMoviesCombinationsByListID mocks base method.
func (m *MockListRepositoryInterface) GetAllMoviesCombinationsByListID(listID string) ([][]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMoviesCombinationsByListID", listID)
	ret0, _ := ret[0].([][]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMoviesCombinationsByListID indicates an expected call of GetAllMoviesCombinationsByListID.
func (mr *MockListRepositoryInterfaceMockRecorder) GetAllMoviesCombinationsByListID(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMoviesCombinationsByListID", reflect.TypeOf((*MockListRepositoryInterface)(nil).GetAllMoviesCombinationsByListID), listID)
}

// GetByID mocks base method.
func (m *MockListRepositoryInterface) GetByID(listID string) (bool, entity.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", listID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(entity.List)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockListRepositoryInterfaceMockRecorder) GetByID(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockListRepositoryInterface)(nil).GetByID), listID)
}

// Update mocks base method.
func (m *MockListRepositoryInterface) Update(list *entity.List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", list)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockListRepositoryInterfaceMockRecorder) Update(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockListRepositoryInterface)(nil).Update), list)
}