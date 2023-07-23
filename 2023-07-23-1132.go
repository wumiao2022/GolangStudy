package main

import "fmt"

func main() {
    // Example 1: Hello World
    fmt.Println("Hello, World!")

    // Example 2: Variables
    var name string = "John"
    var age int = 25
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)

    // Example 3: Arrays
    numbers := [4]int{2, 4, 6, 8}
    fmt.Println(numbers)

    // Example 4: Loops
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Example 5: Functions
    fmt.Println(add(3, 5))
}

func add(a, b int) int {
    return a + b
}