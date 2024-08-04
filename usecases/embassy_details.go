package usecases

import (
	"clean_embassy_helper/clients/coly"
	"clean_embassy_helper/internal/models"
)

var _ ColyClient = (*EmbassyUsecase)(nil)

type EmbassyUsecase struct {
	scraper ColyClient
}

func NewEmbassyUsecase() *EmbassyUsecase {
	colyClient := coly.NewClient()
	return &EmbassyUsecase{
		scraper: colyClient,
	}
}

func (e *EmbassyUsecase) GetEmbassies(homeCountry string, hostCountry string) ([]models.Embassy, error) {
	return e.scraper.GetEmbassies(homeCountry, hostCountry)
}
