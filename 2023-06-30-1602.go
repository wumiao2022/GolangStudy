package main

import "fmt"

func main() {
    // Example 1: Hello World
    fmt.Println("Hello, World!")
    
    // Example 2: Variables
    var name string = "John"
    fmt.Println("My name is", name)
    
    // Example 3: Constants
    const age int = 30
    fmt.Println("I am", age, "years old")
    
    // Example 4: Arrays
    numbers := [3]int{1, 2, 3}
    fmt.Println("The first number is", numbers[0])
    
    // Example 5: Slices
    fruits := []string{"apple", "banana", "orange"}
    fmt.Println("I like", fruits[1])
    
    // Example 6: Loops
    for i := 1; i <= 5; i++ {
        fmt.Println("Loop iteration", i)
    }
    
    // Example 7: Functions
    sum := add(2, 3)
    fmt.Println("The sum is", sum)
}

func add(a, b int) int {
    return a + b
}