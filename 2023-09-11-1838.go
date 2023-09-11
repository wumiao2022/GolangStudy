package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!" using fmt package
    fmt.Println("Hello, World!")

    // Exercise 2: Calculate the sum of two numbers
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // Exercise 3: Swap two numbers using a temporary variable
    a := 5
    b := 10
    fmt.Println("Before swapping - a:", a, "b:", b)
    temp := a
    a = b
    b = temp
    fmt.Println("After swapping - a:", a, "b:", b)

    // Exercise 4: Check if a number is even or odd
    number := 7
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // Exercise 5: Calculate the factorial of a number using recursion
    fact := factorial(5)
    fmt.Println("Factorial of 5:", fact)
}

func factorial(n uint64) uint64 {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}