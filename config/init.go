package config

import (
	"clean_embassy_helper/conf"
	"clean_embassy_helper/services"
	"clean_embassy_helper/usecases"
)

//type Handler struct {
//	GetEmbassiesInHostCountry func(homeCountry, hostCountry string) ([]models.Embassy, error)
//}

type Dependencies struct {
	//TODO: make deps private once there's an initHandler() function
	EmbassyService *services.EmbassyService
	MgoService     *services.MgoService
	GoogleService  *services.GoogleService
}

//
//func InitHandlers() (Handler, error) {
//	deps, err := InitDependencies()
//	if err != nil {
//		return Handler{}, err
//	}
//	return Handler{
//		GetEmbassiesInHostCountry: deps.EmbassyService.GetEmbassies,
//	}, nil
//}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()

	embassyUsecase := usecases.NewEmbassyUsecase(cfg.Domain)
	mgoUsecase := usecases.NewInsertUseCase(cfg.Mgo)
	googleUsecase := usecases.NewGoogleUsecase(cfg.ApiKey)

	return Dependencies{
		EmbassyService: services.NewEmbassyService(embassyUsecase),
		MgoService:     services.NewMgoService(mgoUsecase),
		GoogleService:  services.NewGoogleService(googleUsecase),
	}, nil
}

//func InitHandlers() {
//
//}
