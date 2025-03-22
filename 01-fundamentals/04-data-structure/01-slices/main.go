package main

import "fmt"

// slices has a flexible size

func main() {
	cars := []string{"mustang", "ferrari"}
	fmt.Println(cars[1], cars[0], len(cars), cap(cars))
	cars = append(cars, "BMW")
	fmt.Println(cars)

	prices := make([]float64, 3, 10)
	fmt.Println(prices)
	prices[0] = 9.99
	prices[1] = 4.35
	prices[2] = 1.29
	fmt.Println(prices)
}
