package usecases

import "clean_embassy_helper/internal/models"

//TODO: learn about init function

var _ ColyClient = (*EmbassyDetails)(nil)

type EmbassyDetails struct {
	colyClient ColyClient
}

func NewScraper(colyClient ColyClient) *EmbassyDetails {
	return &EmbassyDetails{
		colyClient: colyClient,
	}
}

func (e *EmbassyDetails) GetEmbassies(homeCountry string, hostCountry string) ([]models.Embassy, error) {
	return e.colyClient.GetEmbassies(homeCountry, hostCountry)
}
