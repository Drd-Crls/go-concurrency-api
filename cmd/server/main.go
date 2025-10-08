package main

import (
	"log"
	"meu-app/internal/router"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const DEFAULT_PORT string = ":8080"

func main() {
	client := resty.New()

	router := router.NewRouter(client)

	log.Fatal(http.ListenAndServe(DEFAULT_PORT, router))
}
