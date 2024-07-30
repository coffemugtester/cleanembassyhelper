package main

import (
	"clean_embassy_helper/clients/coly"
	"clean_embassy_helper/usecases"
	"flag"
	"fmt"
)

func main() {

	homeCountry, hostCountry, err := parseFlags()
	if err != nil {
		fmt.Printf("parseFlags error: %v\n", err)
		return
	}
	fmt.Printf("Home country: %s\nHost country: %s\n", *homeCountry, *hostCountry)

	colyClient := coly.NewClient()

	scrappy := usecases.NewScraper(colyClient)
	embassies := scrappy.GetEmbassies(*homeCountry, *hostCountry)
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
