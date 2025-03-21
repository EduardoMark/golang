package main

import "fmt"

// escopo de pacote
var DBName = "Test"

func main() {
	// escopo de bloco
	var name string = "Eduardo" // inferindo o tipo
	var surname = "Marques"     // o compilador infere automaticamente

	age := 21 // o tipo é inferido pelo compilador

	const PI = 3.14 // o tipo é inferido pelo compilador

	fmt.Println(DBName, name, surname, age, PI)
}
