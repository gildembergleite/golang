package main

import (
	"fmt"
)

func main() {
	// Incluindo mais de uma condição em um mesmo if/else-if
	valor := 7
	if valor > 10 && valor%2 == 0 {
		fmt.Println("O valor é maior que 10 e é par.")
	} else if valor > 10 || valor%2 == 0 {
		fmt.Println("O valor é maior que 10 ou é par.")
	} else {
		fmt.Println("O valor não é maior que 10 nem é par.")
	}
}
