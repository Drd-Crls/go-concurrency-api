package router_test

import (
	"concurrency-go-api/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestRouterRoutes(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
	})
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
	})
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := resty.New()
	client.SetBaseURL(server.URL)

	r := router.NewRouter(client)

	tests := []struct {
		route        string
		expectedCode int
	}{
		{"/", http.StatusOK},
		{"/user", http.StatusOK},
		{"/post", http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		resp := w.Result()

		if resp.StatusCode != tt.expectedCode {
			t.Errorf("para a rota %s esperava %d, recebeu %d", tt.route, tt.expectedCode, resp.StatusCode)
		}
	}
}
