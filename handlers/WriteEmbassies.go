package handlers

import (
	"clean_embassy_helper/config"
	"fmt"
	"net/url"
)

//TODO: put this in a handler

func WriteEmbassies(deps config.Dependencies, homeCountry *string, hostCountry *string) {
	fmt.Printf("WriteEmbassies ----- Home country: %s\nHost country: %s\n", *homeCountry, *hostCountry)
	embassies, err := deps.EmbassyService.GetEmbassies(*homeCountry, *hostCountry)
	if err != nil {
		fmt.Printf("embassyService.GetEmbassies error: %v\n", err)
		return
	}

	for _, embassy := range embassies {

		parsedURL, err := url.Parse(embassy.MapLink)
		if err != nil {
			fmt.Printf("url.Parse error: %v\n", err)
			return
		}

		queryParams := parsedURL.Query()
		qParam := queryParams.Get("q")

		fmt.Printf("Query param: %s\n", qParam)

		embassy.GoogleID = deps.GoogleService.GetGoogleID(qParam)
		embassyDetails, err := deps.GoogleService.GetPlaceDetails(embassy.GoogleID)
		if err != nil {
			fmt.Printf("googleService.GetPlaceDetails error: %v\n", err)
			return
		}

		embassy.PlaceDetails = embassyDetails
		//TODO: this outputs data and whoever requests it either stores it or displays it ...

		//TODO: decide if scraped data should be in a different table than requested data
		insertedID, err := deps.MgoService.InsertDocument(embassy)
		if err != nil {
			fmt.Printf("mgoService.InsertDocument error: %v\n", err)
			return
		}

		fmt.Printf("Embassy: %v\n", embassy)
		fmt.Printf("Inserted ID: %s\n", insertedID)
	}

}
