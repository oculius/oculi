// Code generated by MockGen. DO NOT EDIT.
// Source: request/request.go

// Package mock_request is a generated GoMock package.
package mock_request

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	auth "github.com/ravielze/oculi/common/model/dto/auth"
	sql "github.com/ravielze/oculi/persistent/sql"
	request "github.com/ravielze/oculi/request"
)

// MockReqContext is a mock of ReqContext interface.
type MockReqContext struct {
	ctrl     *gomock.Controller
	recorder *MockReqContextMockRecorder
}

// MockReqContextMockRecorder is the mock recorder for MockReqContext.
type MockReqContextMockRecorder struct {
	mock *MockReqContext
}

// NewMockReqContext creates a new mock instance.
func NewMockReqContext(ctrl *gomock.Controller) *MockReqContext {
	mock := &MockReqContext{ctrl: ctrl}
	mock.recorder = &MockReqContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReqContext) EXPECT() *MockReqContextMockRecorder {
	return m.recorder
}

// AddError mocks base method.
func (m *MockReqContext) AddError(responseCode int, err ...error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{responseCode}
	for _, a := range err {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddError", varargs...)
}

// AddError indicates an expected call of AddError.
func (mr *MockReqContextMockRecorder) AddError(responseCode interface{}, err ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{responseCode}, err...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*MockReqContext)(nil).AddError), varargs...)
}

// AfterCommitDo mocks base method.
func (m *MockReqContext) AfterCommitDo(f func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterCommitDo", f)
}

// AfterCommitDo indicates an expected call of AfterCommitDo.
func (mr *MockReqContextMockRecorder) AfterCommitDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterCommitDo", reflect.TypeOf((*MockReqContext)(nil).AfterCommitDo), f)
}

// AfterRollbackDo mocks base method.
func (m *MockReqContext) AfterRollbackDo(f func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterRollbackDo", f)
}

// AfterRollbackDo indicates an expected call of AfterRollbackDo.
func (mr *MockReqContextMockRecorder) AfterRollbackDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterRollbackDo", reflect.TypeOf((*MockReqContext)(nil).AfterRollbackDo), f)
}

// BeforeCommitDo mocks base method.
func (m *MockReqContext) BeforeCommitDo(f func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeCommitDo", f)
}

// BeforeCommitDo indicates an expected call of BeforeCommitDo.
func (mr *MockReqContextMockRecorder) BeforeCommitDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeCommitDo", reflect.TypeOf((*MockReqContext)(nil).BeforeCommitDo), f)
}

// BeforeRollbackDo mocks base method.
func (m *MockReqContext) BeforeRollbackDo(f func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeRollbackDo", f)
}

// BeforeRollbackDo indicates an expected call of BeforeRollbackDo.
func (mr *MockReqContextMockRecorder) BeforeRollbackDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeRollbackDo", reflect.TypeOf((*MockReqContext)(nil).BeforeRollbackDo), f)
}

// CommitTransaction mocks base method.
func (m *MockReqContext) CommitTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockReqContextMockRecorder) CommitTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockReqContext)(nil).CommitTransaction))
}

// Context mocks base method.
func (m *MockReqContext) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockReqContextMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockReqContext)(nil).Context))
}

// Error mocks base method.
func (m *MockReqContext) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockReqContextMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockReqContext)(nil).Error))
}

// Get mocks base method.
func (m *MockReqContext) Get(key string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReqContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReqContext)(nil).Get), key)
}

// GetOrDefault mocks base method.
func (m *MockReqContext) GetOrDefault(key string, def interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrDefault", key, def)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetOrDefault indicates an expected call of GetOrDefault.
func (mr *MockReqContextMockRecorder) GetOrDefault(key, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrDefault", reflect.TypeOf((*MockReqContext)(nil).GetOrDefault), key, def)
}

// HasError mocks base method.
func (m *MockReqContext) HasError() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasError")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasError indicates an expected call of HasError.
func (mr *MockReqContextMockRecorder) HasError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasError", reflect.TypeOf((*MockReqContext)(nil).HasError))
}

// HasTransaction mocks base method.
func (m *MockReqContext) HasTransaction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTransaction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTransaction indicates an expected call of HasTransaction.
func (mr *MockReqContextMockRecorder) HasTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTransaction", reflect.TypeOf((*MockReqContext)(nil).HasTransaction))
}

// Identifier mocks base method.
func (m *MockReqContext) Identifier() auth.StandardCredentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(auth.StandardCredentials)
	return ret0
}

