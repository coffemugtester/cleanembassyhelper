package main

import (
	"clean_embassy_helper/config"
	"clean_embassy_helper/handlers"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
)

// TODO: behavior vs layers
var rootCmd = &cobra.Command{
	Use:   "angry-embassy",
	Short: "Get embassies for a given home and host country",
	Long:  "This application demonstrates passing parameters through the terminal using Cobra.",
}

var getEmbassies = &cobra.Command{
	Use:   "getembassies",
	Short: "Fetch embassies for a given home and host country",
	Long:  "This command takes two parameters and passes them to the handler function.",
	Run: func(cmd *cobra.Command, args []string) {

		// Retrieve the values of the flags
		home, _ := cmd.Flags().GetString("home")
		host, _ := cmd.Flags().GetString("host")

		fmt.Println("Dependencies initialized")

		deps, err := config.InitDependencies()
		if err != nil {
			fmt.Printf("config.InitDependencies error: %v\n", err)
			return
		}

		// Print the formatted string
		fmt.Printf("Fetching embassies for Home Country: %s and Host Country: %s\n", home, host)
		handlers.WriteEmbassies(deps, &home, &host)
	},
}

func init() {
	fmt.Println("Initializing Cobra CLI")
	rootCmd.AddCommand(getEmbassies)

	// Register flags
	getEmbassies.Flags().String("home", "", "Home country")
	getEmbassies.Flags().String("host", "", "Host country")

	// Mark flags as required if needed
	getEmbassies.MarkFlagRequired("home")
	getEmbassies.MarkFlagRequired("host")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

}

//TODO: remove parseFlags() function

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
