package googl

import "clean_embassy_helper/clients/googl/mapz"

type Client struct {
	*mapz.Client
}

func NewClient(apiKey string) *Client {
	mapzClient := mapz.NewMapzClient(apiKey)
	return &Client{
		mapzClient,
	}
}

func (c Client) GetGoogleID(placeQuery string) string {
	return c.Client.GetGoogleID(placeQuery)
}
