package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!"
    fmt.Println("Hello, World!")

    // Exercise 2: Calculate the sum of 2 numbers and print the result
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("The sum is:", sum)

    // Exercise 3: Print numbers from 1 to 10 using a loop
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 4: Calculate the factorial of a given number and print the result
    n := 5
    factorial := 1
    for i := 1; i <= n; i++ {
        factorial *= i
    }
    fmt.Println("The factorial of", n, "is:", factorial)

    // Exercise 5: Reverse a given string and print the result
    s := "Hello, Go!"
    reversed := ""
    for i := len(s) - 1; i >= 0; i-- {
        reversed += string(s[i])
    }
    fmt.Println("The reversed string is:", reversed)
}