package mongo

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
)

func (mongo *JobListingDBImpl) UpdateJobListingDB(ctx context.Context, collectionName string, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return nil, nil
}
