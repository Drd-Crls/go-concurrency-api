package service_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"concurrency-go-api/internal/model"
	"concurrency-go-api/internal/service"

	"github.com/go-resty/resty/v2"
)

func TestFetchToUserSummary_Success(t *testing.T) {
	users := []model.User{
		{Id: 1, Name: "Alice", Email: "alice@test.com"},
		{Id: 2, Name: "Bob", Email: "bob@test.com"},
	}
	posts := []model.Post{
		{Id: 1, UserId: 1, Title: "Post 1"},
		{Id: 2, UserId: 1, Title: "Post 2"},
		{Id: 3, UserId: 2, Title: "Post 3"},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := resty.New()
	client.SetBaseURL(server.URL)

	summary, err := service.FetchToUserSummary(client, 1)

	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}

	if len(summary) != 1 {
		t.Errorf("esperava 2 summaries, recebeu %d", len(summary))
	}

	if summary[0].PostCount != 2 {
		t.Errorf("esperava 2 posts para Alice, recebeu %d", summary[0].PostCount)
	}
}
