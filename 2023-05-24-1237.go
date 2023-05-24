package main

import (
    "fmt"
)

func main() {
    // Example 1: Print "Hello, World!"
    fmt.Println("Hello, World!")

    // Example 2: Print a triangle of stars
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }

    // Example 3: Calculate the factorial of a number
    num := 6
    fact := 1
    for i := 1; i <= num; i++ {
        fact *= i
    }
    fmt.Printf("Factorial of %d is %d\n", num, fact)

    // Example 4: Check if a number is prime
    num = 17
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Printf("%d is prime\n", num)
    } else {
        fmt.Printf("%d is not prime\n", num)
    }
}