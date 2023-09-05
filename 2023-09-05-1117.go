package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!"
    fmt.Println("Hello, World!")

    // Exercise 2: Print the sum of two numbers, 5 and 3
    num1 := 5
    num2 := 3
    sum := num1 + num2
    fmt.Println("The sum is:", sum)

    // Exercise 3: Print whether a number is even or odd
    number := 13
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // Exercise 4: Print a countdown from 10 to 1
    for i := 10; i >= 1; i-- {
        fmt.Println(i)
    }

    // Exercise 5: Print the Fibonacci sequence up to the 10th number
    num := 10
    fib1 := 0
    fib2 := 1
    fmt.Println(fib1)
    fmt.Println(fib2)
    for i := 2; i < num; i++ {
        fib := fib1 + fib2
        fmt.Println(fib)
        fib1 = fib2
        fib2 = fib
    }
}