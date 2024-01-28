package reservationMaker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/webdriver"
	"github.com/tebeka/selenium"
)

// select room based on room preference order, returning name of room selected, or error otherwise
func selectRoom(driver selenium.WebDriver, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) (string, error) {
	// wait for content to load
	webdriver.WaitForElementReady(driver, ROOM_ITEM_BY, ROOM_ITEM_VAL)

	// map room names to web page room elements
	roomMap := getRoomMap(driver)

	// attempt selecting rooms in order of preference
	for _, roomName := range roomPreferenceOrder {
		err := attemptSelectRoom(driver, roomMap[roomName], reservationDetails)
		if err != nil {
			fmt.Printf("room '%s' selection failed: %v\n", roomName, err)
			continue
		}
		// successfully selected room
		return roomName, nil
	}

	// if preferred rooms are not available
	return "", errors.New("preferred rooms are not available")

}

// get a map of room names to room elements (from web page)
func getRoomMap(driver selenium.WebDriver) map[string]selenium.WebElement {
	// get all rooms
	roomItems := webdriver.FindElements(driver, ROOM_ITEM_BY, ROOM_ITEM_VAL)

	// make map of room name to room element (from web page)
	roomMap := make(map[string]selenium.WebElement)
	for _, roomItem := range roomItems {
		title, err1 := roomItem.GetAttribute("title")
		if err1 != nil {
			log.Fatal("Error: failed to get room item title - ", err1)
		}
		roomMap[title] = roomItem
	}

	return roomMap
}

// attempts to click to select the given room, and fills out the initial pop up, returns error if unsuccessful
func attemptSelectRoom(driver selenium.WebDriver, roomItem selenium.WebElement, reservationDetails myconfig.ReservationDetails) error {
	// find child button for selecting room
	selectBtn := webdriver.FindElementFromElement(roomItem, ROOM_SELECT_BTN_BY, ROOM_SELECT_BTN_VAL)

	// check if btn is displayed
	isDisplayed := webdriver.IsDisplayed(selectBtn)
	if !isDisplayed {
		return fmt.Errorf("room is not available")
	}

	// scroll the select button into view
	webdriver.ScrollElemIntoView(driver, selectBtn, "room select button")

	// click the select button
	err := selectBtn.Click()
	if err != nil {
		return fmt.Errorf("failed to click select button - %v", err)
	}

	fillInitialPopUp(driver, reservationDetails)

	// check if room was successfully selected (shows up in selected rooms)
	err = confirmSelection(driver)
	if err != nil {
		return err
	}
	// selectedRooms := myFindElements(driver, selenium.ByCSSSelector, SELECTED_ROOM_CSS_SELECTOR)
	// if len(selectedRooms) == 0 {
	// 	return fmt.Errorf("room '%s' was not successfully selected", roomName)
	// }

	return nil
}

func fillInitialPopUp(driver selenium.WebDriver, reservationDetails myconfig.ReservationDetails) {
	// wait for content to load
	webdriver.WaitForElementReady(driver, NUMBER_OF_ATTENDEES_INPUT_BY, NUMBER_OF_ATTENDEES_INPUT_VAL)

	// number of attendees
	for i := 0; i < reservationDetails.NumPeople; i++ {
		webdriver.SendKeys(driver, NUMBER_OF_ATTENDEES_INPUT_BY, NUMBER_OF_ATTENDEES_INPUT_VAL, selenium.UpArrowKey)
	}

	// TODO expand to be able to select setup type

	// click add room button
	webdriver.ClickElement(driver, ADD_ROOM_BTN_BY, ADD_ROOM_BTN_VAL)
}

func confirmSelection(driver selenium.WebDriver) error {
	// Check for alert message
	time.Sleep(500 * time.Millisecond)
	alertElem := webdriver.FindElementIfExists(driver, ALERT_BY, ALERT_VAL)
	if alertElem != nil {
		alertMessage := getAlertMessage(alertElem)
		return fmt.Errorf("room selection failed with alert message: %s", alertMessage)
	}

	return nil
}

func getAlertMessage(alertElem selenium.WebElement) string {
	alertMessageElem := webdriver.FindElementFromElement(alertElem, ALERT_MESSAGE_BY, ALERT_MESSAGE_VAL)
	alertMessage, err := alertMessageElem.Text()
	if err != nil {
		alertMessage = "[failed to get alert message]"
	}
	return alertMessage
}
