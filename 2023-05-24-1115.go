package main

import (
    "fmt"
)

func main() {
    // Example 1: sum of two integers
    x := 10
    y := 20
    sum := x + y
    fmt.Println(sum)

    // Example 2: find the maximum of three integers
    a := 5
    b := 3
    c := 7
    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    fmt.Println(max)

    // Example 3: reverse a string
    str := "Hello, World!"
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    reversedStr := string(runes)
    fmt.Println(reversedStr)
}