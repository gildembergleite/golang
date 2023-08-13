package main

import (
	"fmt"
)

func main() {
	// Declarando e inicializando múltiplas variáveis em uma única linha
	var nome, sobrenome string = "João", "Silva"
	idade := 30 // Forma curta de declarar e inicializar uma variável do tipo int

	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Idade:", idade)
}
