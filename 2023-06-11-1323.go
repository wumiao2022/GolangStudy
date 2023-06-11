package main

import (
    "fmt"
)

func main() {
    
    // Example 1: Calculate the average of two numbers
    
    num1 := 10
    num2 := 20
    average := (num1 + num2) / 2
    fmt.Println("The average of", num1, "and", num2, "is", average)
    
    // Example 2: Print even numbers between 1 and 20
    
    for i := 1; i <= 20; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
    
    // Example 3: Reverse a string
    
    message := "hello world"
    reversedMessage := ""
    for i := len(message)-1; i >= 0; i-- {
        reversedMessage += string(message[i])
    }
    fmt.Println(reversedMessage)
    
    // Example 4: Check if a number is prime
    
    number := 17
    isPrime := true
    for i := 2; i < number; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(number, "is prime")
    } else {
        fmt.Println(number, "is not prime")
    }
}