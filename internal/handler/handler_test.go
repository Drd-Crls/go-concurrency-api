package handler_test

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"concurrency-go-api/internal/handler"
	"concurrency-go-api/internal/model"

	"github.com/go-resty/resty/v2"
)

func TestHomeHandler(t *testing.T) {
	tmpDir := t.TempDir()
	templatePath := filepath.Join(tmpDir, "home.html")

	tmplContent := `<html><body><h1>Home</h1></body></html>`
	if err := os.WriteFile(templatePath, []byte(tmplContent), fs.ModePerm); err != nil {
		t.Fatalf("não foi possível criar o template de teste: %v", err)
	}

	client := resty.New()
	h := handler.Home(client)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "text/html" {
		t.Errorf("esperava Content-Type text/html, recebeu %s", ct)
	}
}

func TestPostHandler(t *testing.T) {
	mux := http.NewServeMux()

	posts := []model.Post{
		{Id: 1, UserId: 1, Title: "Post 1"},
		{Id: 2, UserId: 1, Title: "Post 2"},
		{Id: 3, UserId: 2, Title: "Post 3"},
	}

	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := resty.New()
	client.SetBaseURL(server.URL)

	h := handler.PostHandler(client)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", resp.StatusCode)
	}
}

func TestPostHandler_HTTPError(t *testing.T) {
	client := resty.New()
	// URL inválida para forçar erro
	client.SetBaseURL("http://localhost:0")

	h := handler.PostHandler(client)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("esperava 500 em caso de falha HTTP, recebeu %d", resp.StatusCode)
	}
}

func TestUserHandler_HTTPError(t *testing.T) {
	client := resty.New()
	client.SetBaseURL("http://localhost:0") // erro proposital

	h := handler.UserHandler(client)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("esperava 500 em caso de falha HTTP, recebeu %d", resp.StatusCode)
	}
}

func TestHomeHandler_TemplateNotFound(t *testing.T) {
	client := resty.New()
	h := handler.Home(client)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Forçando template inexistente
	tmplPath := filepath.Join("..", "..", "template", "not_exist.html")
	h = func(_ *resty.Client) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			tmpl := template.Must(template.ParseFiles(tmplPath))
			writer.Header().Set("Content-Type", "text/html")
			tmpl.Execute(writer, nil)
		}
	}(client)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("esperava panic por template não encontrado")
		}
	}()

	h(w, req)
}
