// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	command "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/command"
	coredata "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/coredata"

	dtos "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"

	logger "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	mock "github.com/stretchr/testify/mock"

	notifications "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/notifications"

	time "time"
)

// AppFunctionContext is an autogenerated mock type for the AppFunctionContext type
type AppFunctionContext struct {
	mock.Mock
}

// CommandClient provides a mock function with given fields:
func (_m *AppFunctionContext) CommandClient() command.CommandClient {
	ret := _m.Called()

	var r0 command.CommandClient
	if rf, ok := ret.Get(0).(func() command.CommandClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(command.CommandClient)
		}
	}

	return r0
}

// CorrelationID provides a mock function with given fields:
func (_m *AppFunctionContext) CorrelationID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EventClient provides a mock function with given fields:
func (_m *AppFunctionContext) EventClient() coredata.EventClient {
	ret := _m.Called()

	var r0 coredata.EventClient
	if rf, ok := ret.Get(0).(func() coredata.EventClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coredata.EventClient)
		}
	}

	return r0
}

// GetSecret provides a mock function with given fields: path, keys
func (_m *AppFunctionContext) GetSecret(path string, keys ...string) (map[string]string, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, ...string) map[string]string); ok {
		r0 = rf(path, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(path, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InputContentType provides a mock function with given fields:
func (_m *AppFunctionContext) InputContentType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LoggingClient provides a mock function with given fields:
func (_m *AppFunctionContext) LoggingClient() logger.LoggingClient {
	ret := _m.Called()

	var r0 logger.LoggingClient
	if rf, ok := ret.Get(0).(func() logger.LoggingClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.LoggingClient)
		}
	}

	return r0
}

// NotificationsClient provides a mock function with given fields:
func (_m *AppFunctionContext) NotificationsClient() notifications.NotificationsClient {
	ret := _m.Called()

	var r0 notifications.NotificationsClient
	if rf, ok := ret.Get(0).(func() notifications.NotificationsClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(notifications.NotificationsClient)
		}
	}

	return r0
}

// PushToCoreData provides a mock function with given fields: deviceName, readingName, value
func (_m *AppFunctionContext) PushToCoreData(deviceName string, readingName string, value interface{}) (*dtos.Event, error) {
	ret := _m.Called(deviceName, readingName, value)

	var r0 *dtos.Event
	if rf, ok := ret.Get(0).(func(string, string, interface{}) *dtos.Event); ok {
		r0 = rf(deviceName, readingName, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}) error); ok {
		r1 = rf(deviceName, readingName, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResponseContentType provides a mock function with given fields:
func (_m *AppFunctionContext) ResponseContentType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResponseData provides a mock function with given fields:
func (_m *AppFunctionContext) ResponseData() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// SecretsLastUpdated provides a mock function with given fields:
func (_m *AppFunctionContext) SecretsLastUpdated() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// SetResponseContentType provides a mock function with given fields: _a0
func (_m *AppFunctionContext) SetResponseContentType(_a0 string) {
	_m.Called(_a0)
}

// SetResponseData provides a mock function with given fields: data
func (_m *AppFunctionContext) SetResponseData(data []byte) {
	_m.Called(data)
}

// SetRetryData provides a mock function with given fields: data
func (_m *AppFunctionContext) SetRetryData(data []byte) {
	_m.Called(data)
}
