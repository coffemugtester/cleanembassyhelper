package scraper

import "github.com/gocolly/colly/v2"

type ColyClient struct {
	collector *colly.Collector
	BaseURL   string
	//api_key
}

func NewColyClient(baseURL string) *ColyClient {
	c := colly.NewCollector()
	return &ColyClient{
		collector: c,
		BaseURL:   baseURL,
	}
}
