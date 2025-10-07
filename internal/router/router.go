package router

import (
	"meu-app/internal/handler"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func NewRouter() *http.ServeMux {
	server := http.NewServeMux()
	client := resty.New()

	server.HandleFunc("/user", handler.UserHandler(client))
	server.HandleFunc("/userPost", handler.UserPostHandler(client))

	return server
}
