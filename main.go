package main

import (
	"clean_embassy_helper/config"
	"flag"
	"fmt"
)

func main() {

	//TODO: env var for root path

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

	insertedID, err := deps.Mgo.InsertDocument(embassies[0])
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
