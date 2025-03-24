package main

import "fmt"

func main() {
	idade := 21
	fmt.Println(&idade) // address memory &

	var p *int // create a pointer *
	fmt.Println(&p)

	p = &idade      // get a address memory &
	fmt.Println(p)  // log address memory of idade
	fmt.Println(*p) // log value of idade *

	fmt.Println(idade)
	*p = 22 // change the value that the memory address stores
	fmt.Println(idade)

	p = double(p)
	fmt.Println(*p)

	p1 := Person{
		Name: "Eduardo",
		Age:  20,
	}

	fmt.Println(&p1.Name)
	fmt.Println(&p1.Age)
	fmt.Println(p1.Age)
	p1.IncrementAge()
	fmt.Println(p1.Age)
}

// pass by reference
func double(v *int) *int {
	*v *= 2
	return v
}

type Person struct {
	Name string
	Age  uint
}

func (p *Person) IncrementAge() {
	fmt.Println(&p.Name)
	p.Age += 1
}
