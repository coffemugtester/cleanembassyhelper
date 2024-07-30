package services

import (
	"clean_embassy_helper/internal/models"
	"clean_embassy_helper/usecases"
)

//TODO:

var _ EmbassyUseCase = (*EmbassyService)(nil)

type EmbassyService struct {
	embassyUsecase EmbassyUseCase
}

func NewEmbassyService() *EmbassyService {
	embassyUsecase := usecases.NewEmbassyUsecase()
	return &EmbassyService{
		embassyUsecase: embassyUsecase,
	}
}

func (e *EmbassyService) GetEmbassies(homeCountry string, hostCountry string) ([]models.Embassy, error) {
	return e.embassyUsecase.GetEmbassies(homeCountry, hostCountry)
}
