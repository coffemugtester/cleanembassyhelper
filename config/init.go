package config

import (
	"clean_embassy_helper/clients/coly"
	"clean_embassy_helper/clients/coly/scraper"
	"clean_embassy_helper/internal/conf"
	"clean_embassy_helper/services"
	"clean_embassy_helper/usecases"
)

type Dependencies struct {
	//TODO: make deps private once there's an initHandler() function
	EmbassyService *services.EmbassyService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()

	scraperColyClient := scraper.NewColyClient(cfg.Domain)
	colyClient := coly.NewClient(scraperColyClient)
	embassyUsecase := usecases.NewEmbassyUsecase(colyClient)

	return Dependencies{
		EmbassyService: services.NewEmbassyService(embassyUsecase),
	}, nil
}

//func InitHandlers() {
//
//}
