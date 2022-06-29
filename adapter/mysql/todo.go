package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-graphql-sample/entity"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
)

var _ gateway.TodoGateway = (*TodoAccess)(nil)

type todoRow struct {
	ID   string `db:"id"`
	Text string `db:"text"`
	Done bool   `db:"done"`
}

type TodoAccess struct {
	db *sqlx.DB
}

func NewTodoAccess(db *sqlx.DB) *TodoAccess {
	return &TodoAccess{
		db: db,
	}
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
	query := `
SELECT
  id,
  text,
  done
FROM 
  todos
WHERE
  id = ?
`
	var row todoRow
	err := a.db.GetContext(ctx, &row, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	todos, err := a.toEntity(&row)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (a *TodoAccess) Create(ctx context.Context, todo *entity.Todo) error {
	query := `
INSERT INTO todos (id, text, done)
VALUES (:id, :text, :done) 
`
	_, err := a.db.NamedExecContext(ctx, query, a.toRow(todo))
	if err != nil {
		return err
	}
	return nil
}

func (a *TodoAccess) Update(ctx context.Context, todo *entity.Todo) error {
	query := `
UPDATE
  todos
SET
  text = :text,
  done = :done
WHERE
  id = :id
`
	_, err := a.db.NamedExecContext(ctx, query, a.toRow(todo))
	if err != nil {
		return err
	}
	return nil
}

func (a *TodoAccess) Delete(ctx context.Context, id entity.TodoID) (bool, error) {
	query := `
DELETE FROM
  todos
WHERE
  id = ?
`
	result, err := a.db.ExecContext(ctx, query, id.String())
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected != 0, nil
}

func (a *TodoAccess) toRow(e *entity.Todo) *todoRow {
	if e == nil {
		return nil
	}
	return &todoRow{
		ID:   e.ID().String(),
		Text: e.Text(),
		Done: e.Done(),
	}
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
