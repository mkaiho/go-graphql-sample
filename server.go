package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-graphql-sample/adapter/mysql"
	"github.com/mkaiho/go-graphql-sample/adapter/uuid"
	"github.com/mkaiho/go-graphql-sample/graph"
	"github.com/mkaiho/go-graphql-sample/graph/generated"
	"github.com/mkaiho/go-graphql-sample/infrastructure"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
	"github.com/mkaiho/go-graphql-sample/usecase/interactor"
)

const defaultPort = "3080"

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// infrastructure
	var (
		db *sqlx.DB
	)
	{
		db, err = mysql.NewDB(&infrastructure.MySQLConfig{})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// gateway
	var (
		idm         gateway.IDManager
		todoGateway gateway.TodoGateway
	)
	{
		idm = new(uuid.UUIDManager)
		todoGateway = mysql.NewTodoAccess(db)
	}
	// interactor
	var (
		todoInteractor interactor.TodoInteractor
	)
	{
		todoInteractor = interactor.NewCreateTodoInteractor(idm, todoGateway)
	}
	resolver := &graph.Resolver{
		TodoInteractor: todoInteractor,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
