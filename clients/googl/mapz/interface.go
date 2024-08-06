package mapz

type MapsClient interface {
	// GetGoogleID returns the Google id of the location
	GetGoogleID(placeQuery string) string
}
