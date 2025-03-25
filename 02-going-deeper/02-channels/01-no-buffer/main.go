package main

import (
	"fmt"
	"time"
)

// <- on the left side sends a value to the channel
// <- on the right side receives a value from the channel

func numbers(done chan<- bool) {
	for i := range 10 {
		fmt.Printf("%d ", i)
		time.Sleep(200 * time.Millisecond)
	}

	done <- true // done receive true
}

func letters(done chan<- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(300 * time.Millisecond)
	}

	done <- true // done receive true
}

func main() {
	// create a channels
	cn := make(chan bool)
	cl := make(chan bool)

	go numbers(cn)
	go letters(cl)

	// Wait for both goroutines to complete
	// The program will block here until it receives a value from both channels
	<-cn
	<-cl

	fmt.Println("The end!")
}
