package main

import "fmt"

// Main entrypoint
func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}
