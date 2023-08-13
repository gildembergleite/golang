package main

import (
	"fmt"
)

func main() {
	// Utilizando else-if para acrescentar diferentes condições e retornos
	nota := 75
	if nota >= 90 {
		fmt.Println("Ótima nota!")
	} else if nota >= 70 {
		fmt.Println("Boa nota.")
	} else {
		fmt.Println("Precisa melhorar.")
	}
}
