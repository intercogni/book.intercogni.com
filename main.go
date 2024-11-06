package main

import "fmt"

func main() {
	var conference_name = "Intercogni Monthly"
	const conference_tickets = 50
	var remaining_tickets = 50

	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")
}
