package main

import (
    "fmt"
)

func main() {
    // Practice 1: Print Hello World
    fmt.Println("Hello World")

    // Practice 2: Declare Variables
    var name string = "John"
    var age int = 30
    var height float64 = 1.75
    fmt.Printf("My name is %s, I'm %d years old and %.2f meters tall\n", name, age, height)

    // Practice 3: Slice and Range
    letters := []string{"a", "b", "c"}
    for i, letter := range letters {
        fmt.Printf("Index: %d, Letter: %s\n", i, letter)
    }

    // Practice 4: Function and Recursion
    fmt.Printf("The factorial of 5 is %d\n", factorial(5))
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}