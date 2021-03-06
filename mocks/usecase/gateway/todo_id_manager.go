// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/mkaiho/go-graphql-sample/entity"

	mock "github.com/stretchr/testify/mock"
)

// TodoIDManager is an autogenerated mock type for the TodoIDManager type
type TodoIDManager struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *TodoIDManager) Generate() entity.TodoID {
	ret := _m.Called()

	var r0 entity.TodoID
	if rf, ok := ret.Get(0).(func() entity.TodoID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entity.TodoID)
	}

	return r0
}

type mockConstructorTestingTNewTodoIDManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoIDManager creates a new instance of TodoIDManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoIDManager(t mockConstructorTestingTNewTodoIDManager) *TodoIDManager {
	mock := &TodoIDManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
