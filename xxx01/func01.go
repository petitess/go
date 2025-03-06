// https://pkg.go.dev/std
// go mod init booking-app
// go run .
package main

import (
	"fmt"
)

func main() {
	// calculateProfit()
	text := returnValue("Hello Bob")
	fmt.Println(text)
	n1, n2 := returnNumbers(1, 2)
	fmt.Println(n1, n2)
}

func calculateProfit() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Println("How much revenue do you have?")
	fmt.Scan(&revenue)
	fmt.Println("How much do you spend?")
	fmt.Scan(&expenses)
	fmt.Println("What is your tax rate?")
	fmt.Scan(&taxRate)
	futureValue := (revenue - expenses) * taxRate / 100
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, taxRate)
	fmt.Println("After taxes", futureValue)
	fmt.Println("EBT", revenue-expenses)
}

func returnValue(text string) string {
	return text
}

func returnNumbers(n1, n2 float64) (float64, float64) {
	return n1, n2
}
