package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) DeleteJobListingAPI(ctx context.Context, id string) (*model.DeleteJobListingResponse, error) {
	res, err := api.mongo.DeleteJobListingDB(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
