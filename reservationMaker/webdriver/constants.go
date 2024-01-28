// Package webdriver is a package for working with Selenium webdriver & service
//
// It provides InitiliazeWebDriver() for initializing a Selenium webdriver
// service and driver, and api functions for easier interaction with the
// Selenium webdriver.
//
// It is not an actual webdriver itself, but rather a wrapper for working with
// the Selenium webdriver.
package webdriver

const (
	/* -- WEB DRIVER INIT -- */
	DRIVER_DIR             = "drivers"
	DRIVER_TYPE_DIR_PREFIX = "chromedriver-"
	DRIVER_NAME            = "chromedriver"
)
