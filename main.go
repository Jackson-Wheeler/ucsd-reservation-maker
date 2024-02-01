package main

import (
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

// ARGUMENTS: [configFilePath], where configFilePath = DEFAULT_CONFIG_FILE by default

func main() {
	// introductory message
	fmt.Printf("\n%s\n", WELCOME_MESSAGE)

	// read configuration file
	configFilePath := DEFAULT_CONFIG_FILE
	if len(os.Args) >= 2 {
		configFilePath = os.Args[1]
	}
	fmt.Printf("Reading configuration details from: '%s'\n", configFilePath)
	config, err := myconfig.ParseConfigFile(configFilePath)
	if err != nil {
		log.Fatal("Error: failed to parse configuration file:", err)
	}

	// read environment variables
	siteCredentials := readEnvVariables()

	// make reservation
	reservationMaker.MakeReservations(config, siteCredentials)
}

func readEnvVariables() reservationMaker.SiteCredentials {
	godotenv.Load("env")

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username == "" || password == "" {
		log.Fatal("Error: please set USERNAME and PASSWORD environment variables in .env file")
	}
	return reservationMaker.SiteCredentials{Username: username, Password: password}
}
