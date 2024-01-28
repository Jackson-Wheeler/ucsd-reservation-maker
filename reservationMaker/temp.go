package reservationMaker

import (
	"log"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/webdriver"
	"github.com/tebeka/selenium"
)

// begin booking: create reservation btn & booking type btn
func beginBooking(driver selenium.WebDriver, bookingType int) {
	// click the create reservation button
	webdriver.ClickElement(driver, CREATE_RESERVATION_BTN_BY, CREATE_RESERVATION_BTN_VAL)

	// click the 'book now' button for the specified booking type
	switch bookingType {
	case BOOKING_TYPE_STUDENT_ORGS:
		webdriver.ClickElement(driver, BOOKING_TYPE_BTN_STUDENT_ORGS_BY, BOOKING_TYPE_BTN_STUDENT_ORGS_VAL)
	case BOOKING_TYPE_STUDY_ROOM:
		webdriver.ClickElement(driver, BOOKING_TYPE_BTN_STUDY_ROOM_BY, BOOKING_TYPE_BTN_STUDY_ROOM_VAL)
	default:
		log.Fatalf("Error: invalid booking type: %d", bookingType)
	}
}

// set reservation time: booking date, start time, end time, click search
func setReservationTime(driver selenium.WebDriver, resTime myconfig.ReservationTime) {
	// wait for content to load
	webdriver.WaitForElementReady(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL)

	// input the booking date
	webdriver.ClearAndSendKeys(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL, resTime.Date)
	webdriver.SendKeys(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL, selenium.TabKey)

	// input the start time
	webdriver.ClearAndSendKeys(driver, START_TIME_INPUT_BY, START_TIME_INPUT_VAL, resTime.StartTime)

	// input the end time
	webdriver.ClearAndSendKeys(driver, END_TIME_INPUT_BY, END_TIME_INPUT_VAL, resTime.EndTime)

	// click the search button
	webdriver.ClickElement(driver, SEARCH_BTN_BY, SEARCH_BTN_VAL)
}
