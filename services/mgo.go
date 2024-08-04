package services

import "clean_embassy_helper/internal/models"

var _ MgoClientUsecase = (*MgoService)(nil)

type MgoService struct {
	mgoClientUseCase MgoClientUsecase
}

func NewMgoService(mgoClientUseCase MgoClientUsecase) *MgoService {
	return &MgoService{
		mgoClientUseCase: mgoClientUseCase,
	}
}

func (m *MgoService) InsertDocument(document models.Embassy) (string, error) {
	return m.mgoClientUseCase.InsertDocument(document)
}
