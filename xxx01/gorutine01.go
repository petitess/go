package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go slowGreet("Karol", done)
	go greet("Anna", done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}

func greet(text string, doneChan chan bool) {
	fmt.Println("Hello,", text)
	doneChan <- true
}

func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello,", text)
	doneChan <- true
}
