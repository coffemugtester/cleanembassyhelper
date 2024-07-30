package services

import "clean_embassy_helper/internal/models"

type EmbassyUseCase interface {
	GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error)
}
