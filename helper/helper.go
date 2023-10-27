package helper

import (
	"fmt"
	"strconv"
)

type UserData struct {
	firstName       string
	lastName        string
	numberOfTickets uint
}

// function name starts with a capital letter for export reasons
func BookTicket(bookings []map[string]string, firstName string, lastName string, userTickets uint, remainingTickets uint) ([]map[string]string, uint) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

	return bookings, remainingTickets
}

func BookTicketV2(bookings []UserData, firstName string, lastName string, userTickets uint, remainingTickets uint) ([]UserData, uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

	return bookings, remainingTickets
}
