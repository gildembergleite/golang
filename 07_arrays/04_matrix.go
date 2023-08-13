package main

import "fmt"

func main() {
    // Array bidimensional
    var matrix [3][3]int

    // Inicialização do array bidimensional
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            matrix[i][j] = i + j
        }
    }

    fmt.Println(matrix)
}
