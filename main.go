package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "Go conference"

const totalNumberofTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)
var bookingsV2 = make([]helper.UserData, 0)

func main() {
	greetUser()

	for {
		firstName, lastName, userTickets := getUserInput()

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
			continue
		}

		// bookings, remainingTickets = helper.BookTicket(bookings, firstName, lastName, userTickets, remainingTickets)
		bookingsV2, remainingTickets = helper.BookTicketV2(bookingsV2, firstName, lastName, userTickets, remainingTickets)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Conference booked out")
			break
		}
	}
}

func greetUser() {
	// string formatting using Printf
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total number of %v tickets and %v are still available\n", totalNumberofTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Println("fisrtName")
	fmt.Scan(&firstName)
	fmt.Println("lastName")
	fmt.Scan(&lastName)
	fmt.Println("userTickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets

}
