package reservationMaker

import (
	"fmt"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/playwrightwrapper"
)

// MakeReservations makes the reservations on reservations.ucsd.edu according to the given config, using the given site credentials to login.
//
// openFlag: if true, instead of making reservations, program will only open the reservations page (navigates to the date and time of the first reservation in the config file). Purpose: for viewing what dates and times are available.
func MakeReservations(config myconfig.Config, siteCredentials SiteCredentials, openFlag bool) error {
	// initialize PlaywrightWrapper - Playwright is the software controlling the automated browser
	pw := &playwrightwrapper.PlaywrightWrapper{}
	err := pw.Initialize(false, false)
	if err != nil {
		return fmt.Errorf("error initializing Playwright (automated browser controlling software): %v", err)
	}
	defer pw.Close()

	// visit the target page
	err = visitTargetPage(pw, SITE_URL)
	if err != nil {
		return fmt.Errorf("error visiting target page: %v", err)
	}

	// login
	err = login(pw, siteCredentials)
	if err != nil {
		return fmt.Errorf("error logging in: %v", err)
	}

	// create each reservation
	for _, time := range config.ReservationTimes {
		err = createReservation(pw, time, config.RoomPreferenceOrder, config.ReservationDetails)
		if err != nil {
			return fmt.Errorf("error creating reservation for date '%s': %v", time.Date, err)
		}
	}

	// finish up
	fmt.Println("\nDone - see above for log of created reservations")

	return nil
}

func visitTargetPage(pw *playwrightwrapper.PlaywrightWrapper, url string) error {
	fmt.Printf("Navigating to target page: '%s'...\n", url)

	_, err := pw.CurrPage.Goto(SITE_URL)
	if err != nil {
		return fmt.Errorf("failed to navigate to target page: %v", err)
	}

	return nil
}

func login(pw *playwrightwrapper.PlaywrightWrapper, siteCredentials SiteCredentials) error {
	fmt.Printf("Logging in as user: '%s'...\n", siteCredentials.Username)

	// click the login button
	err := pw.FindElemAndClick(LOGIN_BTN_BY, LOGIN_BTN_VAL)
	if err != nil {
		return fmt.Errorf("failed to click login button: %v", err)
	}

	// wait (for new tab to open)
	time.Sleep(500 * time.Millisecond)

	// switch to new tab
	err = pw.SwitchPage(1) // the new tab should be the 2nd tab opened
	if err != nil {
		return fmt.Errorf("failed to switch to new tab: %v", err)
	}

	// Enter username
	err = pw.FindElemAndSendKeys(USERNAME_INPUT_BY, USERNAME_INPUT_VAL, siteCredentials.Username)
	if err != nil {
		return fmt.Errorf("failed to enter username: %v", err)
	}

	// enter password
	err = pw.FindElemAndSendKeys(PASSWORD_INPUT_BY, PASSWORD_INPUT_VAL, siteCredentials.Password)
	if err != nil {
		return fmt.Errorf("failed to enter password: %v", err)
	}

	// click sign in
	err = pw.FindElemAndClick(SIGN_IN_BTN_BY, SIGN_IN_BTN_VAL)
	if err != nil {
		return fmt.Errorf("failed to click sign in button: %v", err)
	}

	return nil
}

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(pw *playwrightwrapper.PlaywrightWrapper, resTime myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) error {

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", resTime.Date, resTime.StartTime, resTime.EndTime)

	// begin booking
	err := beginBooking(pw, BOOKING_TYPE_STUDY_ROOM)
	if err != nil {
		return fmt.Errorf("error beginning booking: %v", err)
	}

	// set reservation time
	err = setReservationTime(pw, resTime)
	if err != nil {
		return fmt.Errorf("error setting reservation time: %v", err)
	}

	// select room
	roomName, err := selectRoom(pw, roomPreferenceOrder, reservationDetails)
	if err != nil {
		fmt.Printf("*no reservation made for %s from %s to %s - %v\n", resTime.Date, resTime.StartTime, resTime.EndTime, err)
		return nil
	}
	fmt.Printf("selected room '%s'\n", roomName)

	// add reservation details
	err = addReservationDetails(pw, reservationDetails)
	if err != nil {
		return fmt.Errorf("error adding reservation details: %v", err)
	}

	// click create reservation button
	err = finishReservation(pw)
	if err != nil {
		return fmt.Errorf("error finishing reservation: %v", err)
	}

	fmt.Printf("*reservation created for '%s'\n", roomName)

	return nil
}
