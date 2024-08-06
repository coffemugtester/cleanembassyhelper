package mapz

import "github.com/stretchr/testify/mock"

type MockMapsClient struct {
	mock.Mock
}

func (m *MockMapsClient) GetGoogleID() string {
	args := m.Called()
	return args.String(0)
}
