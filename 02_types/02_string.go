package main

import "fmt"

func main() {
    // Strings literais
    literalString := "Hello, world!"

    // Strings de várias linhas
    multiLineString := `
        This is a multi-line
        string example.`

    // Strings não interpretadas (raw)
    rawString := `C:\Users\Username\Documents`

    // Strings com interpolação
    name := "Alice"
    interpolatedString := fmt.Sprintf("Hello, %s!", name)

    // Concatenação de strings
    string1 := "Hello, "
    string2 := "Golang!"
    concatenatedString := string1 + string2

    // Indexação e iteração de strings
    str := "Golang"
    for i := 0; i < len(str); i++ {
        fmt.Printf("%c ", str[i])
    }
    fmt.Println()

    // Convertendo string para slice de bytes (byte array)
    byteArray := []byte("Byte Array")

    // Convertendo slice de bytes para string
    byteSlice := []byte{'B', 'y', 't', 'e', 's'}
    convertedString := string(byteSlice)

    fmt.Println("Literal string:", literalString)
    fmt.Println("Multi-line string:", multiLineString)
    fmt.Println("Raw string:", rawString)
    fmt.Println("Interpolated string:", interpolatedString)
    fmt.Println("Concatenated string:", concatenatedString)
    fmt.Println("Indexing and iterating:")
    for _, char := range str {
        fmt.Printf("%c ", char)
    }
    fmt.Println()
    fmt.Println("Byte array:", byteArray)
    fmt.Println("Converted string:", convertedString)
}
