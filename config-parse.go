package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// -- Configuration File Struct Definitions --
type Config struct {
	ReservationDetails       ReservationDetails         `yaml:"reservation_details"`
	RoomPreferenceOrder      []string                   `yaml:"room_preference_order"`
	ReservationDatesAndTimes []ReservationDatesAndTimes `yaml:"reservation_dates_and_times"`
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

type ReservationDatesAndTimes struct {
	StartOfWeekDate string                        `yaml:"start_of_week_date"`
	Reservations    map[DayOfWeek]ReservationTime `yaml:"reservations"`
}

type ReservationTime struct {
	StartTime string `yaml:"startTime"`
	EndTime   string `yaml:"endTime"`
}

// Define Day of the Week
type DayOfWeek string

const (
	Sunday    DayOfWeek = "Sunday"
	Monday    DayOfWeek = "Monday"
	Tuesday   DayOfWeek = "Tuesday"
	Wednesday DayOfWeek = "Wednesday"
	Thursday  DayOfWeek = "Thursday"
	Friday    DayOfWeek = "Friday"
	Saturday  DayOfWeek = "Saturday"
)

func isValidDayOfWeek(day string) bool {
	day = strings.ToLower(day)
	switch day {
	case string(Sunday), string(Monday), string(Tuesday), string(Wednesday), string(Thursday), string(Friday), string(Saturday):
		return true
	default:
		return false
	}
}

// Parse Configuration File
func parseConfigFile(configFilePath string) (Config, error) {
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
	fmt.Printf("Reservation Dates and Times:\n%+v\n\n", config.ReservationDatesAndTimes)

	return config, nil
}
