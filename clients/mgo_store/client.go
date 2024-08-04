package mgo_store

import (
	"clean_embassy_helper/clients/mgo_store/crud"
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
)

type Client struct {
	*crud.Client
}

func NewClient(mgo conf.MgoConfig) *Client {
	crudClient := crud.NewCRUDClient(mgo, crud.MongoNonInter{})
	return &Client{
		crudClient,
	}
}

func (c Client) InsertDocument(document models.Embassy) (string, error) {
	return crud.InsertDocument(c.Client, document)
}
