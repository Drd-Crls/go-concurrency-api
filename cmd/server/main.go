package main

import (
	"log"
	"meu-app/internal/router"
	"net/http"
)

const DEFAULT_PORT = ":8080"

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(DEFAULT_PORT, router))
}
