package main

import (
	"clean_embassy_helper/config"
	"flag"
	"fmt"
	"net/url"
)

func main() {

	//TODO: add controller

	homeCountry, hostCountry, err := parseFlags()
	if err != nil {
		fmt.Printf("parseFlags error: %v\n", err)
		return
	}
	fmt.Printf("Home country: %s\nHost country: %s\n", *homeCountry, *hostCountry)

	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	embassies, err := deps.EmbassyService.GetEmbassies(*homeCountry, *hostCountry)
	if err != nil {
		fmt.Printf("embassyService.GetEmbassies error: %v\n", err)
		return
	}

	parsedURL, err := url.Parse(embassies[0].MapLink)
	if err != nil {
		fmt.Printf("url.Parse error: %v\n", err)
		return
	}

	queryParams := parsedURL.Query()
	qParam := queryParams.Get("q")

	fmt.Printf("Query param: %s\n", qParam)

	embassies[0].GoogleID = deps.GoogleService.GetGoogleID(qParam)

	//TODO: decide if scraped data should be in a different table than requested data

	insertedID, err := deps.MgoService.InsertDocument(embassies[0])
	if err != nil {
		fmt.Printf("mgoService.InsertDocument error: %v\n", err)
		return
	}

	fmt.Printf("Embassies: %v\n", embassies)
	fmt.Printf("Embassy: %v\n", embassies[0])
	fmt.Printf("Inserted ID: %s\n", insertedID)
}

func parseFlags() (*string, *string, error) {
	homeCountry := flag.String("home", "", "The country where the embassy is located")
	hostCountry := flag.String("host", "", "The country represented by the embassy")

	flag.Parse()

	if *homeCountry == "" || *hostCountry == "" {
		fmt.Println("Both home and host country are required.")
		flag.Usage()
	}

	if *homeCountry == *hostCountry {
		fmt.Println("Home and host country cannot be the same.")
		flag.Usage()
		return nil, nil, fmt.Errorf("home and host country cannot be the same")
	}

	return homeCountry, hostCountry, nil
}
