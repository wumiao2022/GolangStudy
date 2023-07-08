package main

import "fmt"

func main() {
    // Example 1
    fmt.Println("Hello, World!")

    // Example 2
    var num1 int = 10
    var num2 int = 15
    fmt.Println("Sum:", num1+num2)

    // Example 3
    var name string = "John"
    fmt.Println("My name is", name)

    // Example 4
    var isTrue bool = true
    if isTrue {
        fmt.Println("It is true!")
    } else {
        fmt.Println("It is false!")
    }

    // Example 5
    var count int = 1
    for count <= 5 {
        fmt.Println("Count:", count)
        count++
    }
}