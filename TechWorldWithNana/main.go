package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get one of the %d remaining tickets here to attend\n", remainingTickets)

	for {
		var bookings []string
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Give your first name")
		fmt.Scan(&firstName)
		fmt.Println("Give your last name")
		fmt.Scan(&lastName)
		fmt.Println("Give your email")
		fmt.Scan(&email)
		fmt.Println("How many tickets")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thanks you %v %v for booking %v tickets, you will receive a email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%d tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])

		}
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
	}

}
