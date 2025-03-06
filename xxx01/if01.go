package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0
	fmt.Println("Wekcome to Go bank")
	fmt.Println("How can we help you?")
	fmt.Println("1. Checko balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Your choice:")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit:")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 3 {
		fmt.Println("How much money do you want to withdraw?")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)
		if withdrawlAmount <= 0 {
			fmt.Println("Invalid amount")
			return
		}
		if withdrawlAmount > accountBalance {
			fmt.Println("Insufficient funds")
			return
		}
		accountBalance -= withdrawlAmount
		fmt.Println("Your balance is", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
