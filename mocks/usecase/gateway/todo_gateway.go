// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/mkaiho/go-graphql-sample/entity"

	mock "github.com/stretchr/testify/mock"
)

// TodoGateway is an autogenerated mock type for the TodoGateway type
type TodoGateway struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, todo
func (_m *TodoGateway) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	ret := _m.Called(ctx, todo)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Todo) *entity.Todo); ok {
		r0 = rf(ctx, todo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Todo) error); ok {
		r1 = rf(ctx, todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TodoGateway) Delete(ctx context.Context, id entity.TodoID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.TodoID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, id
func (_m *TodoGateway) Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, entity.TodoID) *entity.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.TodoID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *TodoGateway) List(ctx context.Context) (entity.Todos, error) {
	ret := _m.Called(ctx)

	var r0 entity.Todos
	if rf, ok := ret.Get(0).(func(context.Context) entity.Todos); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Todos)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, todo
func (_m *TodoGateway) Update(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	ret := _m.Called(ctx, todo)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Todo) *entity.Todo); ok {
		r0 = rf(ctx, todo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Todo) error); ok {
		r1 = rf(ctx, todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoGateway interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoGateway creates a new instance of TodoGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoGateway(t mockConstructorTestingTNewTodoGateway) *TodoGateway {
	mock := &TodoGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}