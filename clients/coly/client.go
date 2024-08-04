package coly

import (
	"clean_embassy_helper/clients/coly/agent"
	"clean_embassy_helper/internal/models"
)

type Client struct {
	colyClient *agent.ColyClient
}

func NewClient(client *agent.ColyClient) *Client {
	return &Client{
		colyClient: client,
	}
}

func (c *Client) GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error) {
	return agent.GetEmbassies(c.colyClient, homeCountry, hostCountry)
}
