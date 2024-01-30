// Package webdriver is a package for working with Selenium webdriver & service
//
// It provides InitiliazeWebDriver() for initializing a Selenium webdriver
// service and driver, and api functions for easier interaction with the
// Selenium webdriver.
//
// It is not an actual webdriver itself, but rather a wrapper for working with
// the Selenium webdriver.
package webdriver

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/tebeka/selenium"
)

const (
	SERVER_PORT = 4444
)

// InitializeWebDriver initializes a Selenium webdriver service & driver. NOTE: the service must be stopped after use by calling service.Stop()
func InitializeWebDriver(driverFolder string, driverName string, maximizeWindow bool) (*selenium.Service, selenium.WebDriver) {
	fmt.Println("Initializing Selenium webdriver...")

	// initialize a Chrome browser instance on port 4444
	driverPath := filepath.Join(".", driverFolder, driverName)

	service, err := selenium.NewChromeDriverService(driverPath, SERVER_PORT)
	if err != nil {
		errMsg := fmt.Sprintf("initializeWebDriver(): failed to create new ChromeDriverService at path '%s' on port '%d'", driverPath, SERVER_PORT)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}

	// configure the browser options
	caps := selenium.Capabilities{}

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		errMsg := "initializeWebDriver(): failed to create new selenium driver"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}

	if maximizeWindow {
		// maximize the current window to avoid responsive rendering
		err = driver.MaximizeWindow("")
		if err != nil {
			errMsg := "initializeWebDriver(): failed to maximize window"
			log.Fatalf("Error: %s - %v", errMsg, err)
		}
	}

	return service, driver
}
