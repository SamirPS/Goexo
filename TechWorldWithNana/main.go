package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket uint = 50

var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			fmt.Printf("The first names of the bookings are: %v\n", getFirstNames())

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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get one of the %d remaining tickets here to attend.We have %d places\n", remainingTickets, conferenceTicket)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thanks you %v %v for booking %v tickets, you will receive a email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %v\n", remainingTickets, conferenceName)
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
