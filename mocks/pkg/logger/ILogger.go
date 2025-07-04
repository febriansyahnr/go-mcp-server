// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks_logger

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	slack "github.com/bluele/slack"

	zap "go.uber.org/zap"

	zapcore "go.uber.org/zap/zapcore"
)

// ILogger is an autogenerated mock type for the ILogger type
type ILogger struct {
	mock.Mock
}

// CleanupSlackLogger provides a mock function with no fields
func (_m *ILogger) CleanupSlackLogger() {
	_m.Called()
}

// CustomSlackAlertFinopsNotification provides a mock function with given fields: attachment
func (_m *ILogger) CustomSlackAlertFinopsNotification(attachment slack.Attachment) {
	_m.Called(attachment)
}

// CustomSlackNotification provides a mock function with given fields: attachment
func (_m *ILogger) CustomSlackNotification(attachment slack.Attachment) {
	_m.Called(attachment)
}

// Debug provides a mock function with given fields: ctx, msg, fields
func (_m *ILogger) Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: ctx, msg, fields
func (_m *ILogger) Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// GetLogger provides a mock function with no fields
func (_m *ILogger) GetLogger() *zap.Logger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLogger")
	}

	var r0 *zap.Logger
	if rf, ok := ret.Get(0).(func() *zap.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zap.Logger)
		}
	}

	return r0
}

// Info provides a mock function with given fields: ctx, msg, fields
func (_m *ILogger) Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Panic provides a mock function with given fields: ctx, msg, fields
func (_m *ILogger) Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Sync provides a mock function with no fields
func (_m *ILogger) Sync() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Sync")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warn provides a mock function with given fields: ctx, msg, fields
func (_m *ILogger) Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NewILogger creates a new instance of ILogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILogger {
	mock := &ILogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
