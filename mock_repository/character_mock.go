// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	repository "github.com/HETIC-MT-P2021/RPGo/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCharacterRepositoryInterface is a mock of CharacterRepositoryInterface interface
type MockCharacterRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterRepositoryInterfaceMockRecorder
}

// MockCharacterRepositoryInterfaceMockRecorder is the mock recorder for MockCharacterRepositoryInterface
type MockCharacterRepositoryInterfaceMockRecorder struct {
	mock *MockCharacterRepositoryInterface
}

// NewMockCharacterRepositoryInterface creates a new mock instance
func NewMockCharacterRepositoryInterface(ctrl *gomock.Controller) *MockCharacterRepositoryInterface {
	mock := &MockCharacterRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCharacterRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterRepositoryInterface) EXPECT() *MockCharacterRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCharacterRepositoryInterface) Create(character *repository.Character) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", character)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCharacterRepositoryInterfaceMockRecorder) Create(character interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockCharacterRepositoryInterface)(nil).Create), character)
}

// GetCharacterByDiscordUserID mocks base method
func (m *MockCharacterRepositoryInterface) GetCharacterByDiscordUserID(discordUserID string) (*repository.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharacterByDiscordUserID", discordUserID)
	ret0, _ := ret[0].(*repository.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharacterByDiscordUserID indicates an expected call of GetCharacterByDiscordUserID
func (mr *MockCharacterRepositoryInterfaceMockRecorder) GetCharacterByDiscordUserID(discordUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharacterByDiscordUserID", reflect.TypeOf((*MockCharacterRepositoryInterface)(nil).GetCharacterByDiscordUserID), discordUserID)
}
