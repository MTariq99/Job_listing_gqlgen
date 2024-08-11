package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type JobListingDB interface {
	CreateJobListingDB(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error)
	DeleteJobListingDB(id string) (*model.DeleteJobListingResponse, error)
	JobsDB() ([]*model.JobListing, error)
	JobDB(id string) (*model.JobListing, error)
	UpdateJobListingDB(id string, input model.UpdateJobListingInput) (*model.JobListing, error)
}

type JobListingDBImpl struct {
	Client *mongo.Client
	// Collection map[string]*mongo.Collection
}

func ConnectMongo() *JobListingDBImpl {
	clientOptions := options.Client().ApplyURI(config.Cfg.MongoURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("connection failed to database: ", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongo database")
	return &JobListingDBImpl{
		Client: client,
	}
}

func LoadCollection(client *mongo.Client) map[string]*mongo.Collection {
	collection := make(map[string]*mongo.Collection, 6)
	collection["jobs"] = collectionHelper(client, "jobs")
	return collection
}

func collectionHelper(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("job_listing").Collection(collectionName)
}

var _ JobListingDB = &JobListingDBImpl{}
