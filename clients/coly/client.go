package coly

import (
	"clean_embassy_helper/clients/coly/scraper"
	"clean_embassy_helper/internal/conf"
	"clean_embassy_helper/internal/models"
)

var cfg = conf.LoadConfig()

type Client struct {
	colyClient *scraper.ColyClient
}

func NewClient() *Client {
	client := scraper.NewColyClient(cfg.Domain)
	return &Client{
		colyClient: client,
	}
}

func (c *Client) GetEmbassies(homeCountry, hostCountry string) ([]models.Embassy, error) {
	return scraper.GetEmbassies(c.colyClient, homeCountry, hostCountry)
}
