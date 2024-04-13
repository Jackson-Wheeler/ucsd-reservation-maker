// page-methods.go: page methods for the PlaywrightWrapper struct
package playwrightwrapper

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// Methods by which to find elements.
const (
	ByText    = "ByText"
	ByTitle   = "ByTitle"
	ByLocator = "ByLocator"
)

/*
FindElement finds an element by the specified method and value. Returns a Locator object, with no garuntee that that object is valid.

method: the method by which to find the element (e.g. ByText, ByLocator)

value: the value to search for
*/
func (pw *PlaywrightWrapper) FindElement(method string, value string) (playwright.Locator, error) {
	if method == ByText {
		return pw.CurrPage.GetByText(value).First(), nil
	} else if method == ByTitle {
		return pw.CurrPage.GetByTitle(value).First(), nil
	} else if method == ByLocator {
		return pw.CurrPage.Locator(value), nil
	} else {
		return nil, fmt.Errorf("FindElement() invalid method: '%s'", method)
	}
}

/*
FindElementFromElement finds an element by the specified method and value, relative to the specified element. Returns a Locator object, with no garuntee that that object is valid.
*/
func (pw *PlaywrightWrapper) FindElementFromElement(elem playwright.Locator, method string, value string) (playwright.Locator, error) {
	if method == ByText {
		return elem.GetByText(value).First(), nil
	} else if method == ByTitle {
		return elem.GetByTitle(value).First(), nil
	} else if method == ByLocator {
		return elem.Locator(value), nil
	} else {
		return nil, fmt.Errorf("FindElement() invalid method: '%s'", method)
	}
}

/*
FindElemAndClick finds an element by the specified method and value, then clicks it.

method: the method by which to find the element (e.g. ByText, ByLocator)

value: the value to search for
*/
func (pw *PlaywrightWrapper) FindElemAndClick(method string, value string) error {
	elem, err := pw.FindElement(method, value)
	if err != nil {
		return fmt.Errorf("error finding element: %v", err)
	}

	err = elem.Click()
	if err != nil {
		return fmt.Errorf("error clicking element: %v", err)
	}

	return nil
}

/*
FindElemAndSendKeys finds an element by the specified method and value, then sends the specified keys to it.

method: the method by which to find the element (e.g. ByText, ByCSSSelector)

value: the value to search for
*/
func (pw *PlaywrightWrapper) FindElemAndSendKeys(method string, value string, keys string) error {
	elem, err := pw.FindElement(method, value)
	if err != nil {
		return fmt.Errorf("error finding element: %v", err)
	}

	trueVal := true
	err = elem.Clear(playwright.LocatorClearOptions{Force: &trueVal})
	if err != nil {
		return fmt.Errorf("error clearing element: %v", err)
	}

	err = elem.PressSequentially(keys)
	if err != nil {
		return fmt.Errorf("error sending keys: %v", err)
	}

	return nil
}

/*
PressKey presses the specified key.
*/
func (pw *PlaywrightWrapper) PressKey(keyName string) error {
	err := pw.CurrPage.Keyboard().Press(keyName)
	if err != nil {
		return fmt.Errorf("error pressing key '%s': %v", keyName, err)
	}

	return nil
}

/*
WaitForElement waits for an element to be ready.
*/
func (pw *PlaywrightWrapper) WaitForElement(method string, value string) error {
	elem, err := pw.FindElement(method, value)
	if err != nil {
		return fmt.Errorf("error finding element: %v", err)
	}

	err = elem.WaitFor()
	if err != nil {
		return fmt.Errorf("error waiting for element: %v", err)
	}

	return nil
}

/*
SelectFromDropdown selects the specified value from the dropdown, which can be found by the value with the given method.
*/
func (pw *PlaywrightWrapper) SelectFromDropdown(method string, value string, selectionValue string) error {
	// Find dropdown
	dropdownElem, err := pw.FindElement(method, value)
	if err != nil {
		return fmt.Errorf("error finding dropdown: %v", err)
	}

	// Select option
	_, err = dropdownElem.SelectOption(playwright.SelectOptionValues{Values: &[]string{selectionValue}})
	if err != nil {
		return fmt.Errorf("error selecting option: %v", err)
	}

	return nil
}
