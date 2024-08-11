package api

import (
	"context"

	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph/model"
)

func (api *JobListingAPIImpl) DeleteJobListingAPI(ctx context.Context, id string) (*model.DeleteJobListingResponse, error) {
	res, err := api.mongo.DeleteJobListingDB(ctx, config.Cfg.Collection, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
