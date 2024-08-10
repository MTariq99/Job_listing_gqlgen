package graph

import "github.com/mtariq99/graphqlexample/api"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	api api.JobListingAPI
}

func NewResolver() *Resolver {
	return &Resolver{
		api: api.NewJobListingAPIImpl(),
	}
}
