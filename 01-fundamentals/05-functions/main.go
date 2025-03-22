package main

import "fmt"

// main is the pricipal func in go
func main() {
	f1("Eduardo")
	f2("Eduardo", "Marques", 21)
	sum := f3(10, 40)
	result := f4(5, 4)
	s, t := f5(50, "Hello")

	fmt.Println(sum)
	fmt.Println(result)
	fmt.Println(s, t)

	p1 := person{
		Name: "Eduardo",
		Age:  21,
	}

	p1.ShowPersonDatails()
}

func f1(name string) {
	fmt.Println("Hello", name)
}

func f2(name, surname string, age int) {
	fmt.Printf("%s %s has %d years old\n", name, surname, age)
}

func f3(a, b int) int {
	return a + b
}

func f4(a, b int) (result int) {
	result = a + b
	return // equal (return result)
}

func f5(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
