package main

import (
	"fmt"
)

func main() {
	// Array de inteiros
	intArray := [5]int{1, 2, 3, 4, 5}

	// Array de strings
	stringArray := [3]string{"hello", "world", "go"}

	// Array de floats
	floatArray := [4]float64{1.1, 2.2, 3.3, 4.4}

	// Array de booleanos
	boolArray := [2]bool{true, false}

	// Array de caracteres (runas)
	runeArray := [3]rune{'a', 'b', 'c'}

	// Imprimindo os arrays
	fmt.Println("IntArray:", intArray)
	fmt.Println("StringArray:", stringArray)
	fmt.Println("FloatArray:", floatArray)
	fmt.Println("BoolArray:", boolArray)
	fmt.Println("RuneArray:", runeArray)
}
