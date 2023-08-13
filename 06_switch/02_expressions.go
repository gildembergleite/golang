package main

import "fmt"

func main() {
		// Switch com expressões simples
    valor := 42

    switch {
    case valor < 0:
        fmt.Println("Valor é negativo")
    case valor == 0:
        fmt.Println("Valor é zero")
    case valor > 0:
        fmt.Println("Valor é positivo")
    }

		// Utilizando múltiplas expressões
		idade := 18

    switch idade {
    case 16, 17:
        fmt.Println("Pode dirigir com permissão")
    case 18, 19, 20:
        fmt.Println("Pode votar")
    default:
        fmt.Println("Outra idade")
    }
}
