package mongo

import (
	"context"

	"github.com/mtariq99/graphqlexample/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mongo *JobListingDBImpl) JobsDB(ctx context.Context, collectionName string) ([]*model.JobListing, error) {
	return nil, nil
}
func (mongo *JobListingDBImpl) JobDB(ctx context.Context, collectionName string, id string) (*model.JobListing, error) {
	JobList := &model.JobListing{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	if err := mongo.Collection[collectionName].FindOne(ctx, filter).Decode(&JobList); err != nil {
		return nil, err
	}

	return JobList, nil
}
