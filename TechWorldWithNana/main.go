package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, remainingTickets, conferenceTicket)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			fmt.Printf("The first names of the bookings are: %v\n", getFirstNames(bookings))

			if remainingTickets == 0 {
				fmt.Printf("Sold out")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Println("Email invalid")
			}
			if !isValidName {
				fmt.Println("Name invalid")
			}

			if !isValidTicketNumber {
				fmt.Println("Ticket number invalid")
			}
		}

	}

}

func greetUsers(conferenceName string, remainingTickets uint, conferenceTicket uint) {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get one of the %d remaining tickets here to attend.We have %d places\n", remainingTickets, conferenceTicket)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thanks you %v %v for booking %v tickets, you will receive a email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %v\n", remainingTickets, conferenceName)
}
