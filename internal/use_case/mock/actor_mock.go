// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_actor.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockActorRepositoryInterface is a mock of ActorRepositoryInterface interface.
type MockActorRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockActorRepositoryInterfaceMockRecorder
}

// MockActorRepositoryInterfaceMockRecorder is the mock recorder for MockActorRepositoryInterface.
type MockActorRepositoryInterfaceMockRecorder struct {
	mock *MockActorRepositoryInterface
}

// NewMockActorRepositoryInterface creates a new mock instance.
func NewMockActorRepositoryInterface(ctrl *gomock.Controller) *MockActorRepositoryInterface {
	mock := &MockActorRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockActorRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorRepositoryInterface) EXPECT() *MockActorRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockActorRepositoryInterface) Create(actor *entity.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockActorRepositoryInterfaceMockRecorder) Create(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActorRepositoryInterface)(nil).Create), actor)
}

// Deactivate mocks base method.
func (m *MockActorRepositoryInterface) Deactivate(actor *entity.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockActorRepositoryInterfaceMockRecorder) Deactivate(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockActorRepositoryInterface)(nil).Deactivate), actor)
}

// DoTheseActorsExist mocks base method.
func (m *MockActorRepositoryInterface) DoTheseActorsExist(actorIDs []string) (bool, []entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTheseActorsExist", actorIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]entity.Actor)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoTheseActorsExist indicates an expected call of DoTheseActorsExist.
func (mr *MockActorRepositoryInterfaceMockRecorder) DoTheseActorsExist(actorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTheseActorsExist", reflect.TypeOf((*MockActorRepositoryInterface)(nil).DoTheseActorsExist), actorIDs)
}

// GetAll mocks base method.
func (m *MockActorRepositoryInterface) GetAll() ([]entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockActorRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockActorRepositoryInterface)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockActorRepositoryInterface) GetByID(actorID string) (entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", actorID)
	ret0, _ := ret[0].(entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockActorRepositoryInterfaceMockRecorder) GetByID(actorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockActorRepositoryInterface)(nil).GetByID), actorID)
}

// Update mocks base method.
func (m *MockActorRepositoryInterface) Update(actor *entity.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockActorRepositoryInterfaceMockRecorder) Update(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockActorRepositoryInterface)(nil).Update), actor)
}