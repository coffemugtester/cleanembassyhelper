package crud

import (
	"clean_embassy_helper/internal/conf"
	"clean_embassy_helper/internal/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//TODO: use mock db for testing

var cfg = conf.LoadConfig()

func TestInsertDocument(t *testing.T) {

	client, err := NewMongoClient(cfg.MongoUri)
	if err != nil {
		return
	}

	insertedDoc, err := InsertDocument(client, models.Embassy{})
	fmt.Printf("insertedDoc: %v\n", insertedDoc)
	assert.Nil(t, err)
	assert.NotNil(t, insertedDoc)
}

func TestMongoClient(t *testing.T) {

	if _, err := NewMongoClient(cfg.MongoUri); err != nil {
		t.Error("Expected true, got false")
		assert.Nil(t, err)
	}
}
