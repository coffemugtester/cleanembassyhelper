package crud

import (
	"clean_embassy_helper/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoClient struct {
	*mongo.Client
}

func NewMongoClient(uri string) (*MgoClient, error) { // (uri string) (*MgoClient, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	client.Ping(context.Background(), nil)

	return &MgoClient{Client: client}, nil
}

func InsertDocument(client *MgoClient, embassy models.Embassy) (interface{}, error) {

	collection := client.Database("embassy").Collection("embassy")

	insertedDoc, err := collection.InsertOne(context.TODO(), embassy)
	if err != nil {
		return nil, err
	}

	return insertedDoc, nil
}
