package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	"geektest/internal/logs"
)

func getCollection(collectionName string) (collection *mongo.Collection) {
	var err error
	client, err = mongo.Connect(ctx, fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort))
	if err != nil {
		logs.Error("mongo new client failure ", err)
	}

	collection = client.Database(dbName).Collection(collectionName)
	return
}

func findOne(collectionName string, filter bson.D) (*mongo.SingleResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := getCollection(collectionName)
	result := collection.FindOne(ctx, filter)
	return result, nil
}
