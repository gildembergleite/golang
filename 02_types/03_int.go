package main

import (
	"fmt"
	"math"
)

func main() {
	// Tipos inteiros com sinal
	var int8Var int8 = math.MinInt8
	var int16Var int16 = math.MinInt16
	var int32Var int32 = math.MinInt32
	var int64Var int64 = math.MinInt64

	// Tipos inteiros sem sinal
	var uint8Var uint8 = math.MaxUint8
	var uint16Var uint16 = math.MaxUint16
	var uint32Var uint32 = math.MaxUint32
	var uint64Var uint64 = math.MaxUint64

	// Tipo int é dependente da arquitetura (32 ou 64 bits)
	var intVar int = math.MinInt32

	fmt.Println("Inteiro com sinal:")
	fmt.Printf("int8: %d\n", int8Var)
	fmt.Printf("int16: %d\n", int16Var)
	fmt.Printf("int32: %d\n", int32Var)
	fmt.Printf("int64: %d\n", int64Var)

	fmt.Println("\nInteiro sem sinal:")
	fmt.Printf("uint8: %d\n", uint8Var)
	fmt.Printf("uint16: %d\n", uint16Var)
	fmt.Printf("uint32: %d\n", uint32Var)
	fmt.Printf("uint64: %d\n", uint64Var)

	fmt.Println("\nInteiro dependente da arquitetura:")
	fmt.Printf("int: %d\n", intVar)
}
