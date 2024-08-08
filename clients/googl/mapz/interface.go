package mapz

import "clean_embassy_helper/internal/models"

type MapsClient interface {
	// GetGoogleID returns the Google id of the location
	GetGoogleID(placeQuery string) string
	GetPlaceDetails(placeQuery string) (models.PlaceDetails, error)
}
