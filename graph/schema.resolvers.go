package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/seb-schulz/hello-world/graph/model"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context, name string) (*model.Hello, error) {
	return &model.Hello{Name: name}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }