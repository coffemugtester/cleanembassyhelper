package services

var _ GoogleUsecase = (*GoogleService)(nil)

type GoogleService struct {
	googleUsecase GoogleUsecase
}

func NewGoogleService(googleUsecase GoogleUsecase) *GoogleService {
	return &GoogleService{
		googleUsecase: googleUsecase,
	}
}

func (g *GoogleService) GetGoogleID(placeQuery string) string {
	return g.googleUsecase.GetGoogleID(placeQuery)
}
