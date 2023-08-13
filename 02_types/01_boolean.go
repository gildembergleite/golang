package main

import "fmt"

func main() {
    // Declaração e inicialização de variáveis booleanas
    var bool1 bool = true
    var bool2 bool = false

    // Operações booleanas
    resultAnd := bool1 && bool2  // AND lógico
    resultOr := bool1 || bool2   // OR lógico
    resultNot := !bool1          // NOT lógico

    // Saída dos resultados
    fmt.Printf("bool1: %v\n", bool1)
    fmt.Printf("bool2: %v\n", bool2)
    fmt.Printf("bool1 && bool2: %v\n", resultAnd)
    fmt.Printf("bool1 || bool2: %v\n", resultOr)
    fmt.Printf("!bool1: %v\n", resultNot)
}
