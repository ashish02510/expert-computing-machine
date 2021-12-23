package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 30
	fmt.Println("welcome to ", conferenceName, "booking application")
	fmt.Println("We have total of ", conferenceTickets, " Tickets and ", remainingTickets, "are still available")

}
