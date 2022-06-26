package interactor

import (
	"context"
	"errors"
	"fmt"

	"github.com/mkaiho/go-graphql-sample/entity"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
)

var _ TodoInteractor = (*todoInteractorImpl)(nil)

type TodoInteractor interface {
	ListTodos(ctx context.Context, input *ListTodoInput) (*ListTodoOutput, error)
	FindTodo(ctx context.Context, input *FindTodoInput) (*FindTodoOutput, error)
	AddTodo(ctx context.Context, input *AddTodoInput) (*AddTodoOutput, error)
	UpdateTodo(ctx context.Context, input *UpdateTodoInput) (*UpdateTodoOutput, error)
	DeleteTodo(ctx context.Context, input *DeleteTodoInput) (*DeleteTodoOutput, error)
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

/**
List todos
**/
type (
	ListTodoInput struct {
	}
	ListTodoOutputTodoItem struct {
		ID   string
		Text string
		Done bool
	}
	ListTodoOutputTodoItems []*ListTodoOutputTodoItem
	ListTodoOutput          struct {
		Todos ListTodoOutputTodoItems
	}
)

func (i *ListTodoInput) Validate() error {
	return nil
}

func (i ListTodoOutputTodoItems) Append(items ...*ListTodoOutputTodoItem) ListTodoOutputTodoItems {
	return append(i, items...)
}

func (u *todoInteractorImpl) ListTodos(ctx context.Context, input *ListTodoInput) (*ListTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	todos, err := u.todoGateway.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("faled to list todos: %w", err)
	}
	var outputTodos ListTodoOutputTodoItems
	for _, todo := range todos {
		outputTodo := ListTodoOutputTodoItem{
			ID:   todo.ID().String(),
			Text: todo.Text(),
			Done: todo.Done(),
		}
		outputTodos = outputTodos.Append(&outputTodo)
	}
	return &ListTodoOutput{
		Todos: outputTodos,
	}, nil
}

/**
Find todo
**/
type (
	FindTodoInput struct {
		ID string
	}
	FindTodoOutput struct {
		ID   string
		Text string
		Done bool
	}
)

func (i *FindTodoInput) Validate() error {
	if i == nil {
		return errors.New("input is empty")
	}
	if _, err := entity.ParseTodoID(i.ID); err != nil {
		return fmt.Errorf("input.ID is invalid: %w", err)
	}
	return nil
}

func (u *todoInteractorImpl) FindTodo(ctx context.Context, input *FindTodoInput) (*FindTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	id, err := entity.ParseTodoID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("faled to find todo: %w", err)
	}
	todo, err := u.todoGateway.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("faled to find todo: %w", err)
	}
	return &FindTodoOutput{
		ID:   todo.ID().String(),
		Text: todo.Text(),
		Done: todo.Done(),
	}, nil
}

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

func (u *todoInteractorImpl) AddTodo(ctx context.Context, input *AddTodoInput) (*AddTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	id := u.idm.Generate()
	todo, err := entity.NewTodo(id, input.Text, false)
	if err != nil {
		return nil, fmt.Errorf("faled to add new todo: %w", err)
	}
	err = u.todoGateway.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("faled to add new todo: %w", err)
	}
	return &AddTodoOutput{
		ID:   todo.ID().String(),
		Text: todo.Text(),
		Done: todo.Done(),
	}, nil
}

/**
Update todo
**/
type (
	UpdateTodoInput struct {
		ID   string
		Text string
		Done bool
	}
	UpdateTodoOutput struct {
		ID   string
		Text string
		Done bool
	}
)

func (i *UpdateTodoInput) Validate() error {
	if i == nil {
		return errors.New("input is empty")
	}
	if _, err := entity.ParseTodoID(i.ID); err != nil {
		return fmt.Errorf("input.ID is invalid: %w", err)
	}
	if len(i.Text) == 0 {
		return errors.New("input.Text is empty")
	}
	return nil
}

func (u *todoInteractorImpl) UpdateTodo(ctx context.Context, input *UpdateTodoInput) (*UpdateTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	id, err := entity.ParseTodoID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("faled to update todo: %w", err)
	}
	todo, err := entity.NewTodo(id, input.Text, input.Done)
	if err != nil {
		return nil, fmt.Errorf("faled to update todo: %w", err)
	}
	err = u.todoGateway.Update(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("faled to update todo: %w", err)
	}
	return &UpdateTodoOutput{
		ID:   todo.ID().String(),
		Text: todo.Text(),
		Done: todo.Done(),
	}, nil
}

/**
Delete todo
**/
type (
	DeleteTodoInput struct {
		ID string
	}
	DeleteTodoOutput struct {
		ID string
	}
)

func (i *DeleteTodoInput) Validate() error {
	if i == nil {
		return errors.New("input is empty")
	}
	if _, err := entity.ParseTodoID(i.ID); err != nil {
		return fmt.Errorf("input.ID is invalid: %w", err)
	}
	return nil
}

func (u *todoInteractorImpl) DeleteTodo(ctx context.Context, input *DeleteTodoInput) (*DeleteTodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	id, err := entity.ParseTodoID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("faled to delete todo: %w", err)
	}
	err = u.todoGateway.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("faled to delete todo: %w", err)
	}
	return &DeleteTodoOutput{
		ID: id.String(),
	}, nil
}
