// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock_discord is a generated GoMock package.
package mock_discord

import (
	reflect "reflect"

	discordgo "github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
)

// MockdiscordSession is a mock of discordSession interface.
type MockdiscordSession struct {
	ctrl     *gomock.Controller
	recorder *MockdiscordSessionMockRecorder
}

// MockdiscordSessionMockRecorder is the mock recorder for MockdiscordSession.
type MockdiscordSessionMockRecorder struct {
	mock *MockdiscordSession
}

// NewMockdiscordSession creates a new mock instance.
func NewMockdiscordSession(ctrl *gomock.Controller) *MockdiscordSession {
	mock := &MockdiscordSession{ctrl: ctrl}
	mock.recorder = &MockdiscordSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdiscordSession) EXPECT() *MockdiscordSessionMockRecorder {
	return m.recorder
}

// Channel mocks base method.
func (m *MockdiscordSession) Channel(arg0 string) (*discordgo.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel", arg0)
	ret0, _ := ret[0].(*discordgo.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel.
func (mr *MockdiscordSessionMockRecorder) Channel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockdiscordSession)(nil).Channel), arg0)
}

// Guild mocks base method.
func (m *MockdiscordSession) Guild(arg0 string) (*discordgo.Guild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Guild", arg0)
	ret0, _ := ret[0].(*discordgo.Guild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Guild indicates an expected call of Guild.
func (mr *MockdiscordSessionMockRecorder) Guild(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Guild", reflect.TypeOf((*MockdiscordSession)(nil).Guild), arg0)
}

// GuildMembers mocks base method.
func (m *MockdiscordSession) GuildMembers(arg0, arg1 string, arg2 int) ([]*discordgo.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GuildMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*discordgo.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GuildMembers indicates an expected call of GuildMembers.
func (mr *MockdiscordSessionMockRecorder) GuildMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GuildMembers", reflect.TypeOf((*MockdiscordSession)(nil).GuildMembers), arg0, arg1, arg2)
}

// InteractionRespond mocks base method.
func (m *MockdiscordSession) InteractionRespond(arg0 *discordgo.Interaction, arg1 *discordgo.InteractionResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InteractionRespond", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InteractionRespond indicates an expected call of InteractionRespond.
func (mr *MockdiscordSessionMockRecorder) InteractionRespond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InteractionRespond", reflect.TypeOf((*MockdiscordSession)(nil).InteractionRespond), arg0, arg1)
}
