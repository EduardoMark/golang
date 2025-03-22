package main

import (
	"fmt"
)

func main() {
	x := 20
	y := 18

	if x > y {
		// executa quando a condição for true
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if y > x {
		// executa caso a codição anterior for false e essa for verdadeira
		fmt.Printf("%d is greater than %d\n", y, x)
	} else {
		// executa quando nenhuma das anteriores forem verdadeiras
		fmt.Printf("%d is equal than %d\n", x, y)
	}

	// if com inicialização de um variável, visível apenas no bloco
	count := 0
	if i := count; i > 10 {
		fmt.Println("I is greather than 10")
	}
}
