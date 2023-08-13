package main

import "fmt"

func main() {
    // Exemplo de loop infinito
    i := 0
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break
        }
    }
}
