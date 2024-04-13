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
func ParseConfigFile(configFilePath string) (*Config, error) {
	var config *Config

	// Read the YAML file
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %v", err)
	}

	// Parse the YAML file into the Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration file: %v", err)
	}

	// Validate the parsed information
	err = validateConfig(*config)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration file: %v", err)
	}

	// Print the parsed information
	fmt.Println("--- Configuration Details ---")
	fmt.Printf("Reservation Details:\n%+v\n\n", config.ReservationDetails)
	fmt.Printf("Room Preference Order:\n%+v\n\n", config.RoomPreferenceOrder)
	fmt.Printf("Reservation Times:\n%+v\n", config.ReservationTimes)
	fmt.Printf("--- End: Configuration Details ---\n\n")

	return config, nil
}

// Validate the parsed configuration. Return error if invalid
func validateConfig(config Config) error {
	// Validate the ReservationDetails
	if config.ReservationDetails.NumPeople <= 0 {
		return fmt.Errorf("num_people must be greater than 0")
	}
	if config.ReservationDetails.EventName == "" {
		return fmt.Errorf("event_name must be non-empty")
	}
	if config.ReservationDetails.ContactName == "" {
		return fmt.Errorf("contact_name must be non-empty")
	}
	if config.ReservationDetails.ContactPhone == "" {
		return fmt.Errorf("contact_phone must be non-empty")
	}
	if config.ReservationDetails.ContactEmail == "" {
		return fmt.Errorf("contact_email must be non-empty")
	}
	if config.ReservationDetails.Description == "" {
		return fmt.Errorf("description must be non-empty")
	}

	// Validate the RoomPreferenceOrder
	if len(config.RoomPreferenceOrder) == 0 {
		return fmt.Errorf("room_preference_order must be non-empty")
	}

	// Validate the ReservationTimes
	if len(config.ReservationTimes) == 0 {
		return fmt.Errorf("reservation_times must be non-empty")
	}
	for _, reservationTime := range config.ReservationTimes {
		if reservationTime.Date == "" {
			return fmt.Errorf("reservation_times date must be non-empty")
		}
		if reservationTime.StartTime == "" {
			return fmt.Errorf("reservation_times startTime must be non-empty")
		}
		if reservationTime.EndTime == "" {
			return fmt.Errorf("reservation_times endTime must be non-empty")
		}
	}

	return nil
}
