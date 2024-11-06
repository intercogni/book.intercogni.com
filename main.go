package main

import (
	"fmt"

	"booking.intercogni.com/utils"
)

func main() {
	const conference_tickets = 50

	var user_email_addr string
	var user_name_first string
	var user_name_last string
	var user_tickets int

	bookings := []string{}
	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	utils.Greet(conference_name, conference_tickets, remaining_tickets)
	for remaining_tickets > 0 && len(bookings) < 50 {
		user_name_first, user_name_last, user_email_addr, user_tickets,
			remaining_tickets, bookings = utils.BookOnce(
			user_name_first, user_name_last, user_email_addr, user_tickets,
			remaining_tickets, bookings,
		)

		utils.PrintSummary(
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
