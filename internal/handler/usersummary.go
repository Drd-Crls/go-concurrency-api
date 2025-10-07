package handler

import (
	"encoding/json"
	"fmt"
	"meu-app/internal/api"
	"meu-app/internal/model"
	"meu-app/internal/service"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func UserPostHandler(client *resty.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		usersCh := make(chan model.Result[[]model.User])
		postsCh := make(chan model.Result[[]model.Post])

		go func() {
			users, err := api.FetchUsers(client)
			usersCh <- model.Result[[]model.User]{Data: users, Err: err}
		}()

		go func() {
			posts, err := api.FetchPosts(client)
			postsCh <- model.Result[[]model.Post]{Data: posts, Err: err}
		}()

		usersRes := <-usersCh
		postsRes := <-postsCh

		for _, user := range usersRes.Data {
			sla := service.CountUserPosts(user.Id, postsRes.Data)
			fmt.Println("Usuario ", user.Id, " postou ", sla)
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(postsRes)
	}
}
