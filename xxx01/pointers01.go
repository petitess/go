package main

import "fmt"

func main() {
	age := 44

	agePointer := &age

	fmt.Println(agePointer)
	fmt.Println(*agePointer)
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
