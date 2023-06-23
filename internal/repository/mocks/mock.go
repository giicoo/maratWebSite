// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	models "github.com/giicoo/maratWebSite/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockRepo) AddUser(user models.UserDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepoMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepo)(nil).AddUser), user)
}

// AddWord mocks base method.
func (m *MockRepo) AddWord(word models.WordDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWord", word)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWord indicates an expected call of AddWord.
func (mr *MockRepoMockRecorder) AddWord(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWord", reflect.TypeOf((*MockRepo)(nil).AddWord), word)
}

// GetUser mocks base method.
func (m *MockRepo) GetUser(login string) (models.UserDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", login)
	ret0, _ := ret[0].(models.UserDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepoMockRecorder) GetUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepo)(nil).GetUser), login)
}

// GetWords mocks base method.
func (m *MockRepo) GetWords() ([]*models.WordDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWords")
	ret0, _ := ret[0].([]*models.WordDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWords indicates an expected call of GetWords.
func (mr *MockRepoMockRecorder) GetWords() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWords", reflect.TypeOf((*MockRepo)(nil).GetWords))
}

// GetWordsByNames mocks base method.
func (m *MockRepo) GetWordsByNames(words []*models.WordDB) ([]*models.WordDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWordsByNames", words)
	ret0, _ := ret[0].([]*models.WordDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWordsByNames indicates an expected call of GetWordsByNames.
func (mr *MockRepoMockRecorder) GetWordsByNames(words interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWordsByNames", reflect.TypeOf((*MockRepo)(nil).GetWordsByNames), words)
}