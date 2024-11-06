package main

import "fmt"

func main() {
	const conference_tickets = 50

	var bookings []string
	var user_email_addr string
	var user_name_first string
	var user_name_last string
	var user_tickets int

	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	// user_name = "Taib"
	// user_tickets = 1

	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("Please enter your first name: ")
	fmt.Scan(&user_name_first)
	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&user_name_last)
	bookings = append(bookings, user_name_first+" "+user_name_last)

	fmt.Printf("Please enter your email address: ")
	fmt.Scan(&user_email_addr)

	fmt.Printf("Please enter the number of tickets to book: ")
	fmt.Scan(&user_tickets)

	remaining_tickets -= user_tickets

	fmt.Printf(
		"Successfully booked %v ticket/s for %v. Confirmation sent to %v \n",
		user_tickets, bookings[0], user_email_addr)

	fmt.Printf(
		"%v tickets now remains for %v\n",
		remaining_tickets, conference_name)

	fmt.Printf("bookings array data: %v\n", bookings)
	fmt.Printf("bookings array type: %T\n", bookings)
	fmt.Printf("bookings array size: %v\n", len(bookings))
}
