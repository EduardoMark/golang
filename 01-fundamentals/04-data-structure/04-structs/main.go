package main

import "fmt"

// store a multiple values of different data types into a single variable

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func main() {
	p1 := Person{
		Name:   "Eduardo",
		Age:    21,
		Job:    "Software Developer",
		Salary: 10000.00,
	}

	fmt.Println(p1)

	p1.Name = "Eduardo Marques"
	p1.Salary = 15000.00

	fmt.Println(p1)
}
