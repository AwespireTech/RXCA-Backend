package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(url string) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverApi).SetConnectTimeout(10)
	database, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = database.Ping(ctx, nil)
	client = database
	return err
}

func GetClient() *mongo.Client {
	return client
}
