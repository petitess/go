package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string
	productNames = [4]string{"Apple", "Banana", "Orange", "Grapes"}
	productNames[2] = "Mango"
	prices := [4]float64{100.0, 200.0, 300.0, 400.0}
	fmt.Println(prices)
	fmt.Println(productNames)

	slicedPrices := prices[1:3]
	fmt.Println(slicedPrices)
	fmt.Println(len(slicedPrices), cap(slicedPrices))
}
