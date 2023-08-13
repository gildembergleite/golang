package main

import "fmt"

func main() {
    source := []int{1, 2, 3}
    destination := make([]int, len(source))

    // Copiando elementos do source para o destination
    copy(destination, source)
    fmt.Println(destination)  // [1 2 3]
}
