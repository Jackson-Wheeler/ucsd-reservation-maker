package webdriver

import (
	"fmt"
	"log"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

const (
	SITE_URL = "https://reservations.ucsd.edu/EmsWebApp/Default.aspx"
)

type SiteCredentials struct {
	Username string
	Password string
}

func MakeReservation(config myconfig.Config, siteCredentials SiteCredentials) {
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

	// wait for delay before closing the browser
	time.Sleep(3 * time.Second)
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
	const LOGIN_BTN_TEXT = "Login"
	const USER_ID_INPUT_ID = "userID_input"
	const PASSWORD_INPUT_ID = "password_input"
	const SIGN_IN_BTN_ID = "pc_btnLogin"

	fmt.Printf("Logging in as user: '%s'...\n", siteCredentials.Username)

	// click the login button
	myClickElement(driver, selenium.ByLinkText, LOGIN_BTN_TEXT)

	// switch to new tab
	myNavToMostRecentTab(driver)

	// enter username
	myInputText(driver, selenium.ByID, USER_ID_INPUT_ID, siteCredentials.Username)

	// find the password input field by its name
	myInputText(driver, selenium.ByID, PASSWORD_INPUT_ID, siteCredentials.Password)

	// find the sign in button by its id
	myClickElement(driver, selenium.ByID, SIGN_IN_BTN_ID)
}
