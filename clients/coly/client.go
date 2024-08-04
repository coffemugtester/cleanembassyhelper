package coly

import (
	"clean_embassy_helper/clients/coly/agent"
	"clean_embassy_helper/internal/models"
)

type Client struct {
	colyAgent *agent.ColyClient
}

func NewClient(domain string) *Client {
	ColyClientAgent := agent.NewColyClient(domain)
	return &Client{
		colyAgent: ColyClientAgent,
	}
}

func (c *Client) GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error) {
	return agent.GetEmbassies(c.colyAgent, homeCountry, hostCountry)
}