// Identifier indicates an expected call of Identifier.
func (mr *MockReqContextMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockReqContext)(nil).Identifier))
}

// NewTransaction mocks base method.
func (m *MockReqContext) NewTransaction() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewTransaction")
}

// NewTransaction indicates an expected call of NewTransaction.
func (mr *MockReqContextMockRecorder) NewTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTransaction", reflect.TypeOf((*MockReqContext)(nil).NewTransaction))
}

// Parse36 mocks base method.
func (m *MockReqContext) Parse36(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse36", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// Parse36 indicates an expected call of Parse36.
func (mr *MockReqContextMockRecorder) Parse36(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse36", reflect.TypeOf((*MockReqContext)(nil).Parse36), key, value)
}

// Parse36UUID mocks base method.
func (m *MockReqContext) Parse36UUID(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse36UUID", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// Parse36UUID indicates an expected call of Parse36UUID.
func (mr *MockReqContextMockRecorder) Parse36UUID(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse36UUID", reflect.TypeOf((*MockReqContext)(nil).Parse36UUID), key, value)
}

// ParseBoolean mocks base method.
func (m *MockReqContext) ParseBoolean(key, value string, def bool) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBoolean", key, value, def)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseBoolean indicates an expected call of ParseBoolean.
func (mr *MockReqContextMockRecorder) ParseBoolean(key, value, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBoolean", reflect.TypeOf((*MockReqContext)(nil).ParseBoolean), key, value, def)
}

// ParseString mocks base method.
func (m *MockReqContext) ParseString(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseString", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseString indicates an expected call of ParseString.
func (mr *MockReqContextMockRecorder) ParseString(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseString", reflect.TypeOf((*MockReqContext)(nil).ParseString), key, value)
}

// ParseStringOrDefault mocks base method.
func (m *MockReqContext) ParseStringOrDefault(key, value, def string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseStringOrDefault", key, value, def)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseStringOrDefault indicates an expected call of ParseStringOrDefault.
func (mr *MockReqContextMockRecorder) ParseStringOrDefault(key, value, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseStringOrDefault", reflect.TypeOf((*MockReqContext)(nil).ParseStringOrDefault), key, value, def)
}

// ParseUUID mocks base method.
func (m *MockReqContext) ParseUUID(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUUID", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseUUID indicates an expected call of ParseUUID.
func (mr *MockReqContextMockRecorder) ParseUUID(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUUID", reflect.TypeOf((*MockReqContext)(nil).ParseUUID), key, value)
}

// ParseUUID36 mocks base method.
func (m *MockReqContext) ParseUUID36(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUUID36", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseUUID36 indicates an expected call of ParseUUID36.
func (mr *MockReqContextMockRecorder) ParseUUID36(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUUID36", reflect.TypeOf((*MockReqContext)(nil).ParseUUID36), key, value)
}

// ResponseCode mocks base method.
func (m *MockReqContext) ResponseCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// ResponseCode indicates an expected call of ResponseCode.
func (mr *MockReqContextMockRecorder) ResponseCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseCode", reflect.TypeOf((*MockReqContext)(nil).ResponseCode))
}

// RollbackTransaction mocks base method.
func (m *MockReqContext) RollbackTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockReqContextMockRecorder) RollbackTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockReqContext)(nil).RollbackTransaction))
}

// Set mocks base method.
func (m *MockReqContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set.
func (mr *MockReqContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockReqContext)(nil).Set), key, val)
}

// SetResponseCode mocks base method.
func (m *MockReqContext) SetResponseCode(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResponseCode", code)
}

// SetResponseCode indicates an expected call of SetResponseCode.
func (mr *MockReqContextMockRecorder) SetResponseCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponseCode", reflect.TypeOf((*MockReqContext)(nil).SetResponseCode), code)
}

// Transaction mocks base method.
func (m *MockReqContext) Transaction() sql.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(sql.API)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockReqContextMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockReqContext)(nil).Transaction))
}

// WithContext mocks base method.
func (m *MockReqContext) WithContext(ctx context.Context) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockReqContextMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockReqContext)(nil).WithContext), ctx)
}

// WithIdentifier mocks base method.
func (m *MockReqContext) WithIdentifier(id auth.StandardCredentials) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithIdentifier", id)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// WithIdentifier indicates an expected call of WithIdentifier.
func (mr *MockReqContextMockRecorder) WithIdentifier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithIdentifier", reflect.TypeOf((*MockReqContext)(nil).WithIdentifier), id)
}

