package main

import "fmt"

type person struct {
	Name string
	Age  int
}

// this is a method of struct
func (p *person) ShowPersonDatails() {
	fmt.Println(p.Name, p.Age)
}
