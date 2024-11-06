package main

import (
	"fmt"

	"booking.intercogni.com/utils"
)

func main() {
	const conference_tickets = 50

	var bookings = make([]utils.UserData, 0)

	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	utils.Greet(conference_name, conference_tickets, remaining_tickets)
	for remaining_tickets > 0 && len(bookings) < 50 {
		user_data := utils.BookOnce(&remaining_tickets)

		bookings = append(bookings, user_data)

		utils.Wg.Add(1)
		go utils.PrintSummary(bookings, remaining_tickets, conference_name)

		is_tickets_unavailable := remaining_tickets == 0
		if is_tickets_unavailable {
			fmt.Printf(
				"This month's %v is booked out. Please come again next month!\n",
				conference_name)
		}
	}

	utils.Wg.Wait()
}
