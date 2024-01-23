package webdriver

import (
	"fmt"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, time myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {
	const CREATE_RESERVATION_BTN_TEXT = "Create A Reservation"
	const STUDY_ROOM_BOOKING_BTN_XPATH = `//*[@aria-label='Book Now With The "Reserve Spaces | Study Rooms & Open Desk" Template']`

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", time.Date, time.StartTime, time.EndTime)

	// click the create reservation button
	myClickElement(driver, selenium.ByLinkText, CREATE_RESERVATION_BTN_TEXT)

	myClickElement(driver, selenium.ByXPATH, STUDY_ROOM_BOOKING_BTN_XPATH)

}
