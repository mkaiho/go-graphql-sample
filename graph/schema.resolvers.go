package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mkaiho/go-graphql-sample/graph/generated"
	"github.com/mkaiho/go-graphql-sample/graph/model"
	"github.com/mkaiho/go-graphql-sample/usecase/interactor"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	out, err := r.TodoInteractor.AddTodo(ctx, &interactor.AddTodoInput{
		Text: input.Text,
	})
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   out.ID,
		Text: out.Text,
		Done: out.Done,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("user_%s", input.UserID),
		},
	}, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	out, err := r.TodoInteractor.UpdateTodo(ctx, &interactor.UpdateTodoInput{
		ID:   input.ID,
		Text: input.Text,
		Done: input.Done,
	})
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   out.ID,
		Text: out.Text,
		Done: out.Done,
		User: &model.User{
			ID:   "user_001",
			Name: "dummy_user",
		},
	}, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) (bool, error) {
	out, err := r.TodoInteractor.DeleteTodo(ctx, &interactor.DeleteTodoInput{
		ID: input,
	})
	if err != nil {
		return false, err
	}
	return out.Deleted, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	out, err := r.TodoInteractor.ListTodos(ctx, &interactor.ListTodoInput{})
	if err != nil {
		return nil, err
	}
	todos := make([]*model.Todo, len(out.Todos))
	for i, todo := range out.Todos {
		todos[i] = &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
			User: &model.User{
				ID:   fmt.Sprintf("user_id_%03d", i+1),
				Name: fmt.Sprintf("user_%03d", i+1),
			},
		}
	}
	return todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, input string) (*model.Todo, error) {
	out, err := r.TodoInteractor.FindTodo(ctx, &interactor.FindTodoInput{
		ID: input,
	})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, nil
	}
	return &model.Todo{
		ID:   out.ID,
		Text: out.Text,
		Done: out.Done,
		User: &model.User{
			ID:   "user_001",
			Name: "dummy_user",
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
