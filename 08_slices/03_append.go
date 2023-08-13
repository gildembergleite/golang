package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}

    // Adicionando elementos ao slice usando append
    nums = append(nums, 4, 5, 6)
    fmt.Println(nums)  // [1 2 3 4 5 6]
}
