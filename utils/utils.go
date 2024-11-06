package utils

import (
	"fmt"
	"net/mail"
)

type UserData struct {
	first_name   string
	last_name    string
	email        string
	ticket_count int
}

func Greet(conference_name string, conference_tickets int, remaining_tickets int) {
	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func GetFirstNames(bookings []UserData) []string {
	first_names := []string{}
	for _, booking := range bookings {
		first_names = append(first_names, booking.first_name)
	}
	return first_names
}

func PrintSummary(
	bookings []UserData, remaining_tickets int, conference_name string,
) {
	first_names := GetFirstNames(bookings)

	fmt.Printf(
		"\nSuccessfully booked %v ticket/s for %v. Confirmation sent to %v \n",
		bookings[len(bookings)-1].ticket_count,
		bookings[len(bookings)-1].first_name,
		bookings[len(bookings)-1].email)

	fmt.Printf("----------\n")
	fmt.Printf("booking reservations: %v\n", first_names)
	fmt.Printf(
		"%v tickets now remains for %v\n",
		remaining_tickets, conference_name)
	fmt.Printf("----------\n\n")
}
func BookOnce(remaining_tickets *int) UserData {
	var user_data UserData
	var _str_in string
	var _int_in int

__input_name:
	fmt.Printf("Please enter your first name: ")
	fmt.Scan(&_str_in)
	user_data.first_name = _str_in

	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&_str_in)
	user_data.last_name = _str_in

	is_valid_user_name := len(user_data.first_name) >= 2 && len(user_data.last_name) >= 2
	if !is_valid_user_name {
		fmt.Printf("!!! You seem to have entered an invalid name, please try again !!!\n")
		goto __input_name
	}

__input_email:
	fmt.Printf("Please enter your email address: ")
	fmt.Scan(&_str_in)
	user_data.email = _str_in

	_, err := mail.ParseAddress(user_data.email)
	if err != nil {
		fmt.Printf("!!! You seem to have entered an invalid email address, please try again !!!\n")
		goto __input_email
	}

__book_tickets:
	fmt.Printf("Please enter the number of tickets to book: ")
	fmt.Scan(&_int_in)
	user_data.ticket_count = _int_in

	if user_data.ticket_count > *remaining_tickets {
		fmt.Printf(
			"!!! Only %v tickets remain, please try booking with less tickets !!!\n",
			*remaining_tickets,
		)
		goto __book_tickets
	}
	*remaining_tickets -= user_data.ticket_count

	return user_data
}
