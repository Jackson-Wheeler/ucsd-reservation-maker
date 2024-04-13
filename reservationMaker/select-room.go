package reservationMaker

import (
	"fmt"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/playwrightwrapper"
	"github.com/playwright-community/playwright-go"
)

// select room based on room preference order, returning name of room selected, or error otherwise
func selectRoom(pw *playwrightwrapper.PlaywrightWrapper, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) (string, error) {
	// wait for content to load
	err := pw.WaitForElement(ROOM_ITEMS_READY_BY, ROOM_ITEMS_READY_VAL)
	if err != nil {
		return "", fmt.Errorf("error waiting for room items: %v", err)
	}

	// map room names to web page room elements
	roomMap, err := getRoomMap(pw)
	if err != nil {
		return "", fmt.Errorf("error getting map of rooms: %v", err)
	}

	// attempt selecting rooms in order of preference
	for _, roomName := range roomPreferenceOrder {
		selectedRoomName, err := attemptSelectRoom(pw, roomMap[roomName], reservationDetails)
		if err != nil {
			fmt.Printf("room '%s' selection failed: %v\n", roomName, err)
			continue
		}
		if selectedRoomName != roomName {
			fmt.Printf("selection inconsistency: instead of '%s', '%s' was selected\n", roomName, selectedRoomName)
		}
		// successfully selected room
		return selectedRoomName, nil
	}

	// // if preferred rooms are not available
	// return "", errors.New("preferred rooms are not available")
	return "", nil
}

// get a map of room names to room elements (from web page)
func getRoomMap(pw *playwrightwrapper.PlaywrightWrapper) (map[string]playwright.Locator, error) {
	// get all rooms
	roomItemLocator, _ := pw.FindElement(ROOM_ITEM_BY, ROOM_ITEM_VAL)
	roomItems, _ := roomItemLocator.All()

	// make map of room name to room element (from web page)
	roomMap := make(map[string]playwright.Locator)
	for _, roomItem := range roomItems {
		title, err1 := roomItem.GetAttribute("title")
		if err1 != nil {
			return nil, fmt.Errorf("failed to get room item title - %v", err1)
		}
		roomMap[title] = roomItem
	}

	return roomMap, nil
}

// attempts to click to select the given room, and fills out the initial pop up, returns error if unsuccessful
func attemptSelectRoom(pw *playwrightwrapper.PlaywrightWrapper, roomItem playwright.Locator, reservationDetails myconfig.ReservationDetails) (string, error) {
	// find child button for selecting room
	selectBtn, err := pw.FindElementFromElement(roomItem, ROOM_SELECT_BTN_BY, ROOM_SELECT_BTN_VAL)
	if err != nil {
		return "", fmt.Errorf("failed to find select button for room - %v", err)
	}

	// check if btn is displayed
	isDisplayed, err := selectBtn.IsVisible()
	if err != nil {
		return "", fmt.Errorf("failed to check if select button is displayed - %v", err)
	}
	if !isDisplayed {
		return "", fmt.Errorf("room is not available")
	}

	// scroll the select button into view
	//webdriver.ScrollElemIntoView(driver, selectBtn, "room select button")

	// click the select button
	err = selectBtn.Click()
	if err != nil {
		return "", fmt.Errorf("failed to click select button - %v", err)
	}

	// fill in initial pop up
	err = fillInitialPopUp(pw, reservationDetails)
	if err != nil {
		return "", fmt.Errorf("failed to fill initial pop up - %v", err)
	}

	// check if room was successfully selected (shows up in selected rooms)
	selectedRoomName, err := confirmSelection(pw)
	if err != nil {
		return "", err
	}

	return selectedRoomName, nil
}

// fills out the initial pop up: number of attendees & setup type
func fillInitialPopUp(pw *playwrightwrapper.PlaywrightWrapper, reservationDetails myconfig.ReservationDetails) error {
	// wait for content to load
	err := pw.WaitForElement(NUMBER_OF_ATTENDEES_INPUT_BY, NUMBER_OF_ATTENDEES_INPUT_VAL)
	if err != nil {
		return fmt.Errorf("error waiting for number of attendees input: %v", err)
	}

	// number of attendees
	err = pw.FindElemAndClick(NUMBER_OF_ATTENDEES_INPUT_BY, NUMBER_OF_ATTENDEES_INPUT_VAL)
	if err != nil {
		return fmt.Errorf("error clicking number of attendees input: %v", err)
	}
	for i := 0; i < reservationDetails.NumPeople; i++ {
		err = pw.PressKey("ArrowUp")
		if err != nil {
			return fmt.Errorf("error pressing key: %v", err)
		}
	}

	// TODO expand to be able to select setup type

	// click add room button
	pw.FindElemAndClick(ADD_ROOM_BTN_BY, ADD_ROOM_BTN_VAL)
	if err != nil {
		return fmt.Errorf("error clicking add room button: %v", err)
	}

	return nil
}

// confirms that the room was successfully selected, returns error if unsuccessful
func confirmSelection(pw *playwrightwrapper.PlaywrightWrapper) (string, error) {
	// check for alert message
	time.Sleep(500 * time.Millisecond) // allow time for alert to appear
	alertElem, _ := pw.FindElement(ALERT_BY, ALERT_VAL)
	if count, _ := alertElem.Count(); count != 0 {
		alertMessage := getAlertMessage(pw, alertElem)
		return "", fmt.Errorf("room selection failed with alert message: %s", alertMessage)
	}

	// confirm room is in selected area
	selectedRoom, _ := pw.FindElement(SELECTED_ROOM_ITEM_BY, SELECTED_ROOM_ITEM_VAL)
	if count, _ := selectedRoom.Count(); count == 0 {
		return "", fmt.Errorf("could not confirm selected room is in the selected aread")
	}

	// get selected room name
	selectedRoomName, err := selectedRoom.InnerText()
	if err != nil {
		return "", fmt.Errorf("failed to get selected room name: %v", err)
	}

	return selectedRoomName, nil
}

// gets the alert message from the alert element
func getAlertMessage(pw *playwrightwrapper.PlaywrightWrapper, alertElem playwright.Locator) string {
	alertMessageElem, err := pw.FindElementFromElement(alertElem, ALERT_MESSAGE_BY, ALERT_MESSAGE_VAL)
	if err != nil {
		return "[failed to get alert message]"
	}
	alertMessage, err := alertMessageElem.InnerText()
	if err != nil {
		alertMessage = "[failed to get alert message]"
	}
	return alertMessage
}
