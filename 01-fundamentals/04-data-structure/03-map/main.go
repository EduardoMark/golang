package main

import "fmt"

// key:value pairs.

func main() {
	cars := map[string]float64{}
	fmt.Println(cars)

	cars["mustang"] = 550000
	cars["ferrari"] = 880000
	// cars["mustang"] = 600000 -> not allowed duplicate, this overwrites the key

	fmt.Println(cars)

	studentsNotes := make(map[string]float64, 0)
	fmt.Println(studentsNotes)
}
