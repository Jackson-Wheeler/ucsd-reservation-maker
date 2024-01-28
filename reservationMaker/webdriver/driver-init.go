// Package webdriver is a driver for working with Selenium webdriver & service
//
// It provides InitiliazeWebDriver() for initializing a Selenium webdriver
// service and driver, and api functions for easier interaction with the
// Selenium webdriver.
package webdriver

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/tebeka/selenium"
)

// InitializeWebDriver initializes a Selenium webdriver service & driver. NOTE: the service must be stopped after use by calling service.Stop()
func InitializeWebDriver(maximizeWindow bool) (*selenium.Service, selenium.WebDriver) {
	fmt.Println("Initializing Selenium webdriver...")

	// initialize a Chrome browser instance on port 4444
	driverPath, err := getDriverPath()
	if err != nil {
		errMsg := "initializeWebDriver(): failed to get chromedriver path"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}

	service, err := selenium.NewChromeDriverService(driverPath, 4444)
	if err != nil {
		errMsg := "initializeWebDriver(): failed to create new ChromeDriverService"
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

// getDriverPath returns the path to the correct chromedriver executable based on the host OS & Architecutre
func getDriverPath() (string, error) {
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
