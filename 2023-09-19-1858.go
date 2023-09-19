package main

import "fmt"

func main() {
    // Exercise 1
    fmt.Println("Exercise 1:")
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 2
    fmt.Println("Exercise 2:")
    num := 5
    if num < 0 {
        fmt.Println("Number is negative")
    } else if num == 0 {
        fmt.Println("Number is zero")
    } else {
        fmt.Println("Number is positive")
    }

    // Exercise 3
    fmt.Println("Exercise 3:")
    colors := []string{"red", "green", "blue"}
    for index, color := range colors {
        fmt.Println(index, color)
    }

    // Exercise 4
    fmt.Println("Exercise 4:")
    i := 1
    for i <= 5 {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
        i++
    }

    // Exercise 5
    fmt.Println("Exercise 5:")
    total := 0
    for i := 1; i <= 100; i++ {
        total += i
    }
    fmt.Println("Sum of numbers from 1 to 100:", total)
}