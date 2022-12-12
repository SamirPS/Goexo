package main

import (
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTicket uint = 50

var remainingTickets uint = 50
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
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
		firstNames = append(firstNames, booking.firstName)

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

	// create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thanks you %v %v for booking %v tickets, you will receive a email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending %v\n to email: %v\n", ticket, email)
	fmt.Println("################")

}
