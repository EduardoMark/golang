package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := range 10 {
		fmt.Printf("%d ", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func letters() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go numbers()
	go letters()

	time.Sleep(3 * time.Second)
	fmt.Println("The end!")
}
