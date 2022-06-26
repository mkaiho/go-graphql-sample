package interactor

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/mkaiho/go-graphql-sample/entity"
	mockgateway "github.com/mkaiho/go-graphql-sample/mocks/usecase/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_todoInteractorImpl_ListTodos(t *testing.T) {
	defaultTodos := func() entity.Todos {
		todos := entity.Todos{}
		for i := 0; i < 3; i++ {
			id, _ := entity.ParseTodoID(fmt.Sprintf("test_id_%03d", i+1))
			todo, _ := entity.NewTodo(id, fmt.Sprintf("my todo %d", i+1), (i%2) == 0)
			todos = append(todos, todo)
		}
		return todos
	}()
	defaultListTodoOutput := func() ListTodoOutput {
		var items ListTodoOutputTodoItems
		for _, todo := range defaultTodos {
			items = append(items, &ListTodoOutputTodoItem{
				ID:   todo.ID().String(),
				Text: todo.Text(),
				Done: todo.Done(),
			})
		}
		return ListTodoOutput{
			Todos: items,
		}
	}()
	type mockTodoGatewayList struct {
		todos entity.Todos
		err   error
	}
	type mocks struct {
		todoGatewayList mockTodoGatewayList
	}
	type args struct {
		ctx   context.Context
		input *ListTodoInput
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *ListTodoOutput
		wantErr bool
	}{
		{
			name: "return output",
			mocks: mocks{
				todoGatewayList: mockTodoGatewayList{
					todos: defaultTodos,
				},
			},
			args: args{
				ctx:   context.Background(),
				input: &ListTodoInput{},
			},
			want:    &defaultListTodoOutput,
			wantErr: false,
		},
		{
			name: "return error when the gateway failed to list todos",
			mocks: mocks{
				todoGatewayList: mockTodoGatewayList{
					todos: nil,
					err:   errors.New("failed to list todos"),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: &ListTodoInput{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoGateway := new(mockgateway.TodoGateway)
			if tt.args.input != nil {
				todoGateway.
					On("List", tt.args.ctx).
					Return(tt.mocks.todoGatewayList.todos, tt.mocks.todoGatewayList.err)
			}
			u := &todoInteractorImpl{
				todoGateway: todoGateway,
			}
			got, err := u.ListTodos(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoInteractorImpl.ListTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "todoInteractorImpl.ListTodo() = %v, want %v", got, tt.want)
			if !tt.wantErr {
				todoGateway.AssertNumberOfCalls(t, "List", 1)
			}
		})
	}
}

func Test_todoInteractorImpl_FindTodo(t *testing.T) {
	var todoID entity.TodoID = "todo_id"
	type mockTodoGatewayFind struct {
		todo *entity.Todo
		err  error
	}
	defalutMockTodoGatewayFind := mockTodoGatewayFind{
		todo: func() *entity.Todo {
			todo, _ := entity.NewTodo(todoID, "my todo", true)
			return todo
		}(),
	}
	type mocks struct {
		todoGatewayFind mockTodoGatewayFind
	}
	type args struct {
		ctx   context.Context
		input *FindTodoInput
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *FindTodoOutput
		wantErr bool
	}{
		{
			name: "return output",
			mocks: mocks{
				todoGatewayFind: defalutMockTodoGatewayFind,
			},
			args: args{
				ctx: context.Background(),
				input: &FindTodoInput{
					ID: todoID.String(),
				},
			},
			want: &FindTodoOutput{
				ID:   todoID.String(),
				Text: "my todo",
				Done: true,
			},
			wantErr: false,
		},
		{
			name: "return error when input is nil",
			mocks: mocks{
				todoGatewayFind: defalutMockTodoGatewayFind,
			},
			args: args{
				ctx:   context.Background(),
				input: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when input.ID is invalid",
			mocks: mocks{
				todoGatewayFind: defalutMockTodoGatewayFind,
			},
			args: args{
				ctx: context.Background(),
				input: &FindTodoInput{
					ID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when the gateway failed to find todo",
			mocks: mocks{
				todoGatewayFind: mockTodoGatewayFind{
					todo: nil,
					err:  errors.New("failed to find todo"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: &FindTodoInput{
					ID: "test_id",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoGateway := new(mockgateway.TodoGateway)
			if tt.args.input != nil {
				todoGateway.
					On("Find", tt.args.ctx, entity.TodoID(tt.args.input.ID)).
					Return(tt.mocks.todoGatewayFind.todo, tt.mocks.todoGatewayFind.err)
			}
			u := &todoInteractorImpl{
				todoGateway: todoGateway,
			}
			got, err := u.FindTodo(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoInteractorImpl.FindTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "todoInteractorImpl.FindTodo() = %v, want %v", got, tt.want)
			if !tt.wantErr {
				todoGateway.AssertNumberOfCalls(t, "Find", 1)
			}
		})
	}
}

func Test_todoInteractorImpl_AddTodo(t *testing.T) {
	var todoID entity.TodoID = "todo_id"
	type errFunc struct {
		err error
	}
	type mocks struct {
		todoGatewayCreate errFunc
	}
	type args struct {
		ctx   context.Context
		input *AddTodoInput
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *AddTodoOutput
		wantErr bool
	}{
		{
			name: "return output",
			mocks: mocks{
				todoGatewayCreate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &AddTodoInput{
					Text: "my todo",
				},
			},
			want: &AddTodoOutput{
				ID:   todoID.String(),
				Text: "my todo",
				Done: false,
			},
			wantErr: false,
		},
		{
			name: "return error when input is nil",
			mocks: mocks{
				todoGatewayCreate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx:   context.Background(),
				input: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when input.Text is empty",
			mocks: mocks{
				todoGatewayCreate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &AddTodoInput{
					Text: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when the gateway failed to create todo",
			mocks: mocks{
				todoGatewayCreate: errFunc{
					err: errors.New("failed to create todo"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: &AddTodoInput{
					Text: "my todo",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idm := new(mockgateway.TodoIDManager)
			idm.
				On("Generate").
				Return(todoID)
			todoGateway := new(mockgateway.TodoGateway)
			if tt.args.input != nil {
				createdTodo, _ := entity.NewTodo(todoID, tt.args.input.Text, false)
				todoGateway.
					On("Create", tt.args.ctx, createdTodo).
					Return(tt.mocks.todoGatewayCreate.err)
			}
			u := &todoInteractorImpl{
				idm:         idm,
				todoGateway: todoGateway,
			}
			got, err := u.AddTodo(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoInteractorImpl.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "todoInteractorImpl.AddTodo() = %v, want %v", got, tt.want)
			if !tt.wantErr {
				todoGateway.AssertNumberOfCalls(t, "Create", 1)
			}
		})
	}
}

func Test_todoInteractorImpl_UpdateTodo(t *testing.T) {
	var todoID entity.TodoID = "todo_id"
	type errFunc struct {
		err error
	}
	type mocks struct {
		todoGatewayUpdate errFunc
	}
	type args struct {
		ctx   context.Context
		input *UpdateTodoInput
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *UpdateTodoOutput
		wantErr bool
	}{
		{
			name: "return output",
			mocks: mocks{
				todoGatewayUpdate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &UpdateTodoInput{
					ID:   todoID.String(),
					Text: "my todo",
					Done: true,
				},
			},
			want: &UpdateTodoOutput{
				ID:   todoID.String(),
				Text: "my todo",
				Done: true,
			},
			wantErr: false,
		},
		{
			name: "return error when input is nil",
			mocks: mocks{
				todoGatewayUpdate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx:   context.Background(),
				input: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when input.ID is invalid",
			mocks: mocks{
				todoGatewayUpdate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &UpdateTodoInput{
					ID:   "",
					Text: "my todo",
					Done: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when input.Text is empty",
			mocks: mocks{
				todoGatewayUpdate: errFunc{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &UpdateTodoInput{
					ID:   todoID.String(),
					Text: "",
					Done: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when the gateway failed to update todo",
			mocks: mocks{
				todoGatewayUpdate: errFunc{
					err: errors.New("failed to update todo"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: &UpdateTodoInput{
					Text: "my todo",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoGateway := new(mockgateway.TodoGateway)
			if tt.args.input != nil {
				todo, _ := entity.NewTodo(todoID, tt.args.input.Text, tt.args.input.Done)
				todoGateway.
					On("Update", tt.args.ctx, todo).
					Return(tt.mocks.todoGatewayUpdate.err)
			}
			u := &todoInteractorImpl{
				todoGateway: todoGateway,
			}
			got, err := u.UpdateTodo(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoInteractorImpl.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "todoInteractorImpl.UpdateTodo() = %v, want %v", got, tt.want)
			if !tt.wantErr {
				todoGateway.AssertNumberOfCalls(t, "Update", 1)
			}
		})
	}
}

func Test_todoInteractorImpl_DeleteTodo(t *testing.T) {
	var todoID entity.TodoID = "todo_id"
	type mockTodoGatewayDelete struct {
		err error
	}
	type mocks struct {
		todoGatewayDelete mockTodoGatewayDelete
	}
	type args struct {
		ctx   context.Context
		input *DeleteTodoInput
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *DeleteTodoOutput
		wantErr bool
	}{
		{
			name: "return output",
			mocks: mocks{
				todoGatewayDelete: mockTodoGatewayDelete{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &DeleteTodoInput{
					ID: todoID.String(),
				},
			},
			want: &DeleteTodoOutput{
				ID: todoID.String(),
			},
			wantErr: false,
		},
		{
			name: "return error when input is nil",
			mocks: mocks{
				todoGatewayDelete: mockTodoGatewayDelete{
					err: nil,
				},
			},
			args: args{
				ctx:   context.Background(),
				input: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when input.ID is invalid",
			mocks: mocks{
				todoGatewayDelete: mockTodoGatewayDelete{
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: &DeleteTodoInput{
					ID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return error when the gateway failed to delete todo",
			mocks: mocks{
				todoGatewayDelete: mockTodoGatewayDelete{
					err: errors.New("failed to delete todo"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: &DeleteTodoInput{
					ID: "test_id",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoGateway := new(mockgateway.TodoGateway)
			if tt.args.input != nil {
				todoGateway.
					On("Delete", tt.args.ctx, entity.TodoID(tt.args.input.ID)).
					Return(tt.mocks.todoGatewayDelete.err)
			}
			u := &todoInteractorImpl{
				todoGateway: todoGateway,
			}
			got, err := u.DeleteTodo(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoInteractorImpl.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "todoInteractorImpl.DeleteTodo() = %v, want %v", got, tt.want)
			if !tt.wantErr {
				todoGateway.AssertNumberOfCalls(t, "Delete", 1)
			}
		})
	}
}
