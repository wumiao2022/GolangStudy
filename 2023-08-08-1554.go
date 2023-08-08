package main

import "fmt"

func main() {
    // Exercise 1: Variables
    var age int = 27
    fmt.Println("I am", age, "years old.")

    // Exercise 2: Arrays
    var fruits [3]string
    fruits[0] = "apple"
    fruits[1] = "banana"
    fruits[2] = "orange"
    fmt.Println("My favorite fruits are", fruits[0], ",", fruits[1], ", and", fruits[2], ".")

    // Exercise 3: Loops
    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }

    // Exercise 4: Functions
    result := multiply(5, 3)
    fmt.Println("5 multiplied by 3 is", result)
}

func multiply(a, b int) int {
    return a * b
}