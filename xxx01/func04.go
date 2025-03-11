package main

import "fmt"

func main() {
	sum := sumup(1, 12, 23, 11)
	fmt.Print(sum)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
