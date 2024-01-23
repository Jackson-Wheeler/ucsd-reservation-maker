package myconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// -- Configuration File Struct Definitions --
type Config struct {
	ReservationDetails  ReservationDetails `yaml:"reservation_details"`
	RoomPreferenceOrder []string           `yaml:"room_preference_order"`
	ReservationTimes    []ReservationTime  `yaml:"reservation_times"`
}

type ReservationDetails struct {
	NumPeople    int    `yaml:"num_people"`
	EventName    string `yaml:"event_name"`
	EventType    string `yaml:"event_type"`
	ContactName  string `yaml:"contact_name"`
	ContactPhone string `yaml:"contact_phone"`
	ContactEmail string `yaml:"contact_email"`
	Description  string `yaml:"description"`
}

type ReservationTime struct {
	Date      string `yaml:"date"`
	StartTime string `yaml:"startTime"`
	EndTime   string `yaml:"endTime"`
}

// Parse Configuration File
func ParseConfigFile(configFilePath string) (Config, error) {
	var config Config

	// Read the YAML file
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, fmt.Errorf("failed to read YAML file: %v", err)
	}

	// Parse the YAML file into the Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse YAML: %v", err)
	}

	// Print the parsed information
	// TODO: remove this
	fmt.Printf("Reservation Details:\n%+v\n\n", config.ReservationDetails)
	fmt.Printf("Room Preference Order:\n%+v\n\n", config.RoomPreferenceOrder)
	fmt.Printf("Reservation Times:\n%+v\n\n", config.ReservationTimes)

	return config, nil
}
