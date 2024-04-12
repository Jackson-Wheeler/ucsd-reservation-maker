/* constants for webdriver package */
package reservationMaker

import (
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/playwrightwrapper"
	"github.com/tebeka/selenium"
)

/* types */
type SiteCredentials struct {
	Username string
	Password string
}

/* constants */
const (
	/* -- WEB DRIVER INIT -- */
	DRIVER_DIR             = "driver"
	DRIVER_NAME            = "chromedriver"
	MAXIMIZE_DRIVER_WINDOW = true

	/* -- WEB PAGE GENERAL -- */
	SITE_URL = "https://reservations.ucsd.edu/EmsWebApp/Default.aspx"

	/* -- BOOKING CONFIG -- */
	BOOKING_TYPE_STUDENT_ORGS = 0
	BOOKING_TYPE_STUDY_ROOM   = 1

	EVENT_TYPE_STUDY_ROOM = "STUDY_ROOM"
	EVENT_TYPE_NAP_NOOK   = "NAP_NOOK"

	ORGANIZATION_CLUB_OPT_VAL        = "910"
	ORGANIZATION_GROUP_STUDY_OPT_VAL = "5708"

	RESERVER_STATUS_STUDENT_OPT_VAL = "53"

	/* -- WEB PAGE LOGIN -- */
	LOGIN_BTN_BY  = playwrightwrapper.ByText
	LOGIN_BTN_VAL = "Login"

	USERNAME_INPUT_BY  = playwrightwrapper.ByLocator
	USERNAME_INPUT_VAL = "#userID_input"

	PASSWORD_INPUT_BY  = playwrightwrapper.ByLocator
	PASSWORD_INPUT_VAL = "#password_input"

	SIGN_IN_BTN_BY  = playwrightwrapper.ByLocator
	SIGN_IN_BTN_VAL = "#pc_btnLogin"

	/* -- WEB PAGE CREATE RESERVATION -- */
	// begin booking
	CREATE_RESERVATION_BTN_BY  = playwrightwrapper.ByTitle
	CREATE_RESERVATION_BTN_VAL = "Create A Reservation"

	BOOKING_TYPE_BTN_STUDENT_ORGS_BY  = playwrightwrapper.ByLocator
	BOOKING_TYPE_BTN_STUDENT_ORGS_VAL = `//*[@aria-label='Book Now With The "Request Spaces | Student Orgs" Template']`
	BOOKING_TYPE_BTN_STUDY_ROOM_BY    = playwrightwrapper.ByLocator
	BOOKING_TYPE_BTN_STUDY_ROOM_VAL   = `//*[@aria-label='Book Now With The "Reserve Spaces | Study Rooms & Open Desk" Template']`

	// reservation time
	BOOKING_DATE_INPUT_BY  = playwrightwrapper.ByLocator
	BOOKING_DATE_INPUT_VAL = "#booking-date-input"
	//BOOKING_DATE_BACKSPACE_STROKES = 15

	START_TIME_INPUT_BY  = playwrightwrapper.ByLocator
	START_TIME_INPUT_VAL = "#start-time-input"

	END_TIME_INPUT_BY  = playwrightwrapper.ByLocator
	END_TIME_INPUT_VAL = "#end-time-input"

	SEARCH_BTN_BY  = playwrightwrapper.ByLocator
	SEARCH_BTN_VAL = "#date-time-collapse .btn-filter-search"

	// room selection
	ROOM_ITEM_BY  = selenium.ByCSSSelector
	ROOM_ITEM_VAL = ".room-column.column"

	ROOM_SELECT_BTN_BY  = selenium.ByCSSSelector
	ROOM_SELECT_BTN_VAL = ".fa-plus-circle"

	NUMBER_OF_ATTENDEES_INPUT_BY  = selenium.ByID
	NUMBER_OF_ATTENDEES_INPUT_VAL = "setup-add-count"

	ADD_ROOM_BTN_BY  = selenium.ByID
	ADD_ROOM_BTN_VAL = "setup--add-modal-save"

	ALERT_BY  = selenium.ByCSSSelector
	ALERT_VAL = ".alert.alert-danger"

	ALERT_MESSAGE_BY  = selenium.ByCSSSelector
	ALERT_MESSAGE_VAL = ".message"

	SELECTED_ROOM_ITEM_BY  = selenium.ByCSSSelector
	SELECTED_ROOM_ITEM_VAL = ".selected-room-name"

	// reservation details
	RESERVATION_DETAILS_BTN_BY  = selenium.ByXPATH
	RESERVATION_DETAILS_BTN_VAL = `//a[@aria-label='Create a Reservation/Reservation Details']`

	EVENT_NAME_INPUT_BY  = selenium.ByID
	EVENT_NAME_INPUT_VAL = "event-name"

	EVENT_TYPE_INPUT_BY  = selenium.ByID
	EVENT_TYPE_INPUT_VAL = "event-type"

	ORGANIZATION_INPUT_BY  = selenium.ByID
	ORGANIZATION_INPUT_VAL = "availablegroups"

	CONTACT_NAME_INPUT_BY  = selenium.ByXPATH
	CONTACT_NAME_INPUT_VAL = `//input[@id='1stContactName']`

	CONTACT_PHONE_INPUT_BY  = selenium.ByXPATH
	CONTACT_PHONE_INPUT_VAL = `//input[@id='1stContactPhone1']`

	CONTACT_EMAIL_INPUT_BY  = selenium.ByXPATH
	CONTACT_EMAIL_INPUT_VAL = `//input[@id='1stContactEmail']`

	RESERVER_STATUS_INPUT_BY  = selenium.ByXPATH
	RESERVER_STATUS_INPUT_VAL = `//select[@id='28']`

	DESCRIPTION_INPUT_BY  = selenium.ByXPATH
	DESCRIPTION_INPUT_VAL = `//textarea[@id='27']`

	// finish reservation
	FINISH_RESERVATION_BTN_BY  = selenium.ByXPATH
	FINISH_RESERVATION_BTN_VAL = `//*[@data-bind='click: function(){ return saveReservation(); }']`

	OK_CONFIRMATION_BTN_BY  = selenium.ByXPATH
	OK_CONFIRMATION_BTN_VAL = `//button[@id='help-text-close-btn']`
)
