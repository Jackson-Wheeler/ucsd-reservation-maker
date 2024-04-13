package reservationMaker

import (
	"fmt"
	"time"

	"github.com/Jackson-Wheeler/ucsd-reservation-maker/myconfig"
	"github.com/Jackson-Wheeler/ucsd-reservation-maker/reservationMaker/playwrightwrapper"
)

// begin booking: create reservation btn & booking type btn
func beginBooking(pw *playwrightwrapper.PlaywrightWrapper, bookingType int) error {
	// click the create reservation button
	err := pw.FindElemAndClick(CREATE_RESERVATION_BTN_BY, CREATE_RESERVATION_BTN_VAL)
	if err != nil {
		return fmt.Errorf("error clicking create reservation button: %v", err)
	}

	// click the 'book now' button for the specified booking type
	switch bookingType {
	case BOOKING_TYPE_STUDENT_ORGS:
		err = pw.FindElemAndClick(BOOKING_TYPE_BTN_STUDENT_ORGS_BY, BOOKING_TYPE_BTN_STUDENT_ORGS_VAL)
		if err != nil {
			return fmt.Errorf("error clicking student org booking button: %v", err)
		}
	case BOOKING_TYPE_STUDY_ROOM:
		err = pw.FindElemAndClick(BOOKING_TYPE_BTN_STUDY_ROOM_BY, BOOKING_TYPE_BTN_STUDY_ROOM_VAL)
		if err != nil {
			return fmt.Errorf("error clicking study room booking button: %v", err)
		}
	default:
		return fmt.Errorf("invalid booking type: %d", bookingType)
	}
	return nil
}

// set reservation time: booking date, start time, end time, click search
func setReservationTime(pw *playwrightwrapper.PlaywrightWrapper, resTime myconfig.ReservationTime) error {
	// wait for content to load
	pw.WaitForElement(BOOKING_DATE_INPUT_BY, BOOKING_DATE_INPUT_VAL)
	time.Sleep(500 * time.Millisecond)

	// input the booking date
	err := pw.FindElemAndSendKeys(BOOKING_DATE_INPUT_BY, BOOKING_DATE_INPUT_VAL, resTime.Date)
	if err != nil {
		return fmt.Errorf("error inputting booking date: %v", err)
	}
	err = pw.PressKey("Tab")
	if err != nil {
		return fmt.Errorf("error pressing tab key: %v", err)
	}

	// input the start time
	err = pw.FindElemAndSendKeys(START_TIME_INPUT_BY, START_TIME_INPUT_VAL, resTime.StartTime)
	if err != nil {
		return fmt.Errorf("error inputting start time: %v", err)
	}

	// input the end time
	err = pw.FindElemAndSendKeys(END_TIME_INPUT_BY, END_TIME_INPUT_VAL, resTime.EndTime)
	if err != nil {
		return fmt.Errorf("error inputting end time: %v", err)
	}

	// click the search button
	err = pw.FindElemAndClick(SEARCH_BTN_BY, SEARCH_BTN_VAL)
	if err != nil {
		return fmt.Errorf("error clicking search button: %v", err)
	}

	return nil
}

// add reservation details to the booking
func addReservationDetails(pw *playwrightwrapper.PlaywrightWrapper, reservationDetails myconfig.ReservationDetails) error {
	// click the reservation details button
	err := pw.FindElemAndClick(RESERVATION_DETAILS_BTN_BY, RESERVATION_DETAILS_BTN_VAL)
	if err != nil {
		return fmt.Errorf("error clicking reservation details button: %v", err)
	}

	// input the event name
	err = pw.FindElemAndSendKeys(EVENT_NAME_INPUT_BY, EVENT_NAME_INPUT_VAL, reservationDetails.EventName)
	if err != nil {
		return fmt.Errorf("error inputting event name: %v", err)
	}

	// do nothing -> event type = study room

	// select organization - only ORGANIZATION_GROUP_STUDY_OPT_VAL is supported at this time
	err = pw.SelectFromDropdown(ORGANIZATION_INPUT_BY, ORGANIZATION_INPUT_VAL, ORGANIZATION_GROUP_STUDY_OPT_VAL)
	if err != nil {
		return fmt.Errorf("error selecting organization: %v", err)
	}

	// input the contact name
	pw.WaitForElement(CONTACT_NAME_INPUT_BY, CONTACT_NAME_INPUT_VAL)
	err = pw.FindElemAndSendKeys(CONTACT_NAME_INPUT_BY, CONTACT_NAME_INPUT_VAL, reservationDetails.ContactName)
	if err != nil {
		return fmt.Errorf("error inputting contact name: %v", err)
	}

	// input the contact phone
	err = pw.FindElemAndSendKeys(CONTACT_PHONE_INPUT_BY, CONTACT_PHONE_INPUT_VAL, reservationDetails.ContactPhone)
	if err != nil {
		return fmt.Errorf("error inputting contact phone: %v", err)
	}

	// input the contact email
	err = pw.FindElemAndSendKeys(CONTACT_EMAIL_INPUT_BY, CONTACT_EMAIL_INPUT_VAL, reservationDetails.ContactEmail)
	if err != nil {
		return fmt.Errorf("error inputting contact email: %v", err)
	}

	// select reserver status
	err = pw.SelectFromDropdown(RESERVER_STATUS_INPUT_BY, RESERVER_STATUS_INPUT_VAL, RESERVER_STATUS_STUDENT_OPT_VAL)
	if err != nil {
		return fmt.Errorf("error selecting reserver status: %v", err)
	}

	// input the description
	pw.WaitForElement(DESCRIPTION_INPUT_BY, DESCRIPTION_INPUT_VAL)
	err = pw.FindElemAndSendKeys(DESCRIPTION_INPUT_BY, DESCRIPTION_INPUT_VAL, reservationDetails.Description)
	if err != nil {
		return fmt.Errorf("error inputting description: %v", err)
	}

	return nil
}

// finish reservation: click create reservation button
func finishReservation(pw *playwrightwrapper.PlaywrightWrapper) error {
	// click the create reservation button (the first one on the page)
	finishResBtn, err := pw.FindElement(FINISH_RESERVATION_BTN_BY, FINISH_RESERVATION_BTN_VAL)
	if err != nil {
		return fmt.Errorf("error finding finish reservation button: %v", err)
	}
	err = finishResBtn.First().Click()
	if err != nil {
		return fmt.Errorf("error clicking finish reservation button: %v", err)
	}

	// dismiss the pop up
	pw.WaitForElement(OK_CONFIRMATION_BTN_BY, OK_CONFIRMATION_BTN_VAL)
	time.Sleep(500 * time.Millisecond) // wait for the pop up to fully load
	pw.FindElemAndClick(OK_CONFIRMATION_BTN_BY, OK_CONFIRMATION_BTN_VAL)

	return nil
}
