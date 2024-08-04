package services

import (
	"clean_embassy_helper/internal/models"
)

var _ EmbassyUsecase = (*EmbassyService)(nil)

type EmbassyService struct {
	embassyUsecase EmbassyUsecase
}

func NewEmbassyService(embassyUsecase EmbassyUsecase) *EmbassyService {
	return &EmbassyService{
		embassyUsecase: embassyUsecase,
	}
}

func (e *EmbassyService) GetEmbassies(homeCountry string, hostCountry string) ([]models.Embassy, error) {
	return e.embassyUsecase.GetEmbassies(homeCountry, hostCountry)
}
