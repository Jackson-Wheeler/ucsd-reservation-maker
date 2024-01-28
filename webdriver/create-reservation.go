package webdriver

import (
	"fmt"
	"log"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, resTime myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {
	const (
		CREATE_RESERVATION_BTN_TEXT  = "Create A Reservation"
		STUDY_ROOM_BOOKING_BTN_XPATH = `//*[@aria-label='Book Now With The "Reserve Spaces | Study Rooms & Open Desk" Template']`
		BOOOKING_DATE_INPUT_ID       = "booking-date-input"
		START_TIME_INPUT_ID          = "start-time-input"
		END_TIME_INPUT_ID            = "end-time-input"
		SEARCH_BTN_XPATH             = `//button[normalize-space()='Search']`
		ROOM_ITEM_CSS_SELECTOR       = ".room-column.column"
	)

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", resTime.Date, resTime.StartTime, resTime.EndTime)

	// click the create reservation button
	myClickElement(driver, selenium.ByLinkText, CREATE_RESERVATION_BTN_TEXT)

	// click the study room booking button
	myClickElement(driver, selenium.ByXPATH, STUDY_ROOM_BOOKING_BTN_XPATH)

	time.Sleep(500 * time.Millisecond) // TODO change to a wait for element ready

	// input the booking date
	myClearAndSendKeys(driver, selenium.ByID, BOOOKING_DATE_INPUT_ID, resTime.Date)
	mySendKeys(driver, selenium.ByID, BOOOKING_DATE_INPUT_ID, selenium.TabKey)

	// input the start time
	myClearAndSendKeys(driver, selenium.ByID, START_TIME_INPUT_ID, resTime.StartTime)

	// input the end time
	myClearAndSendKeys(driver, selenium.ByID, END_TIME_INPUT_ID, resTime.EndTime)

	// click the search button
	myClickElement(driver, selenium.ByXPATH, SEARCH_BTN_XPATH)

	time.Sleep(3 * time.Second) // TODO change to a wait for element ready

	// get all room items
	// TODO get room-column items & save & print their title attribute (make struct for a room)
	// for each room name in roomPreferenceOrder, find that room:
	// (above four points -> function returning err, selectRoom() - error describes what went wrong)

	// get all rooms
	roomItems := myFindElements(driver, selenium.ByCSSSelector, ROOM_ITEM_CSS_SELECTOR)

	// make map of room name to room element (from web page)
	roomMap := make(map[string]selenium.WebElement)
	for _, roomItem := range roomItems {
		title, err1 := roomItem.GetAttribute("title")
		if err1 != nil {
			log.Fatal("Error: failed to get room item text - ", err1)
		}
		roomMap[title] = roomItem
	}

	// attempt booking rooms (in order of preference)
	success := false
	for _, roomName := range roomPreferenceOrder {
		err := selectRoom(driver, roomName, roomMap[roomName])
		if err != nil {
			fmt.Println(err)
			continue
		}
		success = true
		break
	}

	// if no room available, return
	if !success {
		fmt.Printf("*reservation not created - no room available on %s from %s to %s\n", resTime.Date, resTime.StartTime, resTime.EndTime)
		return
	}

	// Reservation Details, button aria-label="Create a Reservation/Reservation Details"
	// event name id="event-name" - clear & send keys
	// event type id="event-type" - just do nothing if event type is study room
	// organization id="availablegroups" - click, down arrow twice, enter
	// id="1stContactName" - clear & send keys
	// id="1stContactPhone1" - clear & send keys
	// id="1stContactEmail"	- clear & send keys
	// choose who you are id="28" - click, down arrow three times, enter
	// description id="27" - clear & send keys

	// Create reservation button, data-bind="click: function(){ return saveReservation(); }" - click
	// confirm reservation was made
}

// - find the child with class "fa-plus-circle"
// - click it
// - if fail, go on to next
// - else, fill out pop up window & check that selected rooms has a room in it, if not, go on to next, else, break out of loop

// attempts to click to select the given room, returns error if fails
func selectRoom(driver selenium.WebDriver, roomName string, roomItem selenium.WebElement) error {
	const (
		ROOM_SELECT_BTN_CSS_SELECTOR = ".fa-plus-circle"
		SELECTED_ROOM_CSS_SELECTOR   = ".selected-room-item"
	)
	// find child button for selecting room
	selectBtn := myFindElementFromElement(roomItem, selenium.ByCSSSelector, ROOM_SELECT_BTN_CSS_SELECTOR)

	// check if it is displayed
	isDisplayed := myIsDisplayed(selectBtn)
	fmt.Printf("selectBtn for room %s is displayed: %t\n", roomName, isDisplayed)
	if !isDisplayed {
		return fmt.Errorf("room '%s' is not available", roomName)
	}

	// scroll the select button into view
	myScrollElemIntoView(driver, selectBtn, "room select button")

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
