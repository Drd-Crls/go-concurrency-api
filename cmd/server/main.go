package main

import (
	"concurrency-go-api/internal/router"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const DEFAULT_PORT string = ":8080"
const API_URL string = "https://jsonplaceholder.typicode.com"

func main() {
	client := resty.New()
	client.SetBaseURL(API_URL)

	router := router.NewRouter(client)

	log.Fatal(http.ListenAndServe(DEFAULT_PORT, router))
}
