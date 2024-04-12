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

func (pw *PlaywrightWrapper) PressKey(keyName string) error {
	err := pw.CurrPage.Keyboard().Press(keyName)
	if err != nil {
		return fmt.Errorf("error pressing key '%s': %v", keyName, err)
	}

	return nil
}
