package main

import "fmt"

func main() {
    // Percorrendo cada elemento da slice
    nums := []int{2, 4, 6, 8, 10}

    // Para cada loop, "index" receberá o indice/posição do elemento
    // e "value" receberá o valor do elemento atual
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
