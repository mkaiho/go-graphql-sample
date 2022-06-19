package interactor

import (
	"context"
	"errors"
	"fmt"

	"github.com/mkaiho/go-graphql-sample/entity"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
)

var _ TodoInteractor = (*todoInteractorImpl)(nil)

/**
Add new todo
**/
type (
	AddTodoInput struct {
		Text string
	}
	AddTodoOutput struct {
		ID   string
		Text string
		Done bool
	}
	TodoInteractor interface {
		AddTodo(ctx context.Context, input *AddTodoInput) (*AddTodoOutput, error)
	}
)

func (i *AddTodoInput) Validate() error {
	if i == nil {
		return errors.New("input is empty")
	}
	if len(i.Text) == 0 {
		return errors.New("input.Text is empty")
	}
	return nil
}

type todoInteractorImpl struct {
	idm         gateway.TodoIDManager
	todoGateway gateway.TodoGateway
}

func NewCreateTodoInteractor(
	idm gateway.TodoIDManager,
	toDoGateway gateway.TodoGateway,
) *todoInteractorImpl {
	return &todoInteractorImpl{
		idm:         idm,
		todoGateway: toDoGateway,
	}
}

func (u *todoInteractorImpl) AddTodo(ctx context.Context, input *AddTodoInput) (*AddTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	id := u.idm.Generate()
	todo, err := entity.NewTodo(id, input.Text, false)
	if err != nil {
		return nil, fmt.Errorf("faled to add new todo: %w", err)
	}
	todo, err = u.todoGateway.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("faled to add new todo: %w", err)
	}
	return &AddTodoOutput{
		ID:   todo.ID().String(),
		Text: todo.Text(),
		Done: todo.Done(),
	}, nil
}
