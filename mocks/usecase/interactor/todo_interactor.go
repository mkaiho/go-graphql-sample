// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interactor "github.com/mkaiho/go-graphql-sample/usecase/interactor"
	mock "github.com/stretchr/testify/mock"
)

// TodoInteractor is an autogenerated mock type for the TodoInteractor type
type TodoInteractor struct {
	mock.Mock
}

// AddTodo provides a mock function with given fields: ctx, input
func (_m *TodoInteractor) AddTodo(ctx context.Context, input *interactor.AddTodoInput) (*interactor.AddTodoOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *interactor.AddTodoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *interactor.AddTodoInput) *interactor.AddTodoOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interactor.AddTodoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *interactor.AddTodoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodo provides a mock function with given fields: ctx, input
func (_m *TodoInteractor) DeleteTodo(ctx context.Context, input *interactor.DeleteTodoInput) (*interactor.DeleteTodoOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *interactor.DeleteTodoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *interactor.DeleteTodoInput) *interactor.DeleteTodoOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interactor.DeleteTodoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *interactor.DeleteTodoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTodo provides a mock function with given fields: ctx, input
func (_m *TodoInteractor) FindTodo(ctx context.Context, input *interactor.FindTodoInput) (*interactor.FindTodoOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *interactor.FindTodoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *interactor.FindTodoInput) *interactor.FindTodoOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interactor.FindTodoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *interactor.FindTodoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTodos provides a mock function with given fields: ctx, input
func (_m *TodoInteractor) ListTodos(ctx context.Context, input *interactor.ListTodoInput) (*interactor.ListTodoOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *interactor.ListTodoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *interactor.ListTodoInput) *interactor.ListTodoOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interactor.ListTodoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *interactor.ListTodoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: ctx, input
func (_m *TodoInteractor) UpdateTodo(ctx context.Context, input *interactor.UpdateTodoInput) (*interactor.UpdateTodoOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *interactor.UpdateTodoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *interactor.UpdateTodoInput) *interactor.UpdateTodoOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interactor.UpdateTodoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *interactor.UpdateTodoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoInteractor interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoInteractor creates a new instance of TodoInteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoInteractor(t mockConstructorTestingTNewTodoInteractor) *TodoInteractor {
	mock := &TodoInteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
