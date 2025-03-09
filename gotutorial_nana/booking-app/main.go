package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println("Get your ticket here to attend")

	var fistName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&fistName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you need? ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	fmt.Printf("User %v booked %v tickets.\n", fistName, userTickets)

	fmt.Printf(
		"We have total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets,
	)
}
