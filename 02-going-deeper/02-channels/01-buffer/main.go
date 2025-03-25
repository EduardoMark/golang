package main

import (
	"fmt"
	"time"
)

func numbers(c chan<- int) {
	for i := range 10 {
		c <- i
		fmt.Printf("Write number in buffer: %d\n", i)
	}

	close(c) // closing channel
}

func main() {
	cn := make(chan int, 2) // create a channel with buffer
	go numbers(cn)          // run a goroutine

	time.Sleep(2 * time.Second)
	for v := range cn {
		fmt.Println("Read value of channel:", v)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("The end!")
}
