package reservationMaker

import (
	"fmt"
	"log"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/webdriver"
	"github.com/tebeka/selenium"
)

// MakeReservations makes the reservations on reservations.ucsd.edu according to the given config, using the given site credentials to login
func MakeReservations(config myconfig.Config, siteCredentials SiteCredentials) {
	// initialize the Selenium service & driver
	service, driver := webdriver.InitializeWebDriver(false)
	defer service.Stop()
	fmt.Println()

	// visit the target page
	visitTargetPage(driver)

	// Login
	login(driver, siteCredentials)

	// Create Reservations
	for _, time := range config.ReservationTimes {
		createReservation(driver, time, config.RoomPreferenceOrder, config.ReservationDetails)
	}

	// TEMP: keep browser by exiting on error
	log.Fatal("Done")
}

func visitTargetPage(driver selenium.WebDriver) {
	fmt.Printf("Navigating to target page: '%s'...\n", SITE_URL)

	err := driver.Get(SITE_URL)
	if err != nil {
		errMsg := fmt.Sprintf("MakeReservation(): failed to navigate to target page: '%s'", SITE_URL)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

func login(driver selenium.WebDriver, siteCredentials SiteCredentials) {
	fmt.Printf("Logging in as user: '%s'...\n", siteCredentials.Username)

	// click the login button
	webdriver.ClickElement(driver, LOGIN_BTN_BY, LOGIN_BTN_VAL)

	// switch to new tab
	webdriver.NavToMostRecentTab(driver)

	// enter username
	webdriver.SendKeys(driver, USERNAME_INPUT_BY, USERNAME_INPUT_VAL, siteCredentials.Username)

	// enter password
	webdriver.SendKeys(driver, PASSWORD_INPUT_BY, PASSWORD_INPUT_VAL, siteCredentials.Password)

	// click sign in button
	webdriver.ClickElement(driver, SIGN_IN_BTN_BY, SIGN_IN_BTN_VAL)
}

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, resTime myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", resTime.Date, resTime.StartTime, resTime.EndTime)

	// begin booking
	beginBooking(driver, BOOKING_TYPE_STUDY_ROOM)

	// set reservation time
	setReservationTime(driver, resTime)

	// select room
	roomName, err := selectRoom(driver, roomPreferenceOrder, reservationDetails)
	if err != nil {
		fmt.Printf("no reservation made for %s from %s to %s - %v\n", resTime.Date, resTime.StartTime, resTime.EndTime, err)
		return
	}
	fmt.Printf("selected room '%s'\n", roomName)

	// Reservation Details, button aria-label="Create a Reservation/Reservation Details"
	// event name id="event-name" - clear & send keys
	// event type id="event-type" - just do nothing if event type is study room
	// organization id="availablegroups" - click, down arrow twice, enter
	// id="1stContactName" - clear & send keys
	// id="1stContactPhone1" - clear & send keys
	// id="1stContactEmail"	- clear & send keys
	// choose who you are id="28" - click, down arrow three times, enter
	// description id="27" - clear & send keys

	// Create reservation button, data-bind="click: function(){ return saveReservation(); }" - click
	// confirm reservation was made
}