// MockNonEchoContext is a mock of NonEchoContext interface.
type MockNonEchoContext struct {
	ctrl     *gomock.Controller
	recorder *MockNonEchoContextMockRecorder
}

// MockNonEchoContextMockRecorder is the mock recorder for MockNonEchoContext.
type MockNonEchoContextMockRecorder struct {
	mock *MockNonEchoContext
}

// NewMockNonEchoContext creates a new mock instance.
func NewMockNonEchoContext(ctrl *gomock.Controller) *MockNonEchoContext {
	mock := &MockNonEchoContext{ctrl: ctrl}
	mock.recorder = &MockNonEchoContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNonEchoContext) EXPECT() *MockNonEchoContextMockRecorder {
	return m.recorder
}

// BindValidate mocks base method.
func (m *MockNonEchoContext) BindValidate(obj interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindValidate", obj)
}

// BindValidate indicates an expected call of BindValidate.
func (mr *MockNonEchoContextMockRecorder) BindValidate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindValidate", reflect.TypeOf((*MockNonEchoContext)(nil).BindValidate), obj)
}

// MockEchoReqContext is a mock of EchoReqContext interface.
type MockEchoReqContext struct {
	ctrl     *gomock.Controller
	recorder *MockEchoReqContextMockRecorder
}

// MockEchoReqContextMockRecorder is the mock recorder for MockEchoReqContext.
type MockEchoReqContextMockRecorder struct {
	mock *MockEchoReqContext
}

// NewMockEchoReqContext creates a new mock instance.
func NewMockEchoReqContext(ctrl *gomock.Controller) *MockEchoReqContext {
	mock := &MockEchoReqContext{ctrl: ctrl}
	mock.recorder = &MockEchoReqContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEchoReqContext) EXPECT() *MockEchoReqContextMockRecorder {
	return m.recorder
}

// AddError mocks base method.
func (m *MockEchoReqContext) AddError(responseCode int, err ...error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{responseCode}
	for _, a := range err {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddError", varargs...)
}

// AddError indicates an expected call of AddError.
func (mr *MockEchoReqContextMockRecorder) AddError(responseCode interface{}, err ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{responseCode}, err...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*MockEchoReqContext)(nil).AddError), varargs...)
}

// AfterCommitDo mocks base method.
func (m *MockEchoReqContext) AfterCommitDo(f func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterCommitDo", f)
}

// AfterCommitDo indicates an expected call of AfterCommitDo.
func (mr *MockEchoReqContextMockRecorder) AfterCommitDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterCommitDo", reflect.TypeOf((*MockEchoReqContext)(nil).AfterCommitDo), f)
}

// AfterRollbackDo mocks base method.
func (m *MockEchoReqContext) AfterRollbackDo(f func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterRollbackDo", f)
}

// AfterRollbackDo indicates an expected call of AfterRollbackDo.
func (mr *MockEchoReqContextMockRecorder) AfterRollbackDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterRollbackDo", reflect.TypeOf((*MockEchoReqContext)(nil).AfterRollbackDo), f)
}

// BeforeCommitDo mocks base method.
func (m *MockEchoReqContext) BeforeCommitDo(f func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeCommitDo", f)
}

// BeforeCommitDo indicates an expected call of BeforeCommitDo.
func (mr *MockEchoReqContextMockRecorder) BeforeCommitDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeCommitDo", reflect.TypeOf((*MockEchoReqContext)(nil).BeforeCommitDo), f)
}

// BeforeRollbackDo mocks base method.
func (m *MockEchoReqContext) BeforeRollbackDo(f func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeRollbackDo", f)
}

// BeforeRollbackDo indicates an expected call of BeforeRollbackDo.
func (mr *MockEchoReqContextMockRecorder) BeforeRollbackDo(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeRollbackDo", reflect.TypeOf((*MockEchoReqContext)(nil).BeforeRollbackDo), f)
}

// CommitTransaction mocks base method.
func (m *MockEchoReqContext) CommitTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockEchoReqContextMockRecorder) CommitTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockEchoReqContext)(nil).CommitTransaction))
}

// Context mocks base method.
func (m *MockEchoReqContext) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEchoReqContextMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEchoReqContext)(nil).Context))
}

// Echo mocks base method.
func (m *MockEchoReqContext) Echo() echo.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo")
	ret0, _ := ret[0].(echo.Context)
	return ret0
}

// Echo indicates an expected call of Echo.
func (mr *MockEchoReqContextMockRecorder) Echo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockEchoReqContext)(nil).Echo))
}

