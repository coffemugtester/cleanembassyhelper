package usecases

import (
	"clean_embassy_helper/clients/coly"
	"clean_embassy_helper/internal/models"
)

//TODO: learn about init function

var _ ColyClient = (*EmbassyUsecase)(nil)

type EmbassyUsecase struct {
	colyClient ColyClient
}

func NewEmbassyUsecase() *EmbassyUsecase {
	colyClient := coly.NewClient()
	return &EmbassyUsecase{
		colyClient: colyClient,
	}
}

func (e *EmbassyUsecase) GetEmbassies(homeCountry string, hostCountry string) ([]models.Embassy, error) {
	return e.colyClient.GetEmbassies(homeCountry, hostCountry)
}
