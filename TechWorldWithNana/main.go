package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get one of the %d tickets here to attend\n", remainingTickets)

}
