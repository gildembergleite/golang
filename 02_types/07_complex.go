package main

import "fmt"

// Definindo uma estrutura (struct)
type Person struct {
    Name    string
    Age     int
    Address Address
}

type Address struct {
    Street  string
    City    string
    Country string
}

func main() {
    // Inicializando uma estrutura (struct)
    person := Person{
        Name: "João",
        Age:  30,
        Address: Address{
            Street:  "123 Main St",
            City:    "Exampleville",
            Country: "Exampleland",
        },
    }

    // Criando uma slice de inteiros
    numbers := []int{1, 2, 3, 4, 5}

    // Criando um mapa (map) de strings para inteiros
    scores := map[string]int{
        "Matemática": 90,
        "Ciências":   85,
        "História":   78,
    }

    // Criando um array de strings
    colors := [3]string{"Vermelho", "Verde", "Azul"}

    // Criando um ponteiro para um inteiro
    var agePointer *int
    age := 25
    agePointer = &age

    // Imprimindo os valores
    fmt.Println("Nome:", person.Name)
    fmt.Println("Idade:", person.Age)
    fmt.Println("Endereço:", person.Address.Street, person.Address.City, person.Address.Country)
    fmt.Println("Números:", numbers)
    fmt.Println("Notas:", scores)
    fmt.Println("Cores:", colors)
    fmt.Println("Idade (ponteiro):", *agePointer)
}
