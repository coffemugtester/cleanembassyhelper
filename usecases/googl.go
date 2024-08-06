package usecases

import "clean_embassy_helper/clients/googl/mapz"

var _ MapsClient = (*GoogleUsecase)(nil)

type GoogleUsecase struct {
	client MapsClient
}

func NewGoogleUsecase(apiKey string) *GoogleUsecase {
	mapzClient := mapz.NewMapzClient(apiKey)
	return &GoogleUsecase{
		client: mapzClient,
	}
}

func (g *GoogleUsecase) GetGoogleID(placeQuery string) string {
	return g.client.GetGoogleID(placeQuery)
}
