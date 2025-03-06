package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0
	fmt.Println("Wekcome to Go bank")
	for {
		fmt.Println("How can we help you?")
		fmt.Println("1. Checko balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Println("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your balance is", accountBalance)
		case 3:
			fmt.Println("How much money do you want to withdraw?")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if withdrawlAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Println("Your balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go bank")
			return
		}
	}
}
