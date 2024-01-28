package webdriver

import (
	"fmt"
	"log"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

type SiteCredentials struct {
	Username string
	Password string
}

func MakeReservations(config myconfig.Config, siteCredentials SiteCredentials) {
	// initialize the Selenium service & webdriver
	service, driver := initializeWebDriver(false)
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
	myClickElement(driver, LOGIN_BTN_BY, LOGIN_BTN_VAL)

	// switch to new tab
	myNavToMostRecentTab(driver)

	// enter username
	mySendKeys(driver, USERNAME_INPUT_BY, USERNAME_INPUT_VAL, siteCredentials.Username)

	// enter password
	mySendKeys(driver, PASSWORD_INPUT_BY, PASSWORD_INPUT_VAL, siteCredentials.Password)

	// click sign in button
	myClickElement(driver, SIGN_IN_BTN_BY, SIGN_IN_BTN_VAL)
}
