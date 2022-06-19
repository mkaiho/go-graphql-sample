package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/mkaiho/go-graphql-sample/entity"
	mockgateway "github.com/mkaiho/go-graphql-sample/mocks/usecase/gateway"
	"github.com/stretchr/testify/assert"
)

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
