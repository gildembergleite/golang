package main

import (
	"fmt"
)

func main() {
	// Runa é um alias para o tipo int32
	var r rune
	r = 'A' // Runa representando o caractere 'A'
	fmt.Printf("Runa como inteiro: %d\n", r)

	// Representação em string
	str := string(r)
	fmt.Printf("Runa como string: %s\n", str)

	// Iterando sobre uma string usando runas
	text := "Olá, mundo!"
	for _, char := range text {
		fmt.Printf("Caractere: %c\n", char)
	}

	// Conversão de inteiro para runa
	var num int32 = 8364 // Código Unicode do símbolo do Euro
	runaEuro := rune(num)
	fmt.Printf("Runa do Euro: %c\n", runaEuro)
}
