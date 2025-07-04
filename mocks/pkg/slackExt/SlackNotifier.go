// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks_slackExt

import (
	context "context"

	slackExt "github.com/paper-indonesia/pg-mcp-server/pkg/slackExt"
	mock "github.com/stretchr/testify/mock"
)

// SlackNotifier is an autogenerated mock type for the SlackNotifier type
type SlackNotifier struct {
	mock.Mock
}

// PostWebhook provides a mock function with given fields: ctx, cmd
func (_m *SlackNotifier) PostWebhook(ctx context.Context, cmd *slackExt.PostWebhookCmd) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for PostWebhook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *slackExt.PostWebhookCmd) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSlackNotifier creates a new instance of SlackNotifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlackNotifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *SlackNotifier {
	mock := &SlackNotifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
