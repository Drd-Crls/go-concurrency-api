package handler

import (
	"encoding/json"
	"meu-app/internal/api"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func UserHandler(client *resty.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := api.FetchUsers(client)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(users)
	}
}
