package handler

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

func Home(*resty.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`
			<html>
				<head><title>Go API</title></head>
				<body>
					<h1>Bem-vindo à API GraphQL Go!</h1>
					<p>Acesse <a href="/playground">GraphQL Playground</a></p>
				</body>
			</html>
		`))
	}
}
