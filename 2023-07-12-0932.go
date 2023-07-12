package main

import "fmt"

func main() {
    // Exercise 1: Hello World
    fmt.Println("Hello, World!")

    // Exercise 2: Variables
    var name string = "John"
    age := 25
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)

    // Exercise 3: Constants
    const pi = 3.14159
    radius := 5.0
    area := pi * radius * radius
    fmt.Printf("The area of a circle with radius %.2f is %.2f\n", radius, area)

    // Exercise 4: Conditionals
    number := 10
    if number%2 == 0 {
        fmt.Println("The number is even.")
    } else {
        fmt.Println("The number is odd.")
    }

    // Exercise 5: Loops
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Exercise 6: Arrays
    numbers := [5]int{1, 2, 3, 4, 5}
    fmt.Println(numbers)

    // Exercise 7: Slices
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // Exercise 8: Functions
    result := sum(3, 4)
    fmt.Println(result)
}

func sum(a, b int) int {
    return a + b
}
