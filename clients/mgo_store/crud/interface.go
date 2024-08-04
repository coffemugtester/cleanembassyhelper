package crud

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Connect(mgoConf conf.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document models.Embassy) (string, error)
}
