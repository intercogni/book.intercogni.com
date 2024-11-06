package main

import "fmt"

func main() {
	const conference_tickets = 50

	var user_first_name string
	var user_last_name string
	var user_email_addr string
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
	fmt.Scan(&user_first_name)

	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&user_last_name)

	fmt.Printf("Please enter your email address: ")
	fmt.Scan(&user_email_addr)

	fmt.Printf("Please enter the number of tickets to book: ")
	fmt.Scan(&user_tickets)

	remaining_tickets -= user_tickets

	fmt.Printf(
		"Successfully booked %v ticket/s for %v %v. Confirmation sent to %v \n",
		user_tickets, user_first_name, user_last_name, user_email_addr)

	fmt.Printf(
		"%v tickets now remains for %v\n",
		remaining_tickets, conference_name)
}
