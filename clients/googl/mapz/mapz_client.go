package mapz

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// implementation check
var _ MapsClient = (*Client)(nil)

// Client mapzImplementation
type Client struct {
	apiKey string
}

func (c Client) GetGoogleID(placeQuery string) (id string) {

	placeQuery = url.QueryEscape(placeQuery)

	mapsURL := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=" + placeQuery + "&inputtype=textquery&fields=place_id&key=" + c.apiKey

	fmt.Printf("placeQuery: %v\n", placeQuery)

	fmt.Printf("apiKey: %v\n", c.apiKey)

	fmt.Printf("Encoded URL: %v\n", mapsURL)

	response, err := http.Get(mapsURL)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	var item struct {
		Candidates []struct {
			PlaceID string `json:"place_id"`
		}
	}

	if err := json.Unmarshal(body, &item); err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	id = item.Candidates[0].PlaceID

	return id
}

// mapzImplementation methods

// NewMapzClient creates a new mapz client
func NewMapzClient(apiKey string) *Client { // add apikey as a parameter
	return &Client{
		apiKey: apiKey,
	}
}
