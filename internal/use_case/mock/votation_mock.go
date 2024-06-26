// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_votation.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockVotationRepositoryInterface is a mock of VotationRepositoryInterface interface.
type MockVotationRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVotationRepositoryInterfaceMockRecorder
}

// MockVotationRepositoryInterfaceMockRecorder is the mock recorder for MockVotationRepositoryInterface.
type MockVotationRepositoryInterfaceMockRecorder struct {
	mock *MockVotationRepositoryInterface
}

// NewMockVotationRepositoryInterface creates a new mock instance.
func NewMockVotationRepositoryInterface(ctrl *gomock.Controller) *MockVotationRepositoryInterface {
	mock := &MockVotationRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockVotationRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVotationRepositoryInterface) EXPECT() *MockVotationRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVotationRepositoryInterface) Create(votation *entity.Votation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", votation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVotationRepositoryInterfaceMockRecorder) Create(votation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).Create), votation)
}

// Deactivate mocks base method.
func (m *MockVotationRepositoryInterface) Deactivate(votation *entity.Votation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", votation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockVotationRepositoryInterfaceMockRecorder) Deactivate(votation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).Deactivate), votation)
}

// GetAll mocks base method.
func (m *MockVotationRepositoryInterface) GetAll() ([]entity.Votation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Votation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVotationRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).GetAll))
}

// GetAllByListIDAndChooserID mocks base method.
func (m *MockVotationRepositoryInterface) GetAllByListIDAndChooserID(listID, chooserID string) ([]entity.Votation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByListIDAndChooserID", listID, chooserID)
	ret0, _ := ret[0].([]entity.Votation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByListIDAndChooserID indicates an expected call of GetAllByListIDAndChooserID.
func (mr *MockVotationRepositoryInterfaceMockRecorder) GetAllByListIDAndChooserID(listID, chooserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByListIDAndChooserID", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).GetAllByListIDAndChooserID), listID, chooserID)
}

// GetByID mocks base method.
func (m *MockVotationRepositoryInterface) GetByID(votationID string) (bool, entity.Votation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", votationID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(entity.Votation)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockVotationRepositoryInterfaceMockRecorder) GetByID(votationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).GetByID), votationID)
}

// Update mocks base method.
func (m *MockVotationRepositoryInterface) Update(votation *entity.Votation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", votation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVotationRepositoryInterfaceMockRecorder) Update(votation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).Update), votation)
}

// VotationAlreadyExists mocks base method.
func (m *MockVotationRepositoryInterface) VotationAlreadyExists(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VotationAlreadyExists", chooserID, listID, firstMovieID, secondMovieID, chosenMovieID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VotationAlreadyExists indicates an expected call of VotationAlreadyExists.
func (mr *MockVotationRepositoryInterfaceMockRecorder) VotationAlreadyExists(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VotationAlreadyExists", reflect.TypeOf((*MockVotationRepositoryInterface)(nil).VotationAlreadyExists), chooserID, listID, firstMovieID, secondMovieID, chosenMovieID)
}
