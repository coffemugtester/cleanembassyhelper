package crud

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ MongoClient = (*MongoNonInter)(nil)

type MongoNonInter struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (t MongoNonInter) Connect(mgoConf conf.MgoConfig) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mgoConf.MongoUri))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	return client, nil
}

func (t MongoNonInter) Ping() error {
	err := t.client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Error pinging MongoDB: %v", err)
		return err
	}
	return nil
}

func (t MongoNonInter) Collection(nameDB, nameCollection string) *mongo.Collection {
	return t.client.Database(nameDB).Collection(nameCollection)
}

func (t MongoNonInter) InsertOne(document models.Embassy) (string, error) {
	id, err := t.collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Printf("Error inserting document: %v", err)
		return "", err
	}

	return fmt.Sprintf("%v", id), nil
}

type Client struct {
	mgoDB         string
	mgoCollection string
	//TODO: needs to be an interface that implements the different mongo operations
	mgoClientInterface MongoClient
	//client             *mongo.Client
}

func NewCRUDClient(mgoConf conf.MgoConfig, m MongoNonInter) *Client {

	fmt.Printf("Creating new CRUD client with config: %v\n", mgoConf)

	m.client, _ = m.Connect(mgoConf)
	m.collection = m.Collection(mgoConf.MongoDb, mgoConf.MongoCollection)

	m.Connect(mgoConf)
	fmt.Printf("Connected to MongoDB\n")
	fmt.Printf("t.client: %v\n", m.client)

	fmt.Printf("Pinging MongoDB\n")

	m.Ping()
	return &Client{
		mgoDB:              mgoConf.MongoDb,
		mgoCollection:      mgoConf.MongoCollection,
		mgoClientInterface: m,
	}
}
