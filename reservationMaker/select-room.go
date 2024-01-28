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
	time.Sleep(3 * time.Second) // TODO change to a wait for element ready

	// map room names to web page room elements
	roomMap := getRoomMap(driver)

	// attempt selecting rooms in order of preference
	for _, roomName := range roomPreferenceOrder {
		err := attemptSelectRoom(driver, roomName, roomMap[roomName])
		if err != nil {
			fmt.Println(err)
			continue
		}
		// successfully selected room
		return roomName, nil
	}

	// if preferred rooms are not available
	return "", errors.New("preferred rooms are not available")

}

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

// - find the child with class "fa-plus-circle"
// - click it
// - if fail, go on to next
// - else, fill out pop up window & check that selected rooms has a room in it, if not, go on to next, else, break out of loop

// attempts to click to select the given room, returns error if fails
func attemptSelectRoom(driver selenium.WebDriver, roomName string, roomItem selenium.WebElement) error {
	const (
		ROOM_SELECT_BTN_CSS_SELECTOR = ".fa-plus-circle"
		SELECTED_ROOM_CSS_SELECTOR   = ".selected-room-item"
	)
	// find child button for selecting room
	selectBtn := webdriver.FindElementFromElement(roomItem, selenium.ByCSSSelector, ROOM_SELECT_BTN_CSS_SELECTOR)

	// check if it is displayed
	isDisplayed := webdriver.IsDisplayed(selectBtn)
	fmt.Printf("selectBtn for room %s is displayed: %t\n", roomName, isDisplayed)
	if !isDisplayed {
		return fmt.Errorf("room '%s' is not available", roomName)
	}

	// scroll the select button into view
	webdriver.ScrollElemIntoView(driver, selectBtn, "room select button")

	// click the select button
	err := selectBtn.Click()
	if err != nil {
		return fmt.Errorf("failed to click select button for room '%s' - %v", roomName, err)
	}

	// check if room was successfully selected (shows up in selected rooms)
	// selectedRooms := myFindElements(driver, selenium.ByCSSSelector, SELECTED_ROOM_CSS_SELECTOR)
	// if len(selectedRooms) == 0 {
	// 	return fmt.Errorf("room '%s' was not successfully selected", roomName)
	// }

	return nil
}

// number of attendees id="setup-add-count" - clear & send keys
// add room btn id="setup--add-modal-save" - click
