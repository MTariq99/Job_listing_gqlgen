package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/database/mongo"
	"github.com/mtariq99/graphqlexample/graph/model"
)

type JobListingAPI interface {
	CreateJobListingAPI(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error)
	UpdateJobListingAPI(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error)
	DeleteJobListingAPI(ctx context.Context, id string) (*model.DeleteJobListingResponse, error)
	JobsAPI(ctx context.Context) ([]*model.JobListing, error)
	JobAPI(ctx context.Context, id string) (*model.JobListing, error)
}

type JobListingAPIImpl struct {
	mongo mongo.JobListingDB
}

func NewJobListingAPIImpl() *JobListingAPIImpl {
	mongo := mongo.ConnectMongo()
	return &JobListingAPIImpl{
		mongo: mongo,
	}
}

var _ JobListingAPI = &JobListingAPIImpl{}
