package mongo

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
)

func (mongo *JobListingDBImpl) DeleteJobListingDB(ctx context.Context, collectionName string, id string) (*model.DeleteJobListingResponse, error) {
	return nil, nil
}
