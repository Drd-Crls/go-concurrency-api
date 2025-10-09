package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"concurrency-go-api/internal/api"
	"concurrency-go-api/internal/model"

	"github.com/go-resty/resty/v2"
)

func TestFetchUsers_Success(t *testing.T) {
	mockUsers := []model.User{
		{Id: 1, Name: "Alice", Email: "alice@test.com"},
		{Id: 2, Name: "Bob", Email: "bob@test.com"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users" {
			t.Errorf("esperava /users, recebeu %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(mockUsers)
	}))
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)

	users, err := api.FetchUsers(client)
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("esperava 2 users, recebeu %d", len(users))
	}
	if users[0].Name != "Alice" {
		t.Errorf("esperava Alice, recebeu %s", users[0].Name)
	}
}

func TestFetchPosts_Success(t *testing.T) {
	mockPosts := []model.Post{
		{Id: 1, UserId: 1, Title: "Post 1"},
		{Id: 2, UserId: 1, Title: "Post 2"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/posts" {
			t.Errorf("esperava /posts, recebeu %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(mockPosts)
	}))
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)

	posts, err := api.FetchPosts(client)
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}

	if len(posts) != 2 {
		t.Errorf("esperava 2 posts, recebeu %d", len(posts))
	}
	if posts[0].Title != "Post 1" {
		t.Errorf("esperava Post 1, recebeu %s", posts[0].Title)
	}
}

func TestFetchUsers_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "erro fake", http.StatusInternalServerError)
	}))
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)

	resp, err := client.R().Get("/users")
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if resp.StatusCode() == http.StatusOK {
		t.Fatal("esperava erro, recebeu status 200")
	}
}

func TestFetchPosts_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "erro fake", http.StatusInternalServerError)
	}))
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)

	resp, err := client.R().Get("/posts")
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if resp.StatusCode() == http.StatusOK {
		t.Fatal("esperava erro, recebeu status 200")
	}
}
