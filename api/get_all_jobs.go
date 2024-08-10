package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) JobsAPI(ctx context.Context) ([]*model.JobListing, error) {
	res, err := api.mongo.JobsDB()
	if err != nil {
		return nil, err
	}
	return res, nil
}
