package gateway

import (
	"context"

	"github.com/mkaiho/go-graphql-sample/entity"
)

type TodoGateway interface {
	List(ctx context.Context) (entity.Todos, error)
	Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error)
	Create(ctx context.Context, todo *entity.Todo) error
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id entity.TodoID) (bool, error)
}
