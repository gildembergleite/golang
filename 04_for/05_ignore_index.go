package main

import "fmt"

func main() {
    // Usando o _ para ignorar o indice
    nums := []int{2, 4, 6, 8, 10}

    for _, value := range nums {
        fmt.Println(value)
    }
}
