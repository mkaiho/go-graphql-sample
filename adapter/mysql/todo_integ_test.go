//go:build mysql

package mysql

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-graphql-sample/entity"
	"github.com/stretchr/testify/assert"
)

func setupTodo() {
	err := load("todo_test_setup.sql")
	if err != nil {
		panic(err)
	}
}
func teadownTodo() {
	err := load("todo_test_teardown.sql")
	if err != nil {
		panic(err)
	}
}
func buildTodosTestData() entity.Todos {
	rows := []*todoRow{
		{
			ID:   "test_todo_id_001",
			Text: "test todo 001",
			Done: false,
		},
		{
			ID:   "test_todo_id_002",
			Text: "test todo 002",
			Done: false,
		},
		{
			ID:   "test_todo_id_003",
			Text: "test todo 003",
			Done: true,
		},
	}
	todos := make(entity.Todos, len(rows))
	for i, row := range rows {
		id, _ := entity.ParseTodoID(row.ID)
		todos[i], _ = entity.NewTodo(id, row.Text, row.Done)
	}
	return todos
}

func TestTodoAccess_List(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		setup   func()
		teadown func()
		fields  fields
		args    args
		want    entity.Todos
		wantErr bool
	}{
		{
			name:    "all TODOs are returned",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want: func() entity.Todos {
				todos := buildTodosTestData()
				return todos
			}(),
			wantErr: false,
		},
		{
			name:  "nil is returned when the todo rows does not exist",
			setup: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "an error is returned when the conversion to an entity fails",
			setup: func() {
				sql := `
INSERT INTO todos (id, text, done)
VALUES
	("test_todo_id", "", 1)
`
				testDB.Exec(sql)
			},
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "an error is returned when the context is canceled",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.teadown != nil {
				defer tt.teadown()
			}
			a := &TodoAccess{
				db: tt.fields.db,
			}
			got, err := a.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoAccess.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "TodoAccess.List() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodoAccess_Find(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id  entity.TodoID
	}
	tests := []struct {
		name    string
		setup   func()
		teadown func()
		fields  fields
		args    args
		want    *entity.Todo
		wantErr bool
	}{
		{
			name:    "the todo for the specified ID is returned",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
				id:  buildTodosTestData()[2].ID(),
			},
			want: func() *entity.Todo {
				return buildTodosTestData()[2]
			}(),
			wantErr: false,
		},
		{
			name:    "nil is returned when the todo rows does not exist",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
				id:  "no-existent-id",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "an error is returned when the conversion to an entity fails",
			setup: func() {
				sql := `
INSERT INTO todos (id, text, done)
VALUES
	("test_todo_id", "", 1)
`
				testDB.Exec(sql)
			},
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.Background(),
				id:  "test_todo_id",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "an error is returned when the context is canceled",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
				id: buildTodosTestData()[2].ID(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.teadown != nil {
				defer tt.teadown()
			}
			a := &TodoAccess{
				db: tt.fields.db,
			}
			got, err := a.Find(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoAccess.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "TodoAccess.Find() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodoAccess_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		todo *entity.Todo
	}
	tests := []struct {
		name    string
		setup   func()
		teadown func()
		fields  fields
		args    args
		want    todoRow
		wantErr bool
	}{
		{
			name:    "todo is created",
			setup:   teadownTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx:  context.Background(),
				todo: buildTodosTestData()[2],
			},
			want: func() todoRow {
				todo := buildTodosTestData()[2]
				return todoRow{
					ID:   todo.ID().String(),
					Text: todo.Text(),
					Done: todo.Done(),
				}
			}(),
			wantErr: false,
		},
		{
			name:    "an error is returned when the entity already exists",
			setup:   setupTodo,
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx:  context.Background(),
				todo: buildTodosTestData()[2],
			},
			want:    todoRow{},
			wantErr: true,
		},
		{
			name:    "an error is returned when the context is canceled",
			teadown: teadownTodo,
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
				todo: buildTodosTestData()[2],
			},
			want:    todoRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.teadown != nil {
				defer tt.teadown()
			}
			a := &TodoAccess{
				db: tt.fields.db,
			}
			err := a.Create(tt.args.ctx, tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoAccess.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var row todoRow
				err := testDB.Get(&row, "SELECT id, text, done FROM todos WHERE id = ?", buildTodosTestData()[2].ID().String())
				if err != nil {
					panic(err)
				}
				assert.Equal(t, tt.want, row)
			}
		})
	}
}
