package main

import "fmt"

func main() {
	// Exemplo simples do switch com nomenclatura
    diaDaSemana := 3

    switch diaDaSemana {
    case 1:
        fmt.Println("Domingo")
    case 2:
        fmt.Println("Segunda-feira")
    case 3:
        fmt.Println("Terça-feira")
    // ... outros casos ...
    default:
        fmt.Println("Dia inválido")
    }

    // Exemplo de switch sem nomenclatura
    booleano := true

    switch {
    case booleano == true :
        fmt.Println("Verdadeiro")
    default:
        fmt.Println("Falso")
    }
}
