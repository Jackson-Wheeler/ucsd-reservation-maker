package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	WELCOME_MESSAGE = "Welcome to the UCSD Reservation maker!\n"
)

func main() {
	// introductory message
	fmt.Printf("\n%s\n", WELCOME_MESSAGE)

	// parse command line arguments
	if len(os.Args) < 2 {
		log.Fatal("Error: please provide a configuration file path\nusage: ./ucsd-reservation-maker <config file path>")
	}

	// read configuration file
	configFilePath := os.Args[1]
	fmt.Printf("reading configuration details from: %s\n", configFilePath)
	config, err := parseConfigFile(configFilePath)
	if err != nil {
		log.Fatal("Error: failed to parse configuration file:", err)
	}

	// read environment variables
	siteCredentials := readEnvVariables()

	// make reservation
	MakeReservation(config, siteCredentials)
}

func readEnvVariables() SiteCredentials {
	godotenv.Load(".env")

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username == "" || password == "" {
		log.Fatal("Error: please set USERNAME and PASSWORD environment variables in .env file")
	}
	return SiteCredentials{username: username, password: password}
}
