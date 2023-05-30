package main

import (
    "fmt"
)

func main() {
    // Example 1: Hello World
    fmt.Println("Hello, World!")

    // Example 2: Variables
    var a string = "Hello"
    b := "World"
    fmt.Println(a, b)

    // Example 3: Conditional Statements
    x := 3
    y := 5
    if x > y {
        fmt.Println("x is greater than y")
    } else {
        fmt.Println("y is greater than x")
    }

    // Example 4: Loops
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
    }

    // Example 5: Arrays
    var numbers [5]int
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    numbers[3] = 4
    numbers[4] = 5
    fmt.Println(numbers)

    // Example 6: Slices
    numbers_slice := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers_slice)

    // Example 7: Maps
    person := map[string]string{
        "name": "John",
        "age":  "30",
        "city": "New York",
    }
    fmt.Println(person)

    // Example 8: Functions
    sum := add(2, 3)
    fmt.Println(sum)
}

func add(a, b int) int {
    return a + b
}