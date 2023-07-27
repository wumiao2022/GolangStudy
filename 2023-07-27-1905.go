package main

import (
    "fmt"
)

func main() {
    // 1. Hello World
    fmt.Println("Hello, World!")

    // 2. Print Number
    num := 10
    fmt.Println(num)

    // 3. Print String
    str := "Golang"
    fmt.Println(str)

    // 4. Arithmetic Operations
    num1 := 10
    num2 := 5
    fmt.Println(num1 + num2)    // Addition
    fmt.Println(num1 - num2)    // Subtraction
    fmt.Println(num1 * num2)    // Multiplication
    fmt.Println(num1 / num2)    // Division
    fmt.Println(num1 % num2)    // Modulus

    // 5. For Loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 6. If-Else Statements
    x := 10
    y := 5
    if x > y {
        fmt.Println("x is greater than y")
    } else {
        fmt.Println("x is less than or equal to y")
    }

    // 7. Switch Statement
    day := "Sunday"
    switch day {
    case "Sunday":
        fmt.Println("Today is Sunday")
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    case "Wednesday":
        fmt.Println("Today is Wednesday")
    case "Thursday":
        fmt.Println("Today is Thursday")
    case "Friday":
        fmt.Println("Today is Friday")
    case "Saturday":
        fmt.Println("Today is Saturday")
    default:
        fmt.Println("Invalid day")
    }
}