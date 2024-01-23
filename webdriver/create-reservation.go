package webdriver

import (
	"fmt"
	"log"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, time myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {
	const (
		CREATE_RESERVATION_BTN_TEXT  = "Create A Reservation"
		STUDY_ROOM_BOOKING_BTN_XPATH = `//*[@aria-label='Book Now With The "Reserve Spaces | Study Rooms & Open Desk" Template']`
		BOOOKING_DATE_INPUT_ID       = "booking-date-input"
		START_TIME_INPUT_ID          = "start-time-input"
		END_TIME_INPUT_ID            = "end-time-input"
		SEARCH_BTN_XPATH             = `//button[normalize-space()='Search']`
		ROOM_ITEM_CLASS              = "room-column column"
	)

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", time.Date, time.StartTime, time.EndTime)

	// click the create reservation button
	myClickElement(driver, selenium.ByLinkText, CREATE_RESERVATION_BTN_TEXT)

	// click the study room booking button
	myClickElement(driver, selenium.ByXPATH, STUDY_ROOM_BOOKING_BTN_XPATH)

	// input the booking date
	myClearAndSendKeys(driver, selenium.ByID, BOOOKING_DATE_INPUT_ID, time.Date)

	// input the start time
	myClearAndSendKeys(driver, selenium.ByID, START_TIME_INPUT_ID, time.StartTime)

	// input the end time
	myClearAndSendKeys(driver, selenium.ByID, END_TIME_INPUT_ID, time.EndTime)

	// click the search button
	myClickElement(driver, selenium.ByXPATH, SEARCH_BTN_XPATH)

	// get all room items
	roomItems := myFindElements(driver, selenium.ByClassName, ROOM_ITEM_CLASS)

	for _, roomItem := range roomItems {
		text, err1 := roomItem.Text()
		outerHTML, err2 := roomItem.GetAttribute("outerHTML")
		fmt.Printf("Room Item Text: %s\n", text)
		fmt.Printf("Room Item OuterHTML: %s\n", outerHTML)
		if err1 != nil || err2 != nil {
			log.Fatal("Error: failed to get room item text or outerHTML - ", err1, err2)
		}
	}
}
