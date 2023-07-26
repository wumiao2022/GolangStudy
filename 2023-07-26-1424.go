package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, world!"
    fmt.Println("Hello, world!")

    // Exercise 2: Print numbers from 1 to 10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 3: Find the sum of numbers from 1 to 100
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // Exercise 4: Calculate the factorial of a given number
    fact := 1
    num := 5
    for i := 1; i <= num; i++ {
        fact *= i
    }
    fmt.Println("Factorial of", num, "is", fact)

    // Exercise 5: Determine if a number is prime
    primeNum := 17
    isPrime := true
    for i := 2; i < primeNum; i++ {
        if primeNum%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(primeNum, "is prime:", isPrime)
}