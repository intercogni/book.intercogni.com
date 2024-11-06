package main

import "fmt"

func main() {
	const conference_tickets = 50

	var user_name string
	var user_tickets int

	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	user_name = "Taib"
	user_tickets = 1

	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("User %v booked %v ticket/s\n", user_name, user_tickets)
}
