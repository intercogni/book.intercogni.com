package main

import "fmt"

func main() {
	var conference_name = "Intercogni Monthly"
	const conference_tickets = 50
	var remaining_tickets = 50

	fmt.Println("Sign up for", conference_name, "for IoT and Robotics")
	fmt.Println(
		conference_tickets, "are given this month and", 
		remaining_tickets, "are still available"
	)
	fmt.Println("Get your tickets here to attend")
}
