package crud

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) Connect(mgoConf conf.MgoConfig) (*mongo.Client, error) {
	args := m.Called(mgoConf)
	return args.Get(0).(*mongo.Client), args.Error(1)
}

func (m *MockMongoClient) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockMongoClient) Collection(nameDB, nameCollection string) *mongo.Collection {
	args := m.Called(nameDB, nameCollection)
	return args.Get(0).(*mongo.Collection)
}

func (m *MockMongoClient) InsertOne(document models.Embassy) (string, error) {
	args := m.Called(document)
	return args.String(0), args.Error(1)
}
