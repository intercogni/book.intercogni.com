package utils

import (
	"fmt"
	"net/mail"
	"strconv"
)

func Greet(conference_name string, conference_tickets int, remaining_tickets int) {
	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func GetFirstNames(bookings []map[string]string) []string {
	first_names := []string{}
	for _, booking := range bookings {
		first_names = append(first_names, booking["first_name"])
	}
	return first_names
}

func PrintSummary(
	bookings []map[string]string, remaining_tickets int, conference_name string,
) {
	first_names := GetFirstNames(bookings)

	fmt.Printf(
		"\nSuccessfully booked %v ticket/s for %v. Confirmation sent to %v \n",
		bookings[len(bookings)-1]["ticket_count"],
		bookings[len(bookings)-1]["first_name"],
		bookings[len(bookings)-1]["email"])

	fmt.Printf("----------\n")
	fmt.Printf("booking reservations: %v\n", first_names)
	fmt.Printf(
		"%v tickets now remains for %v\n",
		remaining_tickets, conference_name)
	fmt.Printf("----------\n\n")
}
func BookOnce(remaining_tickets *int) map[string]string {
	var user_data = make(map[string]string)
	var _in string

__input_name:
	fmt.Printf("Please enter your first name: ")
	fmt.Scan(&_in)
	user_data["first_name"] = _in

	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&_in)
	user_data["last_name"] = _in

	is_valid_user_name := len(user_data["first_name"]) >= 2 && len(user_data["last_name"]) >= 2
	if !is_valid_user_name {
		fmt.Printf("!!! You seem to have entered an invalid name, please try again !!!\n")
		goto __input_name
	}

__input_email:
	fmt.Printf("Please enter your email address: ")
	fmt.Scan(&_in)
	user_data["email"] = _in

	_, err := mail.ParseAddress(user_data["email"])
	if err != nil {
		fmt.Printf("!!! You seem to have entered an invalid email address, please try again !!!\n")
		goto __input_email
	}

__book_tickets:
	fmt.Printf("Please enter the number of tickets to book: ")
	fmt.Scan(&_in)
	user_data["ticket_count"] = _in

	intv_ticket_count, _ := strconv.Atoi(user_data["ticket_count"])
	if intv_ticket_count > *remaining_tickets {
		fmt.Printf(
			"!!! Only %v tickets remain, please try booking with less tickets !!!\n",
			*remaining_tickets,
		)
		goto __book_tickets
	}
	*remaining_tickets -= intv_ticket_count

	return user_data
}
