package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!" to the console.
    fmt.Println("Hello, World!")

    // Exercise 2: Declare and initialize a variable with a string value, then print it.
    message := "Hello, Golang!"
    fmt.Println(message)

    // Exercise 3: Create a function that takes in two integers and returns their sum.
    sum := add(5, 3)
    fmt.Println(sum)

    // Exercise 4: Create a loop that prints the numbers from 1 to 10.
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 5: Create a slice of strings with three elements, then print each element.
    colors := []string{"red", "green", "blue"}
    for _, color := range colors {
        fmt.Println(color)
    }
}

func add(x, y int) int {
    return x + y
}