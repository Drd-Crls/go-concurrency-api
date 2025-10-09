package service

import (
	"concurrency-go-api/internal/model"
)

func CountUserPosts(userId int, posts []model.Post) int {
	count := 0
	for _, post := range posts {
		if post.UserId == userId {
			count++
		}
	}
	return count
}
