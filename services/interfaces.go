package services

import "clean_embassy_helper/internal/models"

type EmbassyUsecase interface {
	GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error)
}

type MgoClientUsecase interface {
	InsertDocument(models.Embassy) (string, error)
}
