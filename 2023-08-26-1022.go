package main

import "fmt"

func main() {
    // Example 1: Hello World
    fmt.Println("Hello, World!")

    // Example 2: Variables
    var age int = 30
    fmt.Println("My age is", age)

    // Example 3: Constants
    const pi float64 = 3.14159
    fmt.Println("The value of pi is", pi)

    // Example 4: Arrays
    var numbers [5]int
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    numbers[3] = 4
    numbers[4] = 5
    fmt.Println("Array:", numbers)

    // Example 5: Slices
    var names []string
    names = append(names, "Alice")
    names = append(names, "Bob")
    names = append(names, "Charlie")
    fmt.Println("Slice:", names)
}