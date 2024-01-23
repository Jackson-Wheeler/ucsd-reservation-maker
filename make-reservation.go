package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/tebeka/selenium"
)

const (
	SITE_URL = "https://reservations.ucsd.edu/EmsWebApp/Default.aspx"
)

type SiteCredentials struct {
	username string
	password string
}

func MakeReservation(config Config, siteCredentials SiteCredentials) {
	// initialize a Chrome browser instance on port 4444
	driverPath, err := getDriverPath()
	if err != nil {
		log.Fatal("Error:", err)
	}

	service, err := selenium.NewChromeDriverService(driverPath, 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	// err = driver.MaximizeWindow("")
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// visit the target page
	err = driver.Get(SITE_URL)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Login
	err = login(driver, siteCredentials)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// wait for 10 seconds before closing the browser
	time.Sleep(3 * time.Second)
}

// getDriverPath returns the path to the correct chromedriver executable based on the host OS & Architecutre
func getDriverPath() (string, error) {
	const DRIVER_DIR = "drivers"
	const DRIVER_TYPE_DIR_PREFIX = "chromedriver-"
	const DRIVER_NAME = "chromedriver"
	var driverType string

	switch runtime.GOOS {
	case "linux":
		driverType = "linux64"
	case "windows":
		driverType = "win64"
	case "darwin":
		if runtime.GOARCH == "amd64" {
			driverType = "mac-x64"
		} else if runtime.GOARCH == "arm64" {
			driverType = "mac-arm64"
		} else {
			return "", errors.New("unsupported architecture with darwin OS")
		}
	default:
		return "", errors.New("unsupported operating system: " + runtime.GOOS)
	}

	fmt.Printf("Selenium chromedriver type: %s\n", driverType)

	driverTypeDir := fmt.Sprintf("%s%s", DRIVER_TYPE_DIR_PREFIX, driverType)
	driverPath := filepath.Join(".", DRIVER_DIR, driverTypeDir, DRIVER_NAME)
	return driverPath, nil
}

func login(driver selenium.WebDriver, siteCredentials SiteCredentials) error {
	const LOGIN_BTN_TEXT = "Login"
	const USER_ID_INPUT_ID = "userID_input"
	const PASSWORD_INPUT_ID = "password_input"
	const SIGN_IN_BTN_ID = "pc_btnLogin"

	// find the login button by its text
	loginButton, err := driver.FindElement(selenium.ByLinkText, LOGIN_BTN_TEXT)
	if err != nil {
		return err
	}

	// click the login button
	err = loginButton.Click()
	if err != nil {
		return err
	}

	// Wait for the new tab to open
	time.Sleep(2 * time.Second)

	// Get the list of window handles
	windowHandles, err := driver.WindowHandles()
	if err != nil {
		return err
	}

	// Switch to the new tab
	err = driver.SwitchWindow(windowHandles[len(windowHandles)-1])
	if err != nil {
		return err
	}

	// find the user ID input field by its name
	userIDInput, err := driver.FindElement(selenium.ByID, USER_ID_INPUT_ID)
	if err != nil {
		return err
	}

	// input text into the user ID input field
	err = userIDInput.SendKeys(siteCredentials.username)
	if err != nil {
		return err
	}

	// find the password input field by its name
	passwordInput, err := driver.FindElement(selenium.ByID, PASSWORD_INPUT_ID)
	if err != nil {
		return err
	}

	// input text into the password input field
	err = passwordInput.SendKeys(siteCredentials.password)
	if err != nil {
		return err
	}

	// find the sign in button by its id
	signInBtn, err := driver.FindElement(selenium.ByID, SIGN_IN_BTN_ID)
	if err != nil {
		return err
	}

	// click the login button
	err = signInBtn.Click()
	if err != nil {
		return err
	}

	return nil
}
