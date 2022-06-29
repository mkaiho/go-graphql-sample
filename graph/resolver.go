package graph

import (
	"github.com/mkaiho/go-graphql-sample/usecase/interactor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoInteractor interactor.TodoInteractor
}
