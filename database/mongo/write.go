package mongo

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
)

func (mongo *JobListingDBImpl) CreateJobListingDB(ctx context.Context, collectionName string, input model.CreateJobListingInput) (*model.JobListing, error) {
	return nil, nil
}
