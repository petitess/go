package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"reading", "swimming", "coding"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	hobbies2 := hobbies[1:3]
	fmt.Println(hobbies2)
	hobbies3 := hobbies[:2]
	fmt.Println(hobbies3)

	goal := []string{"terratest", "golang"}
	goal[1] = "bicep"
	fmt.Println(goal)

	var products []Product
	products = append(products, Product{"book", "1", 9.99})
	products = append(products, Product{"pen", "2", 1.99})
	fmt.Println(products)

	prices := []float64{11.99, 3.99, 25.20}
	discountPrices := []float64{9.99, 1.99, 22.20}
	sum := append(prices, discountPrices...)
	fmt.Println(sum)
}
