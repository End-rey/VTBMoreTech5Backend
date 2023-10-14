// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTranslation is a mock of Translation interface.
type MockTranslation struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationMockRecorder
}

// MockTranslationMockRecorder is the mock recorder for MockTranslation.
type MockTranslationMockRecorder struct {
	mock *MockTranslation
}

// NewMockTranslation creates a new mock instance.
func NewMockTranslation(ctrl *gomock.Controller) *MockTranslation {
	mock := &MockTranslation{ctrl: ctrl}
	mock.recorder = &MockTranslationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslation) EXPECT() *MockTranslationMockRecorder {
	return m.recorder
}

// History mocks base method.
func (m *MockTranslation) History(arg0 context.Context) ([]entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "History", arg0)
	ret0, _ := ret[0].([]entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// History indicates an expected call of History.
func (mr *MockTranslationMockRecorder) History(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "History", reflect.TypeOf((*MockTranslation)(nil).History), arg0)
}

// Translate mocks base method.
func (m *MockTranslation) Translate(arg0 context.Context, arg1 entity.Translation) (entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0, arg1)
	ret0, _ := ret[0].(entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockTranslationMockRecorder) Translate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslation)(nil).Translate), arg0, arg1)
}

// MockTranslationRepo is a mock of TranslationRepo interface.
type MockTranslationRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationRepoMockRecorder
}

// MockTranslationRepoMockRecorder is the mock recorder for MockTranslationRepo.
type MockTranslationRepoMockRecorder struct {
	mock *MockTranslationRepo
}

// NewMockTranslationRepo creates a new mock instance.
func NewMockTranslationRepo(ctrl *gomock.Controller) *MockTranslationRepo {
	mock := &MockTranslationRepo{ctrl: ctrl}
	mock.recorder = &MockTranslationRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationRepo) EXPECT() *MockTranslationRepoMockRecorder {
	return m.recorder
}

// GetHistory mocks base method.
func (m *MockTranslationRepo) GetHistory(arg0 context.Context) ([]entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", arg0)
	ret0, _ := ret[0].([]entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockTranslationRepoMockRecorder) GetHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockTranslationRepo)(nil).GetHistory), arg0)
}

// Store mocks base method.
func (m *MockTranslationRepo) Store(arg0 context.Context, arg1 entity.Translation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTranslationRepoMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTranslationRepo)(nil).Store), arg0, arg1)
}

// MockTranslationWebAPI is a mock of TranslationWebAPI interface.
type MockTranslationWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationWebAPIMockRecorder
}

// MockTranslationWebAPIMockRecorder is the mock recorder for MockTranslationWebAPI.
type MockTranslationWebAPIMockRecorder struct {
	mock *MockTranslationWebAPI
}

// NewMockTranslationWebAPI creates a new mock instance.
func NewMockTranslationWebAPI(ctrl *gomock.Controller) *MockTranslationWebAPI {
	mock := &MockTranslationWebAPI{ctrl: ctrl}
	mock.recorder = &MockTranslationWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationWebAPI) EXPECT() *MockTranslationWebAPIMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *MockTranslationWebAPI) Translate(arg0 entity.Translation) (entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockTranslationWebAPIMockRecorder) Translate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslationWebAPI)(nil).Translate), arg0)
}
