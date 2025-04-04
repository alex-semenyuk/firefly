// Code generated by mockery v2.46.0. DO NOT EDIT.

package eventsmocks

import (
	context "context"

	config "github.com/hyperledger/firefly-common/pkg/config"

	core "github.com/hyperledger/firefly/pkg/core"

	events "github.com/hyperledger/firefly/pkg/events"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// BatchDeliveryRequest provides a mock function with given fields: ctx, connID, sub, _a3
func (_m *Plugin) BatchDeliveryRequest(ctx context.Context, connID string, sub *core.Subscription, _a3 []*core.CombinedEventDataDelivery) error {
	ret := _m.Called(ctx, connID, sub, _a3)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeliveryRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.Subscription, []*core.CombinedEventDataDelivery) error); ok {
		r0 = rf(ctx, connID, sub, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *events.Capabilities {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Capabilities")
	}

	var r0 *events.Capabilities
	if rf, ok := ret.Get(0).(func() *events.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*events.Capabilities)
		}
	}

	return r0
}

// DeliveryRequest provides a mock function with given fields: ctx, connID, sub, event, data
func (_m *Plugin) DeliveryRequest(ctx context.Context, connID string, sub *core.Subscription, event *core.EventDelivery, data core.DataArray) error {
	ret := _m.Called(ctx, connID, sub, event, data)

	if len(ret) == 0 {
		panic("no return value specified for DeliveryRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.Subscription, *core.EventDelivery, core.DataArray) error); ok {
		r0 = rf(ctx, connID, sub, event, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields: ctx, _a1
func (_m *Plugin) Init(ctx context.Context, _a1 config.Section) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Section) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitConfig provides a mock function with given fields: _a0
func (_m *Plugin) InitConfig(_a0 config.Section) {
	_m.Called(_a0)
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NamespaceRestarted provides a mock function with given fields: ns, startTime
func (_m *Plugin) NamespaceRestarted(ns string, startTime time.Time) {
	_m.Called(ns, startTime)
}

// SetHandler provides a mock function with given fields: namespace, handler
func (_m *Plugin) SetHandler(namespace string, handler events.Callbacks) error {
	ret := _m.Called(namespace, handler)

	if len(ret) == 0 {
		panic("no return value specified for SetHandler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.Callbacks) error); ok {
		r0 = rf(namespace, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateOptions provides a mock function with given fields: ctx, options
func (_m *Plugin) ValidateOptions(ctx context.Context, options *core.SubscriptionOptions) error {
	ret := _m.Called(ctx, options)

	if len(ret) == 0 {
		panic("no return value specified for ValidateOptions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.SubscriptionOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
