// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_chooser.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockChooserRepositoryInterface is a mock of ChooserRepositoryInterface interface.
type MockChooserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockChooserRepositoryInterfaceMockRecorder
}

// MockChooserRepositoryInterfaceMockRecorder is the mock recorder for MockChooserRepositoryInterface.
type MockChooserRepositoryInterfaceMockRecorder struct {
	mock *MockChooserRepositoryInterface
}

// NewMockChooserRepositoryInterface creates a new mock instance.
func NewMockChooserRepositoryInterface(ctrl *gomock.Controller) *MockChooserRepositoryInterface {
	mock := &MockChooserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockChooserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChooserRepositoryInterface) EXPECT() *MockChooserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// ChooserAlreadyExists mocks base method.
func (m *MockChooserRepositoryInterface) ChooserAlreadyExists(chooserEmail string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChooserAlreadyExists", chooserEmail)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChooserAlreadyExists indicates an expected call of ChooserAlreadyExists.
func (mr *MockChooserRepositoryInterfaceMockRecorder) ChooserAlreadyExists(chooserEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChooserAlreadyExists", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).ChooserAlreadyExists), chooserEmail)
}

// Create mocks base method.
func (m *MockChooserRepositoryInterface) Create(chooser *entity.Chooser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", chooser)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockChooserRepositoryInterfaceMockRecorder) Create(chooser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).Create), chooser)
}

// Deactivate mocks base method.
func (m *MockChooserRepositoryInterface) Deactivate(chooser *entity.Chooser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", chooser)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockChooserRepositoryInterfaceMockRecorder) Deactivate(chooser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).Deactivate), chooser)
}

// GetAll mocks base method.
func (m *MockChooserRepositoryInterface) GetAll() ([]entity.Chooser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Chooser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChooserRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).GetAll))
}

// GetByEmail mocks base method.
func (m *MockChooserRepositoryInterface) GetByEmail(chooserEmail string) (entity.Chooser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", chooserEmail)
	ret0, _ := ret[0].(entity.Chooser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockChooserRepositoryInterfaceMockRecorder) GetByEmail(chooserEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).GetByEmail), chooserEmail)
}

// GetByID mocks base method.
func (m *MockChooserRepositoryInterface) GetByID(chooserID string) (bool, entity.Chooser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", chooserID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(entity.Chooser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockChooserRepositoryInterfaceMockRecorder) GetByID(chooserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).GetByID), chooserID)
}

// GetVotation mocks base method.
func (m *MockChooserRepositoryInterface) GetVotation(chooserID, listID string) (entity.Votation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotation", chooserID, listID)
	ret0, _ := ret[0].(entity.Votation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotation indicates an expected call of GetVotation.
func (mr *MockChooserRepositoryInterfaceMockRecorder) GetVotation(chooserID, listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotation", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).GetVotation), chooserID, listID)
}

// GetVotations mocks base method.
func (m *MockChooserRepositoryInterface) GetVotations(chooserID string) ([]entity.Votation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotations", chooserID)
	ret0, _ := ret[0].([]entity.Votation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotations indicates an expected call of GetVotations.
func (mr *MockChooserRepositoryInterfaceMockRecorder) GetVotations(chooserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotations", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).GetVotations), chooserID)
}

// Update mocks base method.
func (m *MockChooserRepositoryInterface) Update(chooser *entity.Chooser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", chooser)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockChooserRepositoryInterfaceMockRecorder) Update(chooser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChooserRepositoryInterface)(nil).Update), chooser)
}
