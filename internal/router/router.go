package router

import (
	"concurrency-go-api/graph"
	"concurrency-go-api/internal/handler"
	"net/http"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-resty/resty/v2"
)

func NewRouter(client *resty.Client) *http.ServeMux {
	server := http.NewServeMux()

	gqlServer := gqlHandler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			Client: client,
		}}),
	)

	server.HandleFunc("/", handler.Home(client))

	// rest api apenas para fins de visualizar os dados retornados diretamente pela aplicação
	server.HandleFunc("/user", handler.UserHandler(client))
	server.HandleFunc("/post", handler.PostHandler(client))

	//endpoint GraphQL
	server.Handle("/query", gqlServer)
	server.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	return server
}
