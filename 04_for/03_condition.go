package main

import "fmt"

func main() {
    // Usando condição simples no lugar da inicialização
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    // Usando múltiplas condições
    for i, j := 0, 10; i < 5 && j > 5; i, j = i+1, j-1 {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}
