package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers, double)
	tripled := doubleNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func doubleNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, i := range *numbers {
		dNumbers = append(dNumbers, transform(i))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
