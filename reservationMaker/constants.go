/* constants for webdriver package */
package reservationMaker

import "github.com/tebeka/selenium"

/* types */
type SiteCredentials struct {
	Username string
	Password string
}

/* constants */
const (
	/* -- WEB PAGE GENERAL -- */
	SITE_URL = "https://reservations.ucsd.edu/EmsWebApp/Default.aspx"

	/* -- BOOKING CONFIG -- */
	BOOKING_TYPE_STUDENT_ORGS = 0
	BOOKING_TYPE_STUDY_ROOM   = 1

	/* -- WEB PAGE LOGIN -- */
	LOGIN_BTN_BY  = selenium.ByLinkText
	LOGIN_BTN_VAL = "Login"

	USERNAME_INPUT_BY  = selenium.ByID
	USERNAME_INPUT_VAL = "userID_input"

	PASSWORD_INPUT_BY  = selenium.ByID
	PASSWORD_INPUT_VAL = "password_input"

	SIGN_IN_BTN_BY  = selenium.ByID
	SIGN_IN_BTN_VAL = "pc_btnLogin"

	/* -- WEB PAGE CREATE RESERVATION -- */
	// begin booking
	CREATE_RESERVATION_BTN_BY  = selenium.ByLinkText
	CREATE_RESERVATION_BTN_VAL = "Create A Reservation"

	BOOKING_TYPE_BTN_STUDENT_ORGS_BY  = selenium.ByXPATH
	BOOKING_TYPE_BTN_STUDENT_ORGS_VAL = `//*[@aria-label='Book Now With The "Request Spaces | Student Orgs" Template']`
	BOOKING_TYPE_BTN_STUDY_ROOM_BY    = selenium.ByXPATH
	BOOKING_TYPE_BTN_STUDY_ROOM_VAL   = `//*[@aria-label='Book Now With The "Reserve Spaces | Study Rooms & Open Desk" Template']`

	// reservation time
	BOOOKING_DATE_INPUT_BY  = selenium.ByID
	BOOOKING_DATE_INPUT_VAL = "booking-date-input"

	START_TIME_INPUT_BY  = selenium.ByID
	START_TIME_INPUT_VAL = "start-time-input"

	END_TIME_INPUT_BY  = selenium.ByID
	END_TIME_INPUT_VAL = "end-time-input"

	SEARCH_BTN_BY  = selenium.ByXPATH
	SEARCH_BTN_VAL = `//button[normalize-space()='Search']`

	// room selection
	ROOM_ITEM_BY  = selenium.ByCSSSelector
	ROOM_ITEM_VAL = ".room-column.column"

	ROOM_SELECT_BTN_BY  = selenium.ByCSSSelector
	ROOM_SELECT_BTN_VAL = ".fa-plus-circle"

	SELECTED_ROOM_ITEM_BY  = selenium.ByCSSSelector
	SELECTED_ROOM_ITEM_VAL = ".selected-room-item"

	NUMBER_OF_ATTENDEES_INPUT_BY  = selenium.ByID
	NUMBER_OF_ATTENDEES_INPUT_VAL = "setup-add-count"

	ADD_ROOM_BTN_BY  = selenium.ByID
	ADD_ROOM_BTN_VAL = "setup--add-modal-save"

	ALERT_BY  = selenium.ByCSSSelector
	ALERT_VAL = ".alert"

	ALERT_MESSAGE_BY  = selenium.ByCSSSelector
	ALERT_MESSAGE_VAL = ".message"
)
