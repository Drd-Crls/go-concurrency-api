package api

import (
	"github.com/go-resty/resty/v2"

	"meu-app/internal/model"
)

const (
	USERS_API = "https://jsonplaceholder.typicode.com/users"
	POSTS_API = "https://jsonplaceholder.typicode.com/posts"
)

func FetchUsers(client *resty.Client) ([]model.User, error) {
	var users []model.User
	_, err := client.R().
		SetResult(&users).
		Get(USERS_API)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func FetchPosts(client *resty.Client) ([]model.Post, error) {
	var posts []model.Post
	_, err := client.R().
		SetResult(&posts).
		Get(POSTS_API)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
