// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository_interface/repository_interface_movie.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "youchoose/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockMovieRepositoryInterface is a mock of MovieRepositoryInterface interface.
type MockMovieRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryInterfaceMockRecorder
}

// MockMovieRepositoryInterfaceMockRecorder is the mock recorder for MockMovieRepositoryInterface.
type MockMovieRepositoryInterfaceMockRecorder struct {
	mock *MockMovieRepositoryInterface
}

// NewMockMovieRepositoryInterface creates a new mock instance.
func NewMockMovieRepositoryInterface(ctrl *gomock.Controller) *MockMovieRepositoryInterface {
	mock := &MockMovieRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepositoryInterface) EXPECT() *MockMovieRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMovieRepositoryInterface) Create(movie *entity.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMovieRepositoryInterfaceMockRecorder) Create(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).Create), movie)
}

// Deactivate mocks base method.
func (m *MockMovieRepositoryInterface) Deactivate(movieID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", movieID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockMovieRepositoryInterfaceMockRecorder) Deactivate(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).Deactivate), movieID)
}

// DoTheseMoviesExist mocks base method.
func (m *MockMovieRepositoryInterface) DoTheseMoviesExist(movieIDs []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTheseMoviesExist", movieIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoTheseMoviesExist indicates an expected call of DoTheseMoviesExist.
func (mr *MockMovieRepositoryInterfaceMockRecorder) DoTheseMoviesExist(movieIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTheseMoviesExist", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).DoTheseMoviesExist), movieIDs)
}

// GetAll mocks base method.
func (m *MockMovieRepositoryInterface) GetAll() ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetAll))
}

// GetByActorID mocks base method.
func (m *MockMovieRepositoryInterface) GetByActorID(actorID string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByActorID", actorID)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByActorID indicates an expected call of GetByActorID.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetByActorID(actorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByActorID", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetByActorID), actorID)
}

// GetByDirectorID mocks base method.
func (m *MockMovieRepositoryInterface) GetByDirectorID(directorID string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDirectorID", directorID)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDirectorID indicates an expected call of GetByDirectorID.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetByDirectorID(directorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDirectorID", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetByDirectorID), directorID)
}

// GetByGenreID mocks base method.
func (m *MockMovieRepositoryInterface) GetByGenreID(genreID string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGenreID", genreID)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGenreID indicates an expected call of GetByGenreID.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetByGenreID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGenreID", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetByGenreID), genreID)
}

// GetByID mocks base method.
func (m *MockMovieRepositoryInterface) GetByID(movieID string) (entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", movieID)
	ret0, _ := ret[0].(entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetByID), movieID)
}

// GetByWriterID mocks base method.
func (m *MockMovieRepositoryInterface) GetByWriterID(writerID string) ([]entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByWriterID", writerID)
	ret0, _ := ret[0].([]entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByWriterID indicates an expected call of GetByWriterID.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetByWriterID(writerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByWriterID", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetByWriterID), writerID)
}

// Update mocks base method.
func (m *MockMovieRepositoryInterface) Update(movie *entity.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMovieRepositoryInterfaceMockRecorder) Update(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).Update), movie)
}