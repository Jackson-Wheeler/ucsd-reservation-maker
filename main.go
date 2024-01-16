package main

import (
	"fmt"
	"os"
)

const (
	WELCOME_MESSAGE = "Welcome to the UCSD Reservation maker!\n"
)

func main() {
	// introductory message
	fmt.Printf("\n%s\n", WELCOME_MESSAGE)

	// parse command line arguments
	if len(os.Args) < 2 {
		fmt.Println("usage: ./ucsd-reservation-maker <config file path>")
		fmt.Println("please provide a configuration file path")
		os.Exit(1)
	}

	// read configuration file
	configFilePath := os.Args[1]
	fmt.Printf("reading configuration details from: %s\n", configFilePath)
	_, err := parseConfigFile(configFilePath)
	if err != nil {
		fmt.Printf("failed to parse configuration file: %v\n", err)
		os.Exit(1)
	}

}
