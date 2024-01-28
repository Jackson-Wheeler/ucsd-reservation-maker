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
	webdriver.FindAndClickElement(driver, CREATE_RESERVATION_BTN_BY, CREATE_RESERVATION_BTN_VAL)

	// click the 'book now' button for the specified booking type
	switch bookingType {
	case BOOKING_TYPE_STUDENT_ORGS:
		webdriver.FindAndClickElement(driver, BOOKING_TYPE_BTN_STUDENT_ORGS_BY, BOOKING_TYPE_BTN_STUDENT_ORGS_VAL)
	case BOOKING_TYPE_STUDY_ROOM:
		webdriver.FindAndClickElement(driver, BOOKING_TYPE_BTN_STUDY_ROOM_BY, BOOKING_TYPE_BTN_STUDY_ROOM_VAL)
	default:
		log.Fatalf("Error: invalid booking type: %d", bookingType)
	}
}

// set reservation time: booking date, start time, end time, click search
func setReservationTime(driver selenium.WebDriver, resTime myconfig.ReservationTime) {
	// wait for content to load
	webdriver.WaitForElementReady(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL)

	// input the booking date
	webdriver.PressKey(driver, selenium.BackspaceKey, BOOKING_DATE_BACKSPACE_STROKES)
	webdriver.ClearAndSendKeys(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL, resTime.Date)
	webdriver.PressKey(driver, selenium.TabKey, 1)

	// input the start time
	webdriver.ClearAndSendKeys(driver, START_TIME_INPUT_BY, START_TIME_INPUT_VAL, resTime.StartTime)

	// input the end time
	webdriver.ClearAndSendKeys(driver, END_TIME_INPUT_BY, END_TIME_INPUT_VAL, resTime.EndTime)

	// click the search button
	webdriver.FindAndClickElement(driver, SEARCH_BTN_BY, SEARCH_BTN_VAL)
}

// add reservation details to the booking
func addReservationDetails(driver selenium.WebDriver, reservationDetails myconfig.ReservationDetails) {
	// scroll to top of page
	webdriver.ScrollToTop(driver)

	// click the reservation details button
	webdriver.FindAndClickElement(driver, RESERVATION_DETAILS_BTN_BY, RESERVATION_DETAILS_BTN_VAL)

	// input the event name
	webdriver.ClearAndSendKeys(driver, EVENT_NAME_INPUT_BY, EVENT_NAME_INPUT_VAL, reservationDetails.EventName)

	// event type - do nothing if event type is study room
	if reservationDetails.EventType != EVENT_TYPE_STUDY_ROOM {
		log.Fatalf("Error: invalid event type: '%s' - only '%s' is supported at this time", reservationDetails.EventType, EVENT_TYPE_STUDY_ROOM)
	}

	// select organization - only ORGANIZATION_GROUP_STUDY_OPT_VAL is supported at this time
	webdriver.SelectFromDropdown(driver, ORGANIZATION_INPUT_BY, ORGANIZATION_INPUT_VAL, ORGANIZATION_GROUP_STUDY_OPT_VAL)

	// input the contact name
	webdriver.WaitForElementReady(driver, CONTACT_NAME_INPUT_BY, CONTACT_NAME_INPUT_VAL)
	webdriver.ClearAndSendKeys(driver, CONTACT_NAME_INPUT_BY, CONTACT_NAME_INPUT_VAL, reservationDetails.ContactName)

	// input the contact phone
	webdriver.ClearAndSendKeys(driver, CONTACT_PHONE_INPUT_BY, CONTACT_PHONE_INPUT_VAL, reservationDetails.ContactPhone)

	// input the contact email
	webdriver.ClearAndSendKeys(driver, CONTACT_EMAIL_INPUT_BY, CONTACT_EMAIL_INPUT_VAL, reservationDetails.ContactEmail)

	// select reserver status
	webdriver.SelectFromDropdown(driver, RESERVER_STATUS_INPUT_BY, RESERVER_STATUS_INPUT_VAL, RESERVER_STATUS_STUDENT_OPT_VAL)

	// input the description
	webdriver.WaitForElementReady(driver, DESCRIPTION_INPUT_BY, DESCRIPTION_INPUT_VAL)
	webdriver.ClearAndSendKeys(driver, DESCRIPTION_INPUT_BY, DESCRIPTION_INPUT_VAL, reservationDetails.Description)
}

// Create reservation button, data-bind="click: function(){ return saveReservation(); }" - click
// finish reservation: click create reservation button
func finishReservation(driver selenium.WebDriver) {
	// click the create reservation button
	finishBtn := webdriver.FindElement(driver, FINISH_RESERVATION_BTN_BY, FINISH_RESERVATION_BTN_VAL)
	webdriver.ScrollElemIntoView(driver, finishBtn, "finish reservation button")
	webdriver.ClickElement(finishBtn, "finish reservation button")
}
