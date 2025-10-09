package service

import (
	"strconv"
	"testing"

	"concurrency-go-api/internal/model"
)

func TestCountUserPosts(t *testing.T) {
	posts := []model.Post{
		{UserId: 1, Id: 1, Title: "Post 1"},
		{UserId: 1, Id: 2, Title: "Post 2"},
		{UserId: 2, Id: 3, Title: "Post 3"},
	}

	tests := []struct {
		userId        int
		expectedCount int
	}{
		{userId: 1, expectedCount: 2},
		{userId: 2, expectedCount: 1},
		{userId: 3, expectedCount: 0},
	}

	for _, tt := range tests {
		t.Run(
			"userId="+strconv.Itoa(tt.userId),
			func(t *testing.T) {
				count := CountUserPosts(tt.userId, posts)
				if count != tt.expectedCount {
					t.Errorf("esperava %d, recebeu %d", tt.expectedCount, count)
				}
			},
		)
	}
}
