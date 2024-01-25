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
	// - find the child with class "fa-plus-circle"
	// - click it
	// - if fail, go on to next
	// - else, fill out pop up window & check that selected rooms has a room in it, if not, go on to next, else, break out of loop
	// (above four points -> function returning err, selectRoom() - error describes what went wrong)
	roomItems := myFindElements(driver, selenium.ByCSSSelector, ROOM_ITEM_CSS_SELECTOR)
	fmt.Printf("Length of room items: %d\n", len(roomItems))

	roomMap := make(map[string]selenium.WebElement)
	for _, roomItem := range roomItems {
		title, err1 := roomItem.GetAttribute("title")
		if err1 != nil {
			log.Fatal("Error: failed to get room item text - ", err1)
		}
		fmt.Printf("Room Item Text: %s\n", title)
		roomMap[title] = roomItem
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

	time.Sleep(5 * time.Second)
}
