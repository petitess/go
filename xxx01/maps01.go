package main

import "fmt"

type stringMap map[string]string

func main() {
	websites := map[string]string{
		"google":   "google.com",
		"facebook": "facebook.com",
		"amazon":   "twitter.com",
	}
	websites["twitter"] = "twitter.com"
	delete(websites, "amazon")
	fmt.Println(websites)
	fmt.Println(websites["google"])

	userNames := make(stringMap)
	userNames["first"] = "John"
	userNames["last"] = "Doe"
	userNames.printMethod()

	for _, value := range websites {
		fmt.Println(value)
	}
}

func (sm stringMap) printMethod() {
	fmt.Println(sm)
}
