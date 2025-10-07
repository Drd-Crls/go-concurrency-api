package router

import (
	"meu-app/graph"
	"meu-app/internal/handler"
	"net/http"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"

	"github.com/go-resty/resty/v2"
)

func NewRouter() *http.ServeMux {
	server := http.NewServeMux()
	client := resty.New()

	gqlServer := gqlHandler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}),
	)

	server.HandleFunc("/user", handler.UserHandler(client))
	server.HandleFunc("/userPost", handler.UserPostHandler(client))
	server.Handle("/query", gqlServer)

	return server
}
