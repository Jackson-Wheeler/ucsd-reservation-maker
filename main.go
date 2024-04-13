package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker"
	"github.com/joho/godotenv"
)

const (
	WELCOME_MESSAGE     = "Welcome to the UCSD Reservation maker!\n"
	DEFAULT_CONFIG_FILE = "config.yaml"
)

// Can be run with the following flags:
// -r: make reservations according to the config file
// -o: open reservations page (logs in and opens the 'Reserve Space' view at the date and time of the first reservation in the config file). Does not make any reservations. Purpose: for viewing what dates and times are available.

func main() {
	// Parse command line flags
	openFlag, err := parseCommandLineFlags() // -r is default mode, -o (openFlag) is for only opening reservations page
	if err != nil {
		log.Fatalf("Error parsing command line flags: %v\n", err)
	}

	// introductory message
	fmt.Printf("\n%s\n", WELCOME_MESSAGE)

	// read configuration file
	configFilePath := DEFAULT_CONFIG_FILE
	fmt.Printf("Reading configuration details from: '%s'\n", configFilePath)
	config, err := myconfig.ParseConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	// read environment variables
	siteCredentials, err := readEnvVariables()
	if err != nil {
		log.Fatalf("Error getting site credentials: %v\n", err)
	}

	// make reservation / open reservations page
	err = reservationMaker.MakeReservations(*config, *siteCredentials, openFlag)
	if err != nil {
		log.Fatalf("Error making reservations: %v\n", err)
	}
}

func parseCommandLineFlags() (bool, error) {
	reservationFlag := flag.Bool("r", false, "make reservations according to the config file")
	openFlag := flag.Bool("o", false, "open reservations page (logs in and opens the 'Reserve Space' view at the date and time of the first reservation in the config file). Does not make any reservations. Purpose: for viewing what dates and times are available.")
	flag.Parse()

	if !*reservationFlag && !*openFlag {
		fmt.Printf("\nPlease specify a flag for the type of operation you would like to perform: -r (make reservations) or -o (open reservations page)\n")
		fmt.Printf("Example: %s -r\n", os.Args[0])
		fmt.Println()
		flag.Usage()
		fmt.Println()
		return false, fmt.Errorf("no flag specified")
	}

	if *openFlag {
		fmt.Println("-o mode: Opening reservations page only (not making reservations)...")
	} else if *reservationFlag {
		fmt.Println("-r mode (default): Making reservations...")
	}

	return *openFlag, nil
}

func readEnvVariables() (*reservationMaker.SiteCredentials, error) {
	godotenv.Load("env")

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username == "" || password == "" {
		return nil, fmt.Errorf("please set USERNAME and PASSWORD environment variables in .env file")
	}
	return &reservationMaker.SiteCredentials{Username: username, Password: password}, nil
}
