package main

import (
	"fmt"

	"meu-app/internal/api"

	"github.com/go-resty/resty/v2"
)

func main() {
	fmt.Println("watzap")

	api.FetchUsers(resty.New())
}
