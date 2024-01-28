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

	"github.com/tebeka/selenium"
)

// FindElement finds an element by the specified method and value, exits on error
func FindElement(driver selenium.WebDriver, by string, value string) selenium.WebElement {
	elem, err := driver.FindElement(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("failed to find element by %s with value '%s'", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elem
}

// FindElementIfExists finds an element by the specified method and value, returns nil if not found
func FindElementIfExists(driver selenium.WebDriver, by string, value string) selenium.WebElement {
	elemList := FindElements(driver, by, value)
	if len(elemList) == 0 {
		return nil
	}
	return elemList[0]
}

// FindElementFromElement same as myFindElement() but takes an element as an argument
func FindElementFromElement(element selenium.WebElement, by string, value string) selenium.WebElement {
	elem, err := element.FindElement(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("failed to find element by %s with value '%s', looking from element", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elem
}

// FindElements finds all elements by the specified method and value, exits on error
func FindElements(driver selenium.WebDriver, by string, value string) []selenium.WebElement {
	elems, err := driver.FindElements(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("error when finding elements by %s with value '%s'", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elems
}

// IsDisplayed returns true if given element is displayed, false if not, exits on error
func IsDisplayed(element selenium.WebElement) bool {
	isDisplayed, err := element.IsDisplayed()
	if err != nil {
		errMsg := "failed to check if element is displayed"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return isDisplayed
}

// ClickElement finds & clicks an element by the specified method and value, exits on error
func ClickElement(driver selenium.WebDriver, by string, value string) {
	elem := FindElement(driver, by, value)
	err := elem.Click()
	if err != nil {
		errMsg := fmt.Sprintf("failed to click element with {'%s'='%s'}", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

// NavToMostRecentTab switches to the most recent tab, exits on error
func NavToMostRecentTab(driver selenium.WebDriver) {
	// Get the list of window handles
	windowHandles, err := driver.WindowHandles()
	if err != nil {
		errMsg := "myNavToMostRecentTab(): failed to get window handles"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}

	// Switch to the new tab
	err = driver.SwitchWindow(windowHandles[len(windowHandles)-1])
	if err != nil {
		errMsg := "myNavToMostRecentTab(): failed to switch to most recent tab"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

// SendKeys finds & sends keys to an element by the specified method and value, exits on error
func SendKeys(driver selenium.WebDriver, by string, value string, text string) {
	elem := FindElement(driver, by, value)
	err := elem.SendKeys(text)
	if err != nil {
		errMsg := fmt.Sprintf("failed to input text '%s' into element with {'%s'='%s'}", text, by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

// ClearAndSendKeysfinds & clears & sends keys to an element by the specified method and value, exits on error
func ClearAndSendKeys(driver selenium.WebDriver, by string, value string, text string) {
	elem := FindElement(driver, by, value)

	err := elem.Clear()
	if err != nil {
		errMsg := fmt.Sprintf("failed to clear element with {'%s'='%s'}", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}

	err = elem.SendKeys(text)
	if err != nil {
		errMsg := fmt.Sprintf("failed to input text '%s' into element with {'%s'='%s'}", text, by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

// ScrollElemIntoView scrolls an element into view, exits on error
func ScrollElemIntoView(driver selenium.WebDriver, element selenium.WebElement, elemName string) {
	_, err := driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{element})
	if err != nil {
		errMsg := fmt.Sprintf("failed to scroll %s element into view", elemName)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}
