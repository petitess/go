// https://pkg.go.dev/std
// go mod init booking-app
// go run .
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 5.0
	fmt.Println("How much do you want to invest?")
	fmt.Scan(&investmentAmount)
	fmt.Println("What is the expected return rate?")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("How many years do you want to invest for?")
	fmt.Scan(&years)
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value of investment is", futureRealValue)
}

func mainX() {
	var conferanceName string = "Go Conference"
	const conferenceTickets int = 100
	var remainingTickets uint = 100

	fmt.Printf("conferenceTickets is %T,\nremainingTickets is %T\n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v\n", conferanceName)
	fmt.Println("Total Tickets:", conferenceTickets)
	fmt.Println("Remaining Tickets:", remainingTickets)

	// var bookings list = [100]string{"Nana", "Nicole"}

	var userName string
	var userTickets uint
	fmt.Println("What's you name:")
	fmt.Scan(&userName)
	fmt.Println("How many ticket do you want?")
	fmt.Scan(&userTickets)

	fmt.Println("User", userName, "booked", userTickets, "tickets")

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v ticket remaing for %v\n", remainingTickets, conferanceName)
}
