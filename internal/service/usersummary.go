package service

import (
	"meu-app/internal/api"
	"meu-app/internal/model"

	"github.com/go-resty/resty/v2"
)

func FetchToUserSummary(client *resty.Client, userId int) ([]model.UserSummary, error) {
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
		if userId == 0 {
			postCount := CountUserPosts(user.Id, postsRes.Data)
			usersSummary = append(usersSummary, model.UserSummary{
				Name:      user.Name,
				Email:     user.Email,
				PostCount: postCount,
			})
			return usersSummary, nil
		}
		if userId == user.Id {
			postCount := CountUserPosts(userId, postsRes.Data)
			usersSummary = append(usersSummary, model.UserSummary{
				Name:      user.Name,
				Email:     user.Email,
				PostCount: postCount,
			})
			return usersSummary, nil
		}
	}
	return nil, nil
}
