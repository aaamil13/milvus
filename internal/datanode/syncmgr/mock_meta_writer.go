// Code generated by mockery v2.32.4. DO NOT EDIT.

package syncmgr

import mock "github.com/stretchr/testify/mock"

// MockMetaWriter is an autogenerated mock type for the MetaWriter type
type MockMetaWriter struct {
	mock.Mock
}

type MockMetaWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetaWriter) EXPECT() *MockMetaWriter_Expecter {
	return &MockMetaWriter_Expecter{mock: &_m.Mock}
}

// DropChannel provides a mock function with given fields: _a0
func (_m *MockMetaWriter) DropChannel(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetaWriter_DropChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropChannel'
type MockMetaWriter_DropChannel_Call struct {
	*mock.Call
}

// DropChannel is a helper method to define mock.On call
//   - _a0 string
func (_e *MockMetaWriter_Expecter) DropChannel(_a0 interface{}) *MockMetaWriter_DropChannel_Call {
	return &MockMetaWriter_DropChannel_Call{Call: _e.mock.On("DropChannel", _a0)}
}

func (_c *MockMetaWriter_DropChannel_Call) Run(run func(_a0 string)) *MockMetaWriter_DropChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockMetaWriter_DropChannel_Call) Return(_a0 error) *MockMetaWriter_DropChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaWriter_DropChannel_Call) RunAndReturn(run func(string) error) *MockMetaWriter_DropChannel_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSync provides a mock function with given fields: _a0
func (_m *MockMetaWriter) UpdateSync(_a0 *SyncTask) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*SyncTask) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetaWriter_UpdateSync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSync'
type MockMetaWriter_UpdateSync_Call struct {
	*mock.Call
}

// UpdateSync is a helper method to define mock.On call
//   - _a0 *SyncTask
func (_e *MockMetaWriter_Expecter) UpdateSync(_a0 interface{}) *MockMetaWriter_UpdateSync_Call {
	return &MockMetaWriter_UpdateSync_Call{Call: _e.mock.On("UpdateSync", _a0)}
}

func (_c *MockMetaWriter_UpdateSync_Call) Run(run func(_a0 *SyncTask)) *MockMetaWriter_UpdateSync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*SyncTask))
	})
	return _c
}

func (_c *MockMetaWriter_UpdateSync_Call) Return(_a0 error) *MockMetaWriter_UpdateSync_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaWriter_UpdateSync_Call) RunAndReturn(run func(*SyncTask) error) *MockMetaWriter_UpdateSync_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSyncV2 provides a mock function with given fields: _a0
func (_m *MockMetaWriter) UpdateSyncV2(_a0 *SyncTaskV2) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*SyncTaskV2) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetaWriter_UpdateSyncV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSyncV2'
type MockMetaWriter_UpdateSyncV2_Call struct {
	*mock.Call
}

// UpdateSyncV2 is a helper method to define mock.On call
//   - _a0 *SyncTaskV2
func (_e *MockMetaWriter_Expecter) UpdateSyncV2(_a0 interface{}) *MockMetaWriter_UpdateSyncV2_Call {
	return &MockMetaWriter_UpdateSyncV2_Call{Call: _e.mock.On("UpdateSyncV2", _a0)}
}

func (_c *MockMetaWriter_UpdateSyncV2_Call) Run(run func(_a0 *SyncTaskV2)) *MockMetaWriter_UpdateSyncV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*SyncTaskV2))
	})
	return _c
}

func (_c *MockMetaWriter_UpdateSyncV2_Call) Return(_a0 error) *MockMetaWriter_UpdateSyncV2_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaWriter_UpdateSyncV2_Call) RunAndReturn(run func(*SyncTaskV2) error) *MockMetaWriter_UpdateSyncV2_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetaWriter creates a new instance of MockMetaWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetaWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetaWriter {
	mock := &MockMetaWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}