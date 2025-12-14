package main

import (
	"fmt"
	"strconv"
	"strings"
)

const ConferenceName = "Go-Conference"
const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings []string

func bootApplication() {
	fmt.Println("============= Booting Application ==============")
}

func greetUser(confName string, ticket uint, confTicket uint) {
	fmt.Printf("Welcome to %v Booking Application.\n", confName)
	fmt.Printf("We have total %v tickets, and %v are still available.\n", ticket, confTicket)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you want? ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

// booked ticket
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName, lastName, email string) []string {

	remainingTickets -= userTickets
	booking := fmt.Sprintf("%s %s (%s)", firstName, lastName, email)
	bookings = append(bookings, booking)

	fmt.Printf("\nThank you %s %s! You have booked %v ticket(s).\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation email will be sent to %s\n", email)
	fmt.Printf("Remaining : %v\n", remainingTickets)
	return bookings
}

func printFirstNames(bookings []string) []string {
	var firstNames []string
	for _, booking := range bookings {
		names := strings.Fields(booking)
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
	}
	return firstNames
}

func list0fUsers(firstName, lastName, email string, userTickets uint) map[string]string {
	// store user activity using map
	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["LastName"] = lastName
	userData["Email"] = email
	userData["UserTicket"] = strconv.FormatUint(uint64(userTickets), 10)

	return userData
}

func printUsers(firstName, lastName, email string, userTickets uint) {
	data := list0fUsers(firstName, lastName, email, userTickets)
	fmt.Println("\n ======== YOUR information =========")
	for i, v := range data {
		fmt.Printf("%-28s : %s\n", v, i)
	}
}

func firstNAme() {
	// Helper function to print first names
	fNames := printFirstNames(bookings)
	fmt.Printf("Current bookings: %v\n", fNames)
}

