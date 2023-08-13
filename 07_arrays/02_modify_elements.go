package main

import "fmt"

func main() {
    intArray := [5]int{10, 20, 30, 40, 50}
    intArray[2] = 35
    fmt.Println(intArray) // Saída: [10 20 35 40 50]
}
