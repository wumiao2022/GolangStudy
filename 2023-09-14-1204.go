package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!"
    fmt.Println("Hello, World!")

    // Exercise 2: Print the sum of two integers
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // Exercise 3: Print the result of multiplying two floats
    float1 := 3.14
    float2 := 2.5
    result := float1 * float2
    fmt.Println("Result:", result)

    // Exercise 4: Print the remainder of dividing two integers
    dividend := 20
    divisor := 3
    remainder := dividend % divisor
    fmt.Println("Remainder:", remainder)

    // Exercise 5: Print the length of a string
    str := "Golang"
    length := len(str)
    fmt.Println("String length:", length)
}