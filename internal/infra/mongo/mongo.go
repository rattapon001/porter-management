package mongo_db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDbConnection(config *MongoConfig) (*mongo.Client, error) {

	uri := "mongodb://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port
	log.Println("Mongo uri :: ", uri)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}
	return client, nil
}
