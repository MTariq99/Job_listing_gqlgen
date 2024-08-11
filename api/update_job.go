package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) UpdateJobListingAPI(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	res, err := api.mongo.UpdateJobListingDB(ctx, config.Cfg.Collection, id, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
