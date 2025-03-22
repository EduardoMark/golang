package main

import "fmt"

func main() {
	day := 1

	// switch com case único
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	default:
		fmt.Println("Other days")
	}

	// switch com multi cases
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Evend weekday")
	case 6, 7:
		fmt.Println("Weekend")
	}
}
