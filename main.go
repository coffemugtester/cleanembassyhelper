package main

import (
	"clean_embassy_helper/services"
	"flag"
	"fmt"
)

func main() {

	//TODO: use init function to load config and shared resources
	//TODO: env var for root path

	homeCountry, hostCountry, err := parseFlags()
	if err != nil {
		fmt.Printf("parseFlags error: %v\n", err)
		return
	}
	fmt.Printf("Home country: %s\nHost country: %s\n", *homeCountry, *hostCountry)

	embassyService := services.NewEmbassyService()
	embassies, err := embassyService.GetEmbassies(*homeCountry, *hostCountry)
	if err != nil {
		fmt.Printf("embassyService.GetEmbassies error: %v\n", err)
		return
	}
	fmt.Printf("Embassies: %v\n", embassies)
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