// Error mocks base method.
func (m *MockEchoReqContext) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockEchoReqContextMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockEchoReqContext)(nil).Error))
}

// Get mocks base method.
func (m *MockEchoReqContext) Get(key string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEchoReqContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEchoReqContext)(nil).Get), key)
}

// GetOrDefault mocks base method.
func (m *MockEchoReqContext) GetOrDefault(key string, def interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrDefault", key, def)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetOrDefault indicates an expected call of GetOrDefault.
func (mr *MockEchoReqContextMockRecorder) GetOrDefault(key, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrDefault", reflect.TypeOf((*MockEchoReqContext)(nil).GetOrDefault), key, def)
}

// HasError mocks base method.
func (m *MockEchoReqContext) HasError() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasError")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasError indicates an expected call of HasError.
func (mr *MockEchoReqContextMockRecorder) HasError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasError", reflect.TypeOf((*MockEchoReqContext)(nil).HasError))
}

// HasTransaction mocks base method.
func (m *MockEchoReqContext) HasTransaction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTransaction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTransaction indicates an expected call of HasTransaction.
func (mr *MockEchoReqContextMockRecorder) HasTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTransaction", reflect.TypeOf((*MockEchoReqContext)(nil).HasTransaction))
}

// Identifier mocks base method.
func (m *MockEchoReqContext) Identifier() auth.StandardCredentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(auth.StandardCredentials)
	return ret0
}

// Identifier indicates an expected call of Identifier.
func (mr *MockEchoReqContextMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockEchoReqContext)(nil).Identifier))
}

// NewTransaction mocks base method.
func (m *MockEchoReqContext) NewTransaction() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewTransaction")
}

// NewTransaction indicates an expected call of NewTransaction.
func (mr *MockEchoReqContextMockRecorder) NewTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTransaction", reflect.TypeOf((*MockEchoReqContext)(nil).NewTransaction))
}

// Param mocks base method.
func (m *MockEchoReqContext) Param(param string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", param)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockEchoReqContextMockRecorder) Param(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockEchoReqContext)(nil).Param), param)
}

