package main

import "fmt"

func main() {
    // Declaração e inicialização de variáveis dos tipos float
    var f32 float32 = 3.14159
    var f64 float64 = 2.71828

    // Impressão dos valores das variáveis
    fmt.Printf("Valor float32: %f\n", f32)
    fmt.Printf("Valor float64: %f\n", f64)

    // Operações matemáticas simples
    soma := f32 + float32(f64) // Precisamos converter f64 para float32 para realizar a operação
    subtracao := f64 - float64(f32) // Precisamos converter f32 para float64 para realizar a operação
    multiplicacao := f32 * float32(f64)
    divisao := f64 / float64(f32)

    fmt.Printf("Soma: %f\n", soma)
    fmt.Printf("Subtração: %f\n", subtracao)
    fmt.Printf("Multiplicação: %f\n", multiplicacao)
    fmt.Printf("Divisão: %f\n", divisao)
}
