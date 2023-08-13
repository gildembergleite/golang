package main

import "fmt"

func main() {
    originalSlice := []int{1, 2, 3, 4, 5}

    // Criando um slice usando uma parte do original
    newSlice := originalSlice[1:4]  // [2 3 4]
    fmt.Println(newSlice)

    // Modificando o novo slice
    newSlice[0] = 10
    fmt.Println(originalSlice)  // [1 10 3 4 5]
}