// Param36 mocks base method.
func (m *MockEchoReqContext) Param36(param string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param36", param)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// Param36 indicates an expected call of Param36.
func (mr *MockEchoReqContextMockRecorder) Param36(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param36", reflect.TypeOf((*MockEchoReqContext)(nil).Param36), param)
}

// Param36UUID mocks base method.
func (m *MockEchoReqContext) Param36UUID(param string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param36UUID", param)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// Param36UUID indicates an expected call of Param36UUID.
func (mr *MockEchoReqContextMockRecorder) Param36UUID(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param36UUID", reflect.TypeOf((*MockEchoReqContext)(nil).Param36UUID), param)
}

// ParamUUID mocks base method.
func (m *MockEchoReqContext) ParamUUID(param string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamUUID", param)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// ParamUUID indicates an expected call of ParamUUID.
func (mr *MockEchoReqContextMockRecorder) ParamUUID(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamUUID", reflect.TypeOf((*MockEchoReqContext)(nil).ParamUUID), param)
}

// ParamUUID36 mocks base method.
func (m *MockEchoReqContext) ParamUUID36(param string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamUUID36", param)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// ParamUUID36 indicates an expected call of ParamUUID36.
func (mr *MockEchoReqContextMockRecorder) ParamUUID36(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamUUID36", reflect.TypeOf((*MockEchoReqContext)(nil).ParamUUID36), param)
}

// Parse36 mocks base method.
func (m *MockEchoReqContext) Parse36(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse36", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// Parse36 indicates an expected call of Parse36.
func (mr *MockEchoReqContextMockRecorder) Parse36(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse36", reflect.TypeOf((*MockEchoReqContext)(nil).Parse36), key, value)
}

// Parse36UUID mocks base method.
func (m *MockEchoReqContext) Parse36UUID(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse36UUID", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// Parse36UUID indicates an expected call of Parse36UUID.
func (mr *MockEchoReqContextMockRecorder) Parse36UUID(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse36UUID", reflect.TypeOf((*MockEchoReqContext)(nil).Parse36UUID), key, value)
}

// ParseBoolean mocks base method.
func (m *MockEchoReqContext) ParseBoolean(key, value string, def bool) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBoolean", key, value, def)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseBoolean indicates an expected call of ParseBoolean.
func (mr *MockEchoReqContextMockRecorder) ParseBoolean(key, value, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBoolean", reflect.TypeOf((*MockEchoReqContext)(nil).ParseBoolean), key, value, def)
}

// ParseString mocks base method.
func (m *MockEchoReqContext) ParseString(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseString", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseString indicates an expected call of ParseString.
func (mr *MockEchoReqContextMockRecorder) ParseString(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseString", reflect.TypeOf((*MockEchoReqContext)(nil).ParseString), key, value)
}

// ParseStringOrDefault mocks base method.
func (m *MockEchoReqContext) ParseStringOrDefault(key, value, def string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseStringOrDefault", key, value, def)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseStringOrDefault indicates an expected call of ParseStringOrDefault.
func (mr *MockEchoReqContextMockRecorder) ParseStringOrDefault(key, value, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseStringOrDefault", reflect.TypeOf((*MockEchoReqContext)(nil).ParseStringOrDefault), key, value, def)
}

// ParseUUID mocks base method.
func (m *MockEchoReqContext) ParseUUID(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUUID", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseUUID indicates an expected call of ParseUUID.
func (mr *MockEchoReqContextMockRecorder) ParseUUID(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUUID", reflect.TypeOf((*MockEchoReqContext)(nil).ParseUUID), key, value)
}

// ParseUUID36 mocks base method.
func (m *MockEchoReqContext) ParseUUID36(key, value string) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUUID36", key, value)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// ParseUUID36 indicates an expected call of ParseUUID36.
func (mr *MockEchoReqContextMockRecorder) ParseUUID36(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUUID36", reflect.TypeOf((*MockEchoReqContext)(nil).ParseUUID36), key, value)
}

// Query mocks base method.
func (m *MockEchoReqContext) Query(query, def string) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", query, def)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockEchoReqContextMockRecorder) Query(query, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockEchoReqContext)(nil).Query), query, def)
}

// QueryBoolean mocks base method.
func (m *MockEchoReqContext) QueryBoolean(query string, def bool) request.EchoReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBoolean", query, def)
	ret0, _ := ret[0].(request.EchoReqContext)
	return ret0
}

// QueryBoolean indicates an expected call of QueryBoolean.
func (mr *MockEchoReqContextMockRecorder) QueryBoolean(query, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBoolean", reflect.TypeOf((*MockEchoReqContext)(nil).QueryBoolean), query, def)
}

// ResponseCode mocks base method.
func (m *MockEchoReqContext) ResponseCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// ResponseCode indicates an expected call of ResponseCode.
func (mr *MockEchoReqContextMockRecorder) ResponseCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseCode", reflect.TypeOf((*MockEchoReqContext)(nil).ResponseCode))
}

// RollbackTransaction mocks base method.
func (m *MockEchoReqContext) RollbackTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockEchoReqContextMockRecorder) RollbackTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockEchoReqContext)(nil).RollbackTransaction))
}

// Set mocks base method.
func (m *MockEchoReqContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set.
func (mr *MockEchoReqContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockEchoReqContext)(nil).Set), key, val)
}

// SetResponseCode mocks base method.
func (m *MockEchoReqContext) SetResponseCode(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResponseCode", code)
}

// SetResponseCode indicates an expected call of SetResponseCode.
func (mr *MockEchoReqContextMockRecorder) SetResponseCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponseCode", reflect.TypeOf((*MockEchoReqContext)(nil).SetResponseCode), code)
}

// Transaction mocks base method.
func (m *MockEchoReqContext) Transaction() sql.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(sql.API)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockEchoReqContextMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockEchoReqContext)(nil).Transaction))
}

// Transform mocks base method.
func (m *MockEchoReqContext) Transform() request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform")
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// Transform indicates an expected call of Transform.
func (mr *MockEchoReqContextMockRecorder) Transform() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockEchoReqContext)(nil).Transform))
}

// WithContext mocks base method.
func (m *MockEchoReqContext) WithContext(ctx context.Context) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockEchoReqContextMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockEchoReqContext)(nil).WithContext), ctx)
}

// WithIdentifier mocks base method.
func (m *MockEchoReqContext) WithIdentifier(id auth.StandardCredentials) request.ReqContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithIdentifier", id)
	ret0, _ := ret[0].(request.ReqContext)
	return ret0
}

// WithIdentifier indicates an expected call of WithIdentifier.
func (mr *MockEchoReqContextMockRecorder) WithIdentifier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithIdentifier", reflect.TypeOf((*MockEchoReqContext)(nil).WithIdentifier), id)
}
