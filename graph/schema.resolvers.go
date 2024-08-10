package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
    "github.com/mtariq99/graphqlexample/database"
	"github.com/mtariq99/graphqlexample/graph/model"
)
db:=database.Connect()
// CreateJobListing is the resolver for the CreateJobListing field.
func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	return db.CreateJobListingDB(input),nil
}

// UpdateJobListing is the resolver for the UpdateJobListing field.
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return db.UpdateJobListingDB(id,input),nil
}

// DeleteJobListing is the resolver for the DeleteJobListing field.
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobListingResponse, error) {
	return db.DeleteJobListingDB(id),nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	return db.GetJobs(),nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	return db.GetJobDB(id),nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
