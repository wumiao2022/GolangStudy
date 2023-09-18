package main

import "fmt"

func main() {
    // Exercise 1: Print Hello, World!
    fmt.Println("Hello, World!")

    // Exercise 2: Print your name
    fmt.Println("Your Name")

    // Exercise 3: Print the sum of two numbers
    num1 := 5
    num2 := 10
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // Exercise 4: Print a pattern
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }

    // Exercise 5: Calculate the factorial of a number
    num := 5
    factorial := 1
    for i := 1; i <= num; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", num, ":", factorial)
}