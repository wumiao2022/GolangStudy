package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Enter a number: ")
    fmt.Scanln(&num1)

    fmt.Print("Enter another number: ")
    fmt.Scanln(&num2)

    fmt.Println("Sum:", num1 + num2)
}