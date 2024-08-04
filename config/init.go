package config

import (
	"clean_embassy_helper/clients/coly"
	"clean_embassy_helper/clients/coly/agent"
	"clean_embassy_helper/conf"
	"clean_embassy_helper/services"
	"clean_embassy_helper/usecases"
)

type Dependencies struct {
	//TODO: make deps private once there's an initHandler() function
	EmbassyService *services.EmbassyService
	//Mgo 		  *services.MgoService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()

	scraperColyClient := agent.NewColyClient(cfg.Domain)
	colyClient := coly.NewClient(scraperColyClient)
	embassyUsecase := usecases.NewEmbassyUsecase(colyClient)

	return Dependencies{
		EmbassyService: services.NewEmbassyService(embassyUsecase),
	}, nil
}

//func InitHandlers() {
//
//}
