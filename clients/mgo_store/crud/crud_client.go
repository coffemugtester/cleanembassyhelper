package crud

import (
	"clean_embassy_helper/conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	mgoDB         string
	mgoCollection string
	client        *mongo.Client
}

func NewCRUDClient(mgoConf conf.MgoConfig) *Client {
	mgo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mgoConf.MongoUri))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
		return nil
	}

	if err := mgo.Ping(context.Background(), nil); err != nil {
		fmt.Printf("Error pinging MongoDB: %v", err)
		return nil
	}

	return &Client{
		mgoDB:         mgoConf.MongoDb,
		mgoCollection: mgoConf.MongoCollection,
		client:        mgo,
	}
}
