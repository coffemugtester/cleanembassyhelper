package config

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/services"
	"clean_embassy_helper/usecases"
)

type Dependencies struct {
	//TODO: make deps private once there's an initHandler() function
	EmbassyService *services.EmbassyService
	Mgo            *services.MgoService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()

	embassyUsecase := usecases.NewEmbassyUsecase(cfg.Domain)
	mgoUsecase := usecases.NewInsertUseCase(cfg.Mgo)

	return Dependencies{
		EmbassyService: services.NewEmbassyService(embassyUsecase),
		Mgo:            services.NewMgoService(mgoUsecase),
	}, nil
}

//func InitHandlers() {
//
//}
