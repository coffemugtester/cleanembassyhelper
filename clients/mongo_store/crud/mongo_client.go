package crud

import (
	"clean_embassy_helper/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	*mongo.Client
}

func NewMongoClient(uri string) (*MongoClient, error) { // (uri string) (*MongoClient, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	client.Ping(context.Background(), nil)

	return &MongoClient{Client: client}, nil
}

func InsertDocument(client *MongoClient, embassy models.Embassy) (interface{}, error) {

	collection := client.Database("embassy").Collection("embassy")

	insertedDoc, err := collection.InsertOne(context.TODO(), embassy)
	if err != nil {
		return nil, err
	}

	return insertedDoc, nil
}
