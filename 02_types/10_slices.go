package main

import "fmt"

func main() {
    // Slice de inteiros
    intSlice := []int{1, 2, 3, 4, 5}

    // Slice de strings
    stringSlice := []string{"apple", "banana", "cherry", "date"}

    // Slice de floats
    floatSlice := []float64{3.14, 1.618, 2.718}

    // Slice de booleanos
    boolSlice := []bool{true, false, true, false}

    // Slice de bytes
    byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}

    // Imprimindo os slices
    fmt.Println("intSlice:", intSlice)
    fmt.Println("stringSlice:", stringSlice)
    fmt.Println("floatSlice:", floatSlice)
    fmt.Println("boolSlice:", boolSlice)
    fmt.Println("byteSlice:", byteSlice)
}
