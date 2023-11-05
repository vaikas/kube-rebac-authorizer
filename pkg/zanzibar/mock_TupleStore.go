// Code generated by mockery v2.36.0. DO NOT EDIT.

package zanzibar

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTupleStore is an autogenerated mock type for the TupleStore type
type MockTupleStore struct {
	mock.Mock
}

type MockTupleStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTupleStore) EXPECT() *MockTupleStore_Expecter {
	return &MockTupleStore_Expecter{mock: &_m.Mock}
}

// ReadTuples provides a mock function with given fields: ctx, filter
func (_m *MockTupleStore) ReadTuples(ctx context.Context, filter TupleFilter) ([]Tuple, error) {
	ret := _m.Called(ctx, filter)

	var r0 []Tuple
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, TupleFilter) ([]Tuple, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, TupleFilter) []Tuple); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Tuple)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, TupleFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTupleStore_ReadTuples_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadTuples'
type MockTupleStore_ReadTuples_Call struct {
	*mock.Call
}

// ReadTuples is a helper method to define mock.On call
//   - ctx context.Context
//   - filter TupleFilter
func (_e *MockTupleStore_Expecter) ReadTuples(ctx interface{}, filter interface{}) *MockTupleStore_ReadTuples_Call {
	return &MockTupleStore_ReadTuples_Call{Call: _e.mock.On("ReadTuples", ctx, filter)}
}

func (_c *MockTupleStore_ReadTuples_Call) Run(run func(ctx context.Context, filter TupleFilter)) *MockTupleStore_ReadTuples_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(TupleFilter))
	})
	return _c
}

func (_c *MockTupleStore_ReadTuples_Call) Return(_a0 []Tuple, _a1 error) *MockTupleStore_ReadTuples_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTupleStore_ReadTuples_Call) RunAndReturn(run func(context.Context, TupleFilter) ([]Tuple, error)) *MockTupleStore_ReadTuples_Call {
	_c.Call.Return(run)
	return _c
}

// WriteTuples provides a mock function with given fields: ctx, writes, deletes
func (_m *MockTupleStore) WriteTuples(ctx context.Context, writes []Tuple, deletes []Tuple) error {
	ret := _m.Called(ctx, writes, deletes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []Tuple, []Tuple) error); ok {
		r0 = rf(ctx, writes, deletes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTupleStore_WriteTuples_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteTuples'
type MockTupleStore_WriteTuples_Call struct {
	*mock.Call
}

// WriteTuples is a helper method to define mock.On call
//   - ctx context.Context
//   - writes []Tuple
//   - deletes []Tuple
func (_e *MockTupleStore_Expecter) WriteTuples(ctx interface{}, writes interface{}, deletes interface{}) *MockTupleStore_WriteTuples_Call {
	return &MockTupleStore_WriteTuples_Call{Call: _e.mock.On("WriteTuples", ctx, writes, deletes)}
}

func (_c *MockTupleStore_WriteTuples_Call) Run(run func(ctx context.Context, writes []Tuple, deletes []Tuple)) *MockTupleStore_WriteTuples_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]Tuple), args[2].([]Tuple))
	})
	return _c
}

func (_c *MockTupleStore_WriteTuples_Call) Return(_a0 error) *MockTupleStore_WriteTuples_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTupleStore_WriteTuples_Call) RunAndReturn(run func(context.Context, []Tuple, []Tuple) error) *MockTupleStore_WriteTuples_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTupleStore creates a new instance of MockTupleStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTupleStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTupleStore {
	mock := &MockTupleStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}