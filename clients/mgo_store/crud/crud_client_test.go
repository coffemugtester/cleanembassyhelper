package crud

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

// TODO: implement test containers
func TestMockMongoClient_Connect(t *testing.T) {

	mongo := new(mongo.Client)

	mockClient := MockMongoClient{}

	mgoConf := conf.MgoConfig{
		MongoUri:        "",
		MongoCollection: "",
		MongoDb:         "",
	}

	mockClient.On("Connect", mgoConf).Return(mongo, nil)

	client, err := mockClient.Connect(mgoConf)
	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, client)

	mockClient.AssertExpectations(t)

}

func TestMockMongoClient_InsertOne(t *testing.T) {

	mockClient := MockMongoClient{}

	document := models.Embassy{}

	mockClient.On("InsertOne", document).Return("123", nil)

	id, err := mockClient.InsertOne(document)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "123", id)

	mockClient.AssertExpectations(t)
}
