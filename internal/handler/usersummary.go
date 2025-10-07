package handler

import (
	"encoding/json"
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

		var usersSummary []model.UserSummary

		for _, user := range usersRes.Data {
			postCount := service.CountUserPosts(user.Id, postsRes.Data)
			usersSummary = append(usersSummary, model.UserSummary{
				Name:      user.Name,
				Email:     user.Email,
				PostCount: postCount,
			})
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(usersSummary)
	}
}
