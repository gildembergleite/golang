package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}

    // Acessando elementos
    fmt.Println(numbers[0])  // 10
    fmt.Println(numbers[2])  // 30

    // Modificando elementos
    numbers[1] = 25
    fmt.Println(numbers)     // [10 25 30 40 50]
}
