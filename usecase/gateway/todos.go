package gateway

import (
	"context"

	"github.com/mkaiho/go-graphql-sample/entity"
)

type TodoIDManager interface {
	Generate() entity.TodoID
}

type TodoGateway interface {
	List(ctx context.Context) (entity.Todos, error)
	Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error)
	Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	ChangeDoneStatus(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	Delete(ctx context.Context, id entity.TodoID) error
}
