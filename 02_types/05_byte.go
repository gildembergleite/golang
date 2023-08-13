package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Diferentes tipos de dados relacionados a bytes

	// Slice de bytes
	byteSlice := []byte{65, 66, 67} // Representa 'A', 'B' e 'C' em ASCII
	fmt.Println("Slice de bytes:", byteSlice)

	// Array de bytes
	var byteArray [3]byte
	byteArray[0] = 97 // Representa 'a' em ASCII
	byteArray[1] = 98 // Representa 'b' em ASCII
	byteArray[2] = 99 // Representa 'c' em ASCII
	fmt.Println("Array de bytes:", byteArray)

	// String
	str := "Olá, mundo!"
	byteString := []byte(str)
	fmt.Println("String em bytes:", byteString)

	// Conversões entre tipos
	byteValue := byte(65)
	fmt.Println("Valor byte:", byteValue)

	// Tamanho em bytes dos tipos
	sizeOfByte := unsafe.Sizeof(byte(0))
	sizeOfByteArray := unsafe.Sizeof(byteArray)
	sizeOfString := unsafe.Sizeof(str)
	fmt.Printf("Tamanho de byte: %d bytes\n", sizeOfByte)
	fmt.Printf("Tamanho de array de bytes: %d bytes\n", sizeOfByteArray)
	fmt.Printf("Tamanho de string: %d bytes\n", sizeOfString)
}
