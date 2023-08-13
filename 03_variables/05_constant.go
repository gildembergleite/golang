package main

import (
	"fmt"
)

// Declaração de constantes
const (
	pi       = 3.14159
	gravity  = 9.81
	username = "john_doe"
)

func main() {
	// Uso das constantes
	fmt.Printf("O valor de pi é %.2f\n", pi)
	fmt.Printf("A aceleração de gravidade é %.2f m/s²\n", gravity)
	fmt.Printf("Nome de usuário: %s\n", username)

	// Tentativa de atribuir um valor a uma constante (gerará um erro de compilação)
	// pi = 3.14

	// Tentativa de redeclarar uma constante (gerará um erro de compilação)
	// const pi = 3.14
}
