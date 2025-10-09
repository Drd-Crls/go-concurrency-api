package api

import (
	"github.com/go-resty/resty/v2"

	"concurrency-go-api/internal/model"
)

const (
	USERS_ENDPOINT = "/users"
	POSTS_ENDPOINT = "/posts"
)

func FetchUsers(client *resty.Client) ([]model.User, error) {
	var users []model.User

	_, err := client.R().
		SetResult(&users).
		Get(USERS_ENDPOINT)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func FetchPosts(client *resty.Client) ([]model.Post, error) {
	var posts []model.Post
	_, err := client.R().
		SetResult(&posts).
		Get(POSTS_ENDPOINT)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
