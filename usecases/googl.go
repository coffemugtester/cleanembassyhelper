package usecases

import (
	"clean_embassy_helper/clients/googl"
	"clean_embassy_helper/internal/models"
)

var _ MapsClient = (*GoogleUsecase)(nil)

type GoogleUsecase struct {
	client MapsClient
}

func NewGoogleUsecase(apiKey string) *GoogleUsecase {
	client := googl.NewClient(apiKey)
	return &GoogleUsecase{
		client: client,
	}
}

func (g *GoogleUsecase) GetGoogleID(placeQuery string) string {
	return g.client.GetGoogleID(placeQuery)
}

func (g *GoogleUsecase) GetPlaceDetails(placeQuery string) (models.PlaceDetails, error) {
	return g.client.GetPlaceDetails(placeQuery)
}
