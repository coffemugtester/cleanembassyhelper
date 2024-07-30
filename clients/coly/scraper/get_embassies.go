package scraper

import (
	"clean_embassy_helper/internal/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
)

func GetEmbassies(client *ColyClient, homeCountry, hostCountry string) []models.Embassy {

	var result []models.Embassy

	url := fmt.Sprintf("%s/en/%s/embassy/%s/", client.BaseURL, strings.ToLower(homeCountry), strings.ToLower(hostCountry))
	fmt.Printf("URL: %s\n", url)

	client.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", url)
	})

	client.collector.OnHTML(".embassy__list", func(e *colly.HTMLElement) {
		// Create a GoQuery document from the Colly element
		doc := e.DOM

		// Iterate over all elements with the class "embassy__name"
		doc.Find(".embassy__name").Each(func(i int, s *goquery.Selection) {
			var embassy models.Embassy

			embassyName := s.Text()

			// Find the sibling element for the map link
			// Since "embassy__map-link" is in a sibling "embassy__list__wrap"
			mapLink := s.Next().Find(".embassy__list__wrap .embassy__map-link").AttrOr("href", "")

			fmt.Printf("Name: \"%s\",\n", embassyName)
			fmt.Printf("MapLink: \"%s\",\n", mapLink)

			embassy.HostCountry = hostCountry
			embassy.HomeCountry = homeCountry
			embassy.Name = strings.Replace(embassyName, "\n ", "", -1)
			embassy.MapLink = strings.Replace(mapLink, "\n+", "", -1)
			embassy.City = extractLastWord(embassyName)
			result = append(result, embassy)
		})
	})

	err := client.collector.Visit(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}

	if len(result) == 0 {
		fmt.Errorf("no embassies found for %s in %s", homeCountry, hostCountry)
	}

	return result
}

func extractLastWord(s string) string {
	// Remove leading and trailing whitespace
	s = strings.TrimSpace(s)
	// Split the string into words
	words := strings.Fields(s)
	// Return the last word
	if len(words) > 0 {
		return words[len(words)-1]
	}
	return ""
}
