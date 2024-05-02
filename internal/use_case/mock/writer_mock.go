// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_writer.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockWriterRepositoryInterface is a mock of WriterRepositoryInterface interface.
type MockWriterRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWriterRepositoryInterfaceMockRecorder
}

// MockWriterRepositoryInterfaceMockRecorder is the mock recorder for MockWriterRepositoryInterface.
type MockWriterRepositoryInterfaceMockRecorder struct {
	mock *MockWriterRepositoryInterface
}

// NewMockWriterRepositoryInterface creates a new mock instance.
func NewMockWriterRepositoryInterface(ctrl *gomock.Controller) *MockWriterRepositoryInterface {
	mock := &MockWriterRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockWriterRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriterRepositoryInterface) EXPECT() *MockWriterRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriterRepositoryInterface) Create(writer *entity.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWriterRepositoryInterfaceMockRecorder) Create(writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).Create), writer)
}

// Deactivate mocks base method.
func (m *MockWriterRepositoryInterface) Deactivate(writer *entity.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockWriterRepositoryInterfaceMockRecorder) Deactivate(writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).Deactivate), writer)
}

// DoTheseWritersAreIncludedInTheMovie mocks base method.
func (m *MockWriterRepositoryInterface) DoTheseWritersAreIncludedInTheMovie(movieID string, writersIDs []string) (bool, []entity.Writer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTheseWritersAreIncludedInTheMovie", movieID, writersIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]entity.Writer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoTheseWritersAreIncludedInTheMovie indicates an expected call of DoTheseWritersAreIncludedInTheMovie.
func (mr *MockWriterRepositoryInterfaceMockRecorder) DoTheseWritersAreIncludedInTheMovie(movieID, writersIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTheseWritersAreIncludedInTheMovie", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).DoTheseWritersAreIncludedInTheMovie), movieID, writersIDs)
}

// DoTheseWritersExist mocks base method.
func (m *MockWriterRepositoryInterface) DoTheseWritersExist(writerIDs []string) (bool, []entity.Writer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTheseWritersExist", writerIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]entity.Writer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoTheseWritersExist indicates an expected call of DoTheseWritersExist.
func (mr *MockWriterRepositoryInterfaceMockRecorder) DoTheseWritersExist(writerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTheseWritersExist", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).DoTheseWritersExist), writerIDs)
}

// GetAll mocks base method.
func (m *MockWriterRepositoryInterface) GetAll() ([]entity.Writer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWriterRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockWriterRepositoryInterface) GetByID(writerID string) (entity.Writer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", writerID)
	ret0, _ := ret[0].(entity.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockWriterRepositoryInterfaceMockRecorder) GetByID(writerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).GetByID), writerID)
}

// Update mocks base method.
func (m *MockWriterRepositoryInterface) Update(writer *entity.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWriterRepositoryInterfaceMockRecorder) Update(writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriterRepositoryInterface)(nil).Update), writer)
}
