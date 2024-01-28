package webdriver

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/tebeka/selenium"
)

// creates a reservation for the specified time given the room preference order and reservation details
func createReservation(driver selenium.WebDriver, resTime myconfig.ReservationTime, roomPreferenceOrder []string, reservationDetails myconfig.ReservationDetails) {

	fmt.Printf("\nCreating reservation for %s from %s to %s...\n", resTime.Date, resTime.StartTime, resTime.EndTime)

	// begin booking
	beginBooking(driver, BOOKING_TYPE_STUDY_ROOM)

	// set reservation time
	setReservationTime(driver, resTime)

	// select room
	roomName, err := selectRoom(driver, roomPreferenceOrder, reservationDetails)
	if err != nil {
		fmt.Printf("no reservation made for %s from %s to %s - %v\n", resTime.Date, resTime.StartTime, resTime.EndTime, err)
		return
	}
	fmt.Printf("selected room '%s'\n", roomName)

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

// begin booking: create reservation btn & booking type btn
func beginBooking(driver selenium.WebDriver, bookingType int) {
	var by, val string

	switch bookingType {
	case BOOKING_TYPE_STUDENT_ORGS:
		by = BOOKING_TYPE_BTN_STUDENT_ORGS_BY
		val = BOOKING_TYPE_BTN_STUDENT_ORGS_VAL
	case BOOKING_TYPE_STUDY_ROOM:
		by = BOOKING_TYPE_BTN_STUDY_ROOM_BY
		val = BOOKING_TYPE_BTN_STUDY_ROOM_VAL
	default:
		log.Fatalf("Error: invalid booking type: %d", bookingType)
	}

	// click the create reservation button
	myClickElement(driver, by, val)

	// click the study room booking button
	myClickElement(driver, by, val)
}

// set reservation time: booking date, start time, end time, click search
func setReservationTime(driver selenium.WebDriver, resTime myconfig.ReservationTime) {
	time.Sleep(500 * time.Millisecond) // TODO change to a wait for element ready

	// input the booking date
	myClearAndSendKeys(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL, resTime.Date)
	mySendKeys(driver, BOOOKING_DATE_INPUT_BY, BOOOKING_DATE_INPUT_VAL, selenium.TabKey)

	// input the start time
	myClearAndSendKeys(driver, START_TIME_INPUT_BY, START_TIME_INPUT_VAL, resTime.StartTime)

	// input the end time
	myClearAndSendKeys(driver, END_TIME_INPUT_BY, END_TIME_INPUT_VAL, resTime.EndTime)

	// click the search button
	myClickElement(driver, SEARCH_BTN_BY, SEARCH_BTN_VAL)
}

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
	roomItems := myFindElements(driver, ROOM_ITEM_BY, ROOM_ITEM_VAL)

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
