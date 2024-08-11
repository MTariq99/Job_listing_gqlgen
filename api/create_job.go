package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) CreateJobListingAPI(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	res, err := api.mongo.CreateJobListingDB(ctx, config.Cfg.Collection, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
