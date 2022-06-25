package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-graphql-sample/entity"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
)

var _ gateway.TodoGateway = (*TodoAccess)(nil)

type todoRow struct {
	ID   string
	Text string
	Done bool
}

type TodoAccess struct {
	db *sqlx.DB
}

func (a *TodoAccess) List(ctx context.Context) (entity.Todos, error) {
	query := `
SELECT
  id,
  text,
  done
FROM 
  todos
`
	var rows []*todoRow
	err := a.db.SelectContext(ctx, &rows, query)
	if err != nil {
		return nil, err
	}
	todos, err := a.toEntities(rows)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (a *TodoAccess) Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error) {
	return nil, nil
}
func (a *TodoAccess) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	return nil, nil
}
func (a *TodoAccess) Update(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	return nil, nil
}
func (a *TodoAccess) Delete(ctx context.Context, id entity.TodoID) error {
	return nil
}

func (a *TodoAccess) toEntity(row *todoRow) (*entity.Todo, error) {
	id, err := entity.ParseTodoID(row.ID)
	if err != nil {
		return nil, err
	}
	todo, err := entity.NewTodo(id, row.Text, row.Done)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (a *TodoAccess) toEntities(rows []*todoRow) (entity.Todos, error) {
	if rows == nil {
		return nil, nil
	}
	todos := make(entity.Todos, len(rows))
	for i, row := range rows {
		todo, err := a.toEntity(row)
		if err != nil {
			return nil, err
		}
		todos[i] = todo
	}
	return todos, nil
}
