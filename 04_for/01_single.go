package main

import "fmt"

func main() {
    // Usando um loop "for" para imprimir números de 0 a 4
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Usando um loop "for" para imprimir números pares de 1 a 10
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}
