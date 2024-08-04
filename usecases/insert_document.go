package usecases

import (
	"clean_embassy_helper/clients/mgo_store"
	"clean_embassy_helper/conf"
	"clean_embassy_helper/internal/models"
)

var _ MgoClient = (*InsertUseCase)(nil)

type InsertUseCase struct {
	mgoClient MgoClient
}

func NewInsertUseCase(mgoConf conf.MgoConfig) *InsertUseCase {
	mgoClient := mgo_store.NewClient(mgoConf)
	return &InsertUseCase{
		mgoClient: mgoClient,
	}
}

func (i *InsertUseCase) InsertDocument(document models.Embassy) (string, error) {
	return i.mgoClient.InsertDocument(document)
}
