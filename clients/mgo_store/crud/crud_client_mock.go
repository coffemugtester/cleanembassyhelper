package crud

import (
	"clean_embassy_helper/internal/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockClient struct {
	mock.Mock
}

type MockClientExpectations func(*MockClient)

func NewMockCRUDClient(expectations MockClientExpectations) *MockClient {
	m := &MockClient{}
	expectations(m)
	return m
}

func (m MockClient) InsertDocument(_ *Client, document models.Embassy) (*mongo.InsertOneResult, error) {
	args := m.Called(document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
