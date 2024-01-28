/* constants for webdriver package */
package webdriver

import "github.com/tebeka/selenium"

const (
	/* -- WEB DRIVER INIT -- */
	DRIVER_DIR             = "drivers"
	DRIVER_TYPE_DIR_PREFIX = "chromedriver-"
	DRIVER_NAME            = "chromedriver"

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
)
