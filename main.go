package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func greet(conference_name string, conference_tickets int, remaining_tickets int) {
	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func getFirstNames(bookings []string) []string {
	first_names := []string{}
	for _, booking := range bookings {
		this_first_name := strings.Fields(booking)[0]
		first_names = append(first_names, this_first_name)
	}
	return first_names
}

func printSummary(
	user_tickets int, bookings []string, user_email_addr string,
	remaining_tickets int, conference_name string,
) {
	first_names := getFirstNames(bookings)

	fmt.Printf(
		"\nSuccessfully booked %v ticket/s for %v. Confirmation sent to %v \n",
		user_tickets, bookings[len(bookings)-1], user_email_addr)
	fmt.Printf("----------\n")
	fmt.Printf("booking reservations: %v\n", first_names)
	fmt.Printf(
		"%v tickets now remains for %v\n",
		remaining_tickets, conference_name)
	fmt.Printf("----------\n\n")
}
func bookOnce(
	user_name_first, user_name_last, user_email_addr string, user_tickets,
	remaining_tickets int, bookings []string,
) (
	string, string, string, int, int, []string,
) {
__input_name:
	fmt.Printf("Please enter your first name: ")
	fmt.Scan(&user_name_first)
	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&user_name_last)

	is_valid_user_name := len(user_name_first) >= 2 && len(user_name_last) >= 2
	if !is_valid_user_name {
		fmt.Printf("!!! You seem to have entered an invalid name, please try again !!!\n")
		goto __input_name
	}
	bookings = append(bookings, user_name_first+" "+user_name_last)
__input_email:
	fmt.Printf("Please enter your email address: ")
	fmt.Scan(&user_email_addr)
	_, err := mail.ParseAddress(user_email_addr)
	if err != nil {
		fmt.Printf("!!! You seem to have entered an invalid email address, please try again !!!\n")
		goto __input_email
	}
__book_tickets:
	fmt.Printf("Please enter the number of tickets to book: ")
	fmt.Scan(&user_tickets)
	if user_tickets > remaining_tickets {
		fmt.Printf(
			"!!! Only %v tickets remain, please try booking with less tickets !!!\n",
			remaining_tickets,
		)
		goto __book_tickets
	}
	remaining_tickets -= user_tickets

	return user_name_first, user_name_last, user_email_addr, user_tickets,
		remaining_tickets, bookings
}

func main() {
	const conference_tickets = 50

	var user_email_addr string
	var user_name_first string
	var user_name_last string
	var user_tickets int

	bookings := []string{}
	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	greet(conference_name, conference_tickets, remaining_tickets)
	for remaining_tickets > 0 && len(bookings) < 50 {

		user_name_first, user_name_last, user_email_addr, user_tickets,
			remaining_tickets, bookings = bookOnce(
			user_name_first, user_name_last, user_email_addr, user_tickets,
			remaining_tickets, bookings,
		)

		printSummary(
			user_tickets, bookings, user_email_addr,
			remaining_tickets, conference_name,
		)

		is_tickets_unavailable := remaining_tickets == 0
		if is_tickets_unavailable {
			fmt.Printf(
				"This month's %v is booked out. Please come again next month!\n",
				conference_name)
			break
		}
	}
}
