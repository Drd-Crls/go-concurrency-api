package graph

import "github.com/go-resty/resty/v2"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *resty.Client
}
