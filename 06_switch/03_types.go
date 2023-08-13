package main

import "fmt"

func main() {
    var x interface{} = "Olá"

    switch x.(type) {
    case int:
        fmt.Println("É um inteiro")
    case string:
        fmt.Println("É uma string")
    default:
        fmt.Println("Tipo desconhecido")
    }
}
