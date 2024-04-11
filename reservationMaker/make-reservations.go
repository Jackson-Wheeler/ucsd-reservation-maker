package reservationMaker

import (
	"fmt"
	"log"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/myplaywright"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/webdriver"
	"github.com/tebeka/selenium"
)

// MakeReservations makes the reservations on reservations.ucsd.edu according to the given config, using the given site credentials to login.
//
// openFlag: if true, instead of making reservations, program will only open the reservations page (navigates to the date and time of the first reservation in the config file). Purpose: for viewing what dates and times are available.
func MakeReservations(config myconfig.Config, siteCredentials SiteCredentials, openFlag bool) error {
	// start Playwright
	pw := &myplaywright.MyPlaywright{}
	err := pw.Initialize(false)
	if err != nil {
		return fmt.Errorf("error initializing Playwright (automated browser controlling software): %v", err)
	}
	defer pw.Close()

	// // visit the target page
	// visitTargetPage(driver)

	// // login
	// login(driver, siteCredentials)

	// // create each reservation
	// for _, time := range config.ReservationTimes {
	// 	createReservation(driver, time, config.RoomPreferenceOrder, config.ReservationDetails)
	// }

	// // finish up
	// fmt.Println("\nDone - see above for log of created reservations")

	return nil
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
	webdriver.FindAndClickElement(driver, LOGIN_BTN_BY, LOGIN_BTN_VAL)

	// switch to new tab
	webdriver.NavToMostRecentTab(driver)

	// enter username
	webdriver.SendKeys(driver, USERNAME_INPUT_BY, USERNAME_INPUT_VAL, siteCredentials.Username)

	// enter password
	webdriver.SendKeys(driver, PASSWORD_INPUT_BY, PASSWORD_INPUT_VAL, siteCredentials.Password)

	// click sign in button
	webdriver.FindAndClickElement(driver, SIGN_IN_BTN_BY, SIGN_IN_BTN_VAL)
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
		fmt.Printf("*no reservation made for %s from %s to %s - %v\n", resTime.Date, resTime.StartTime, resTime.EndTime, err)
		return
	}
	fmt.Printf("selected room '%s'\n", roomName)

	// add reservation details
	addReservationDetails(driver, reservationDetails)

	// click create reservation button
	finishReservation(driver)

	fmt.Printf("*reservation created for '%s'\n", roomName)
}
