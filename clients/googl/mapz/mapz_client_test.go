package mapz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapsImpl_GetGoogleID(t *testing.T) {
	m := MockMapsClient{}

	m.On("GetGoogleID").Return("123")
	assert.Equal(t, "mapz", m.GetGoogleID())
}
