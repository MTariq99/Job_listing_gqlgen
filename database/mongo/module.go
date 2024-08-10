package mongo

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobListingDB interface {
	CreateJobListingDB(ctx context.Context, input model.JobListing) (*model.JobListing, error)
	DeleteJobListingDB(ctx context.Context, id string) (*model.DeleteJobListingResponse, error)
	JobsDB() ([]*model.JobListing, error)
	JobDB(id string) (*model.JobListing, error)
	UpdateJobListingDB(id string, input model.UpdateJobListingInput) (*model.JobListing, error)
}

type JobListingDBImpl struct {
	Client     *mongo.Client
	Collection map[string]*mongo.Collection
}

func ConnectMongo() *JobListingDBImpl {
	return &JobListingDBImpl{}
}

var _ JobListingDB = &JobListingDBImpl{}
