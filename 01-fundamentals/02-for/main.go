package main

import "fmt"

func main() {
	// for tradicional
	for i := 0; i < 5; i++ {
		fmt.Print(i) // 1 2 3 4 5
	}

	fmt.Println("\n-------------------")

	// for range (array, slice, string, map, channel)
	club := "Real Madrid"
	for i, value := range club {
		fmt.Println(i, "-", string(value))
	}

	fmt.Println("-------------------")

	// for condicional
	count := 0
	for count < 10 {
		count++
	}

	// for
	loop := 1
	for {
		fmt.Print(loop)
		loop++
		if loop == 5 {
			break
		}
	}
}
