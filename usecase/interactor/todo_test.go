package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/mkaiho/go-graphql-sample/entity"
	mockgateway "github.com/mkaiho/go-graphql-sample/mocks/usecase/gateway"
	"github.com/stretchr/testify/assert"
)

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
	type mockTodoGatewayCreate struct {
		todo *entity.Todo
		err  error
	}
	defalutMockTodoGatewayCreate := mockTodoGatewayCreate{
		todo: func() *entity.Todo {
			todo, _ := entity.NewTodo(todoID, "my todo", false)
			return todo
		}(),
	}
	type mocks struct {
		todoGatewayCreate mockTodoGatewayCreate
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
				todoGatewayCreate: defalutMockTodoGatewayCreate,
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
				todoGatewayCreate: defalutMockTodoGatewayCreate,
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
				todoGatewayCreate: defalutMockTodoGatewayCreate,
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
				todoGatewayCreate: mockTodoGatewayCreate{
					todo: nil,
					err:  errors.New("failed to create todo"),
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
					Return(tt.mocks.todoGatewayCreate.todo, tt.mocks.todoGatewayCreate.err)
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
	type mockTodoGatewayUpdate struct {
		todo *entity.Todo
		err  error
	}
	defalutMockTodoGatewayUpdate := mockTodoGatewayUpdate{
		todo: func() *entity.Todo {
			todo, _ := entity.NewTodo(todoID, "my todo", true)
			return todo
		}(),
	}
	type mocks struct {
		todoGatewayUpdate mockTodoGatewayUpdate
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
				todoGatewayUpdate: defalutMockTodoGatewayUpdate,
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
				todoGatewayUpdate: defalutMockTodoGatewayUpdate,
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
				todoGatewayUpdate: defalutMockTodoGatewayUpdate,
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
				todoGatewayUpdate: defalutMockTodoGatewayUpdate,
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
				todoGatewayUpdate: mockTodoGatewayUpdate{
					todo: nil,
					err:  errors.New("failed to update todo"),
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
					Return(tt.mocks.todoGatewayUpdate.todo, tt.mocks.todoGatewayUpdate.err)
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
