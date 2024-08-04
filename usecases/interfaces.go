package usecases

import "clean_embassy_helper/internal/models"

type ColyClient interface {
	GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error)
}

type MgoClient interface {
	InsertDocument(models.Embassy) (string, error)
}
