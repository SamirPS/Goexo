package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get one of the %d remaining tickets here to attend\n", remainingTickets)

	var userName string
	var userTickets uint
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %d tickets \n", userName, userTickets)

}
