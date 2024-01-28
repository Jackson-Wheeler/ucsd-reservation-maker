package webdriver

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

func myFindElement(driver selenium.WebDriver, by string, value string) selenium.WebElement {
	elem, err := driver.FindElement(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("failed to find element by %s with value '%s'", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elem
}

// same as myFindElement() but takes an element as an argument
func myFindElementFromElement(element selenium.WebElement, by string, value string) selenium.WebElement {
	elem, err := element.FindElement(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("failed to find element by %s with value '%s', looking from element", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elem
}

func myFindElements(driver selenium.WebDriver, by string, value string) []selenium.WebElement {
	elems, err := driver.FindElements(by, value)
	if err != nil {
		errMsg := fmt.Sprintf("failed to find element by %s with value '%s'", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return elems
}

func myIsDisplayed(element selenium.WebElement) bool {
	isDisplayed, err := element.IsDisplayed()
	if err != nil {
		errMsg := "failed to check if element is displayed"
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
	return isDisplayed
}

func myClickElement(driver selenium.WebDriver, by string, value string) {
	elem := myFindElement(driver, by, value)
	err := elem.Click()
	if err != nil {
		errMsg := fmt.Sprintf("failed to click element with {'%s'='%s'}", by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

func myNavToMostRecentTab(driver selenium.WebDriver) {
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

func mySendKeys(driver selenium.WebDriver, by string, value string, text string) {
	elem := myFindElement(driver, by, value)
	err := elem.SendKeys(text)
	if err != nil {
		errMsg := fmt.Sprintf("failed to input text '%s' into element with {'%s'='%s'}", text, by, value)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}

func myClearAndSendKeys(driver selenium.WebDriver, by string, value string, text string) {
	elem := myFindElement(driver, by, value)

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

func myScrollElemIntoView(driver selenium.WebDriver, element selenium.WebElement, elemName string) {
	_, err := driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{element})
	if err != nil {
		errMsg := fmt.Sprintf("failed to scroll %s element into view", elemName)
		log.Fatalf("Error: %s - %v", errMsg, err)
	}
}
