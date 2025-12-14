package main

import (
	"Booking-app/helpers"
	"fmt"
	"strings"
)

func main() {
	bootApplication()
	greetUser(ConferenceName, conferenceTickets, remainingTickets)

	for {
		if remainingTickets == 0 {
			fmt.Println("Sorry, our conference is fully booked! Come back next year.")
			printFirstNames(bookings)
			break
		}
		firstName, lastName, email, userTickets := getUserInput()

		// ============== input Validation ==================
		if !helpers.IsOnlyLetters(firstName) || len(firstName) < 2 {
			fmt.Println("Invalid first name. Must contain only letters and be at least 2 characters.")
			continue
		}

		if !helpers.IsOnlyLetters(lastName) || len(lastName) < 2 {
			fmt.Println("Invalid last name. Must contain only letters and be at least 2 characters.")
			continue
		}

		if !strings.Contains(email, "@") || len(email) < 3 {
			fmt.Println("Invalid email address. It must contain '@' and be at least 3 characters long.")
			continue
		}
		bookings = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)
		var firstNames []string
		for _, booking := range bookings {
			names := strings.Fields(booking)
			if len(names) > 0 {
				firstNames = append(firstNames, names[0])
			}
		}
		if userTickets < 1 || userTickets > remainingTickets {
			fmt.Printf("Invalid number of tickets. We have only %v remaining.\n", remainingTickets)
			continue
		}

		// Valid booking → process it
		bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)

		printUsers(firstName, lastName, email, userTickets)
		if remainingTickets == 0 {
			fmt.Println("\nAll tickets sold out! See you at the conference!")
			break
		}
	}
}
