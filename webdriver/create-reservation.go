package webdriver

import (
	"fmt"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, time myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {
	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", time.Date, time.StartTime, time.EndTime)
}
