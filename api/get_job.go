package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) JobAPI(ctx context.Context, id string) (*model.JobListing, error) {
	res, err := api.mongo.JobDB(ctx, config.Cfg.Collection, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
