package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

func Home(*resty.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tmplPath := filepath.Join("templates", "index.html")
		tmpl := template.Must(template.ParseFiles(tmplPath))
		writer.Header().Set("Content-Type", "text/html")
		tmpl.Execute(writer, nil)
	}
}
