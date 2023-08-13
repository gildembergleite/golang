package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Exemplo com ponteiro para int
	var num int = 42
	var ptrInt *int
	ptrInt = &num

	fmt.Println("Valor de num:", num)
	fmt.Println("Valor apontado por ptrInt:", *ptrInt)

	// Exemplo com ponteiro para float64
	var floatNum float64 = 3.14
	var ptrFloat *float64
	ptrFloat = &floatNum

	fmt.Println("Valor de floatNum:", floatNum)
	fmt.Println("Valor apontado por ptrFloat:", *ptrFloat)

	// Exemplo com ponteiro para string
	var str string = "Olá, mundo!"
	var ptrStr *string
	ptrStr = &str

	fmt.Println("Valor de str:", str)
	fmt.Println("Valor apontado por ptrStr:", *ptrStr)
	fmt.Println(reflect.TypeOf(ptrStr))
}
